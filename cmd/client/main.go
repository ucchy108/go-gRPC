package main

import (
	hellopb "RPC/pkg/grpc"
	"bufio"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	log.Println("start gRPC Client.")

	// 1.gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3.gRPCクライアントを作成
	client = hellopb.NewGreetingServiceClient(conn)

	Hello("Hoge")
}

func Hello(name string) {
	req := &hellopb.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(res.GetMessage())
	}
}
