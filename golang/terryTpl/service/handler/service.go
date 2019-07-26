package handler

import (
	"context"
	"github.com/astaxie/beego"
	"xiaozhupeiqi/golang/terryTpl/model"

	"github.com/micro/go-micro/util/log"

	service "xiaozhupeiqi/golang/terryTpl/service/proto/service"
)

type Service struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Service) Call(ctx context.Context, req *service.Request, rsp *service.Response) error {
	log.Log("Received Service.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// 登录
func (e *Service) Login(ctx context.Context, input *service.Input, output *service.Output) error {
	beego.Info("call Login")

	if input == nil || output == nil {
		return nil
	}
	user := model.User{}
	user.Phone = input.Phone
	user.Password = input.Password

	user, err := model.GetUserByPhoneAndPassword(user.Phone, user.Password)
	if err != nil {
		output.Code = 2
		output.Msg = "login_err"
		return nil
	}
	// 登录成功
	output.Msg = "success"
	output.Code = 1
	return nil

}

// 注册
func (e *Service) Register(ctx context.Context, input *service.Input, output *service.Output) error {
	beego.Info("call Register")
	if input == nil || output == nil {
		return nil
	}

	user := model.User{}

	user.Phone = input.Phone
	user, err := model.GetUserByPhone(input.Phone)
	// 表示这个用户已经注册了
	if user.ID != 0 {
		output.Code = 2
		output.Msg = "register_err"
		return nil
	}

	user.Phone = input.Phone
	user.Password = input.Password
	err = model.CreateUser(&user)
	if err != nil {
		output.Code = 2
		output.Msg = "register_err"
		return nil
	}

	output.Code = 1
	output.Msg = "success"

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Service) Stream(ctx context.Context, req *service.StreamingRequest, stream service.Service_StreamStream) error {
	log.Logf("Received Service.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&service.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Service) PingPong(ctx context.Context, stream service.Service_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&service.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
