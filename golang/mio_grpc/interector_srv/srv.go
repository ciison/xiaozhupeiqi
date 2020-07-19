package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mio_grpc/pb"
	"net"
	"os"
	"time"
)

type greeterImpl struct {
}
type HandlerFunc func(*Context)
type Context struct {
	req         interface{}
	resp        interface{}
	info        *grpc.UnaryServerInfo
	ctx         context.Context
	err         error             // 错误
	reqJsData   string            // 输入参数js格式的字符串
	respJsData  string            // 输入参数的js格式的字符串
	handlerFunc []HandlerFunc     // 这里默认有一个处理 handler
	handler     grpc.UnaryHandler //
	index       int               // 当前回调所在的层数
	data        map[string]interface{}
}

func newContext(ctx context.Context, req, resp interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) *Context {
	data, _ := json.Marshal(req)
	return &Context{
		req:         req,
		resp:        resp,
		info:        info,
		err:         nil,
		ctx:         ctx,
		reqJsData:   string(data),
		respJsData:  "",
		handlerFunc: make([]HandlerFunc, 0),
		handler:     handler,
		data:        make(map[string]interface{}, 16),
	}
}

// 参数检查
func validate() HandlerFunc {
	var (
		validate = validator.New()
		uni      = ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh")
	)
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	return func(c *Context) {
		// 参数校验
		fmt.Println("validate ", c.GetReq())
		if err := validate.Struct(c.req); err != nil {
			if transErr, ok := err.(validator.ValidationErrors); ok {
				translations := transErr.Translate(trans)
				var buf bytes.Buffer
				for _, s2 := range translations {
					buf.WriteString(s2)
				}
				err = status.New(codes.InvalidArgument, buf.String()).Err()
				c.AbortWith(err)
				return
			}
			err = status.New(codes.Unknown, fmt.Sprintf("error%s", err)).Err()
			c.AbortWith(err)
			return
			fmt.Println("error ", err)
		} else {
			fmt.Println("参数校验", err, c.GetReqJsData())
		}
	}

}

// 调用 处理函数
func procHandler(ctx *Context) {
	ctx.resp, ctx.err = ctx.handler(ctx.ctx, ctx.req)
}

// 获取请求参数
func (c *Context) GetReq() interface{} {
	return c.req
}

func (c *Context) GetReqJsData() string {
	if c == nil {
		return ""
	}
	return c.reqJsData
}

var log2es = func() HandlerFunc {
	// 模拟输入日志到 es
	file, err := os.OpenFile("./my.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()
	return func(c *Context) {
		start := time.Now()

		c.Next() // 请求下一个方法
		writeString, err2 := file.WriteString(
			fmt.Sprintf(
				"method:%s, status:%v, latency:%v, req:%s, resp:%s\n", c.FullMethod(), status.Convert(c.err).Code().String(), time.Now().Sub(start), c.GetReqJsData(), ""))
		fmt.Println("log2es after ", time.Now(), writeString, err2, c.FullMethod())
	}
}

func (c *Context) SetReqJsData(str string) {
	if c == nil {
		return
	}
	if json.Valid([]byte(str)) {
		c.reqJsData = str
	}
}

func (c *Context) SetRespJsData(str string) {
	if c == nil {
		return
	}
	if json.Valid([]byte(str)) {
		c.respJsData = str
	}
}

func (c *Context) FullMethod() string {
	if c == nil || c.info == nil {
		return ""
	}
	return c.info.FullMethod
}

func (c *Context) SetData(key string, value interface{}) {
	if c == nil {
		return
	}
	c.data[key] = value
}

func (c *Context) GetData(key string) interface{} {
	if c == nil {
		return nil
	}
	return c.data[key]
}

func (c *Context) Next() {
	if c == nil {
		return
	}
	c.index++
	for (c.index) < len(c.handlerFunc) {
		c.handlerFunc[c.index](c)
		c.index++
	}

}

func (c *Context) AbortWith(err error) {
	c.err = err
	c.index = (1 << 32)
}

// 包裹处理多个 handler func
func WrapperHandler(handFunc ...HandlerFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		c := newContext(ctx, req, resp, info, handler)
		c.handlerFunc = append(c.handlerFunc, func(c *Context) {
			defer func() {
				if err := recover(); err != nil {
					c.AbortWith(status.New(codes.Internal, fmt.Sprintf("errors:%v", err)).Err())
				}
			}()
			c.Next()
		})
		c.handlerFunc = append(c.handlerFunc, handFunc...)
		c.handlerFunc = append(c.handlerFunc, procHandler)

		for c.index = 0; c.index < len(c.handlerFunc); c.index++ {
			c.handlerFunc[c.index](c)
		}
		return c.resp, c.err
	}
}

func (g greeterImpl) OutOfIndex(ctx context.Context, request *pb.OutOfIndexRequest) (resp *pb.OutOfIndexResponse, err error) {
	fmt.Println("OutOfIndex", request)
	time.Sleep(time.Second * 4)
	//defer func() {
	//	// 拦截错误
	//	if err := recover(); err != nil {
	//		glog.Errorf("panic:%s\n", string(debug.Stack()))
	//	}
	//}()
	resp = &pb.OutOfIndexResponse{Data: "ok"}
	request.Ids = make([]int64, 0)
	request.Ids[1] = 1
	return
}

func (g greeterImpl) NilPointer(ctx context.Context, request *pb.NilPointerRequest) (*pb.NilPointerResponse, error) {
	request.Data.Data = "work man"
	return &pb.NilPointerResponse{Data: "ok"}, nil

}

func (g greeterImpl) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	//panic("implement me")
	fmt.Println("Hello:", request)
	return &pb.HelloResponse{ErrCode: "err_code"}, nil
}

func main() {
	flag.Parse()
	srv := grpc.NewServer(grpc.UnaryInterceptor(WrapperHandler(log2es(), validate())))
	listen, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err)
	}
	pb.RegisterGreeterServer(srv, &greeterImpl{})
	if err = srv.Serve(listen); err != nil {
		panic(err)
	}

}
