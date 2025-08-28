package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/database"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/grpc/services"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web"
	"github.com/Higor-ViniciusDev/CleanArchiteture/internal/infra/web/webserver"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ordens")

	if err != nil {
		panic(err)
	}

	novoRepository := database.NewOrdemRepository(db)
	ordersHandle := web.NewOrdensHandler(novoRepository)

	webServeR := webserver.NewWebServer(":8080")
	webServeR.AdicionarHandle("/ordens", ordersHandle.CriarOrdem, "POST")

	go webServeR.StartWebServer()

	ordensService := services.NewOrdemService(novoRepository)

	grpcServe := grpc.NewServer()
	pb.RegisterOrdemServiceServer(grpcServe, ordensService)
	reflection.Register(grpcServe)

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	fmt.Println("Servidor GRPC Rodando na porta 50051")
	if err := grpcServe.Serve(listen); err != nil {
		panic(err)
	}
}
