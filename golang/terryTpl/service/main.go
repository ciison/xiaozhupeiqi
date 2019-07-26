package main

import (
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"xiaozhupeiqi/golang/terryTpl/service/handler"
	"xiaozhupeiqi/golang/terryTpl/service/subscriber"

	srv "xiaozhupeiqi/golang/terryTpl/service/proto/service"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("terry.user.srv.service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	srv.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("terry.user.srv.service", service.Server(), new(subscriber.Service))

	// Register Function as Subscriber
	micro.RegisterSubscriber("terry.user.srv.service", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
