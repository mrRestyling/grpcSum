package main

import (
	"context"
	"flag"
	"log"
	"myGRPC/proto/createdProto"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Println(x, " не цифра")
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Println(y, " не цифра")
	}

	// подключаемся к серверу gRPC
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
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
