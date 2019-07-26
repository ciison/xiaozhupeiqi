package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	service "xiaozhupeiqi/golang/terryTpl/service/proto/service"
)

type Service struct{}

func (e *Service) Handle(ctx context.Context, msg *service.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *service.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
