package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	greeter "xiaozhupeiqi/golang/little_little_micro/grpc/greeter/proto/greeter"
)

type Greeter struct{}

func (e *Greeter) Handle(ctx context.Context, msg *greeter.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *greeter.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
