package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	greeter "xiaozhupeiqi/golang/little_little_micro/grpc/greeter/proto/greeter"
)

type Greeter struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Greeter) Call(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Log("Received Greeter.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Greeter) Stream(ctx context.Context, req *greeter.StreamingRequest, stream greeter.Greeter_StreamStream) error {
	log.Logf("Received Greeter.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&greeter.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Greeter) PingPong(ctx context.Context, stream greeter.Greeter_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&greeter.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
