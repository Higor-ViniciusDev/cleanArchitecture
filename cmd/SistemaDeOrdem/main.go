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
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/database"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/graph"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/services"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web/webserver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	configs, err := configsPackge.LoadConfig("./cmd/SistemaDeOrdem")
	if err != nil {
		panic(err)
	}

	//"root:root@tcp(localhost:3306)/ordens"
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUsuario, configs.DBSenha, configs.DBHost, configs.DBPorta, configs.DBNome))

	if err != nil {
		panic(err)
	}

	novoRepository := database.NewOrdemRepository(db)
	ordersHandle := web.NewOrdensHandler(novoRepository)

	webServeR := webserver.NewWebServer(fmt.Sprintf(":%s", configs.WebServerPorta))
	webServeR.AdicionarHandle("/ordens", ordersHandle.CriarOrdem, "POST")

	go webServeR.StartWebServer()

	ordensService := services.NewOrdemService(novoRepository)

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
		Repository: novoRepository,
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
