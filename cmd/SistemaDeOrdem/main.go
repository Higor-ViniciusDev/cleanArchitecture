package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	configsPackge "github.com/Higor-ViniciusDev/CleanArchiteture/configs"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/events/handlers"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/graph"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/services"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web/webserver"
	"github.com/Higor-ViniciusDev/CleanArchiteture/pkg/events"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs, err := configsPackge.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	//"root:root@tcp(localhost:3306)/ordens"
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUsuario, configs.DBSenha, configs.DBHost, configs.DBPorta, configs.DBNome))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMqCanal := getRabbitMqChannel(configs)

	eventorDisparador := events.NewEventoDisparador()
	eventorDisparador.RegistrarHandler("OrdemCreated", &handlers.OrderCreatedHandler{
		RabbitMQChannel: rabbitMqCanal,
	})

	createOrderUseCase := NewCreateOrdemUseCaseInje(db, eventorDisparador)
	listAllOrderUseCase := NewListAllOrdemUseCaseInje(db)

	webOrdersHandle := NewWebOrdersHandleInje(createOrderUseCase, listAllOrderUseCase)
	webServeR := webserver.NewWebServer(fmt.Sprintf(":%s", configs.WebServerPorta))
	webServeR.AdicionarHandle("/ordens", webOrdersHandle.CriarOrdem, "POST")
	webServeR.AdicionarHandle("/order", webOrdersHandle.ListarOrdens, "GET")

	go webServeR.StartWebServer()

	ordensService := services.NewOrderService(*createOrderUseCase)

	grpcServe := grpc.NewServer()
	pb.RegisterOrdemServiceServer(grpcServe, ordensService)
	reflection.Register(grpcServe)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPorta))

	if err != nil {
		panic(err)
	}

	fmt.Println("Servidor GRPC Rodando na porta", configs.GRPCServerPorta)
	go grpcServe.Serve(listen)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UseCaseOrder: *createOrderUseCase,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("server graphql iniciado na porta %s", configs.GraphQLServerPorta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.GraphQLServerPorta), nil))
}

func getRabbitMqChannel(configs *configsPackge.Conf) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		configs.RabbitMQUser, configs.RabbitMQPass, configs.RabbitMQHost, configs.RabbitMQPort))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
