package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go/api"
	"log"
	"os"
	"strconv"
)

func main() {

	conn, _ := grpc.Dial("127.0.0.1:8886", grpc.WithInsecure())

	client := api.NewServiceClient(conn)

	p1 := os.Args[1]
	p2 := os.Args[2]

	a, _ := strconv.Atoi(p1)
	b, _ := strconv.Atoi(p2)

	r, err := client.Sum(context.Background(), &api.Request{A: int32(a), B: int32(b)})

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	log.Println("Sum:", r.Sum)
}
