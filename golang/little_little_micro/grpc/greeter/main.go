package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"xiaozhupeiqi/golang/little_little_micro/grpc/greeter/handler"
	"xiaozhupeiqi/golang/little_little_micro/grpc/greeter/subscriber"

	greeter "xiaozhupeiqi/golang/little_little_micro/grpc/greeter/proto/greeter"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.greeter", service.Server(), new(subscriber.Greeter))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.greeter", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
