package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"mio_grpc/pb"
	"time"
)

func main() {
	cli()
	nilPoint()
	outOfIndex()
}

func cli() {
	conn, err := grpc.DialContext(context.Background(), ":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	hello, err := client.Hello(context.TODO(), &pb.HelloRequest{BookTime: time.Now().Unix(), Id: -233})
	if err != nil {
		fmt.Println(status.Convert(err).Message())
		fmt.Println(status.Convert(err).Code().String())
		return
	}
	fmt.Println(hello)
}

func nilPoint() {
	conn, err := grpc.DialContext(context.Background(), ":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.NilPointer(context.TODO(), &pb.NilPointerRequest{})
	if err != nil {
		fmt.Println(status.Convert(err).Message())
		return
	}
	fmt.Println(resp)
}

func outOfIndex() {
	conn, err := grpc.DialContext(context.Background(), ":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.OutOfIndex(context.TODO(), &pb.OutOfIndexRequest{})
	if err != nil {
		fmt.Println("outOfIndex:", status.Convert(err).Message())
		return
	}
	fmt.Println(resp)
}
