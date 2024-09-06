package main

import (
	"context"
	"flag"
	"log"
	"myGRPC/proto/createdProto"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Println(x, " не цифра")
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Println(y, " не цифра")
		log.Fatal(err)
	}

	// подключаемся к серверу gRPC
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// создаем клиент gRPC
	// он уже сгенерирован из proto файла в DEMO.pb.go

	client := createdProto.NewAdderClient(conn)
	res, err := client.SendMessage(context.Background(), &createdProto.AddRequest{X: int64(x), Y: int64(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
	log.Println(res.Result)
}
