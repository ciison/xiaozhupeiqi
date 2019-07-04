package main

import (
	"github.com/micro/go-grpc" // grpc
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"xiaozhupeiqi/golang/little_little_micro/grpc/grpcmicro/service/handler"
	"xiaozhupeiqi/golang/little_little_micro/grpc/grpcmicro/service/subscriber"

	srv "xiaozhupeiqi/golang/little_little_micro/grpc/grpcmicro/service/proto/service"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	srv.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.service", service.Server(), new(subscriber.Service))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.service", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
