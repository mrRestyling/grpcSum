package main

import (
	"log"
	"myGRPC/internal/adder"
	"myGRPC/proto/createdProto"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// grpc.NewServer() - создание нового сервера gRPC
	s := grpc.NewServer()

	// srv - создает новый экземпляр структуры

	srv := &adder.GRPCServer{}

	// Экземпляр srv будет использоваться как обработчик RPC-запросов для сервера gRPC
	// Регистрирует экземпляр srv в сервере s как обработчик RPC-запросов для сервиса Adder, определенного в файле протокола.
	createdProto.RegisterAdderServer(s, srv)

	// net.Listen - Создает TCP-сокет и начинает слушать входящие соединения на порту 8080.
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Запуск сервера gRPC
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}
