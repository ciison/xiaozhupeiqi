package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/util/log"

	service "xiaozhupeiqi/golang/little_little_micro/grpc/grpcmicro/service/proto/service"
)

type Service struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Service) Call(ctx context.Context, req *service.Request, rsp *service.Response) error {
	log.Log("Received Service.Call request")
	// 这里是 rpc 的调用返回结果
	rsp.Msg = "大家好！ 我是 xiao 成, 现在是 "+ time.Now().Format("2006-01-02")
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
