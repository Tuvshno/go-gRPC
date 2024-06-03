package main

import (
	"context"
	"flag"
	"fmt"
	"go-gRPC/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	serverAddr := flag.String(
		"server", "localhost:8080",
		"The server address in the format of host:port",
	)
	flag.Parse()

	// TLS credentials for secure communication
	//creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})

	// gRPC DialOptions
	//opts := []grpc.DialOption{
	//	grpc.WithTransportCredentials(creds),
	//}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewCalculatorClient(conn)
	res, err := client.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{10, 10, 10, 10, 10},
	})

	if err != nil {
		log.Fatalln("error sending request:", err)
	}

	fmt.Println("result:", res)
}
