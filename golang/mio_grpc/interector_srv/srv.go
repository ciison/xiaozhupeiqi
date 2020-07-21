package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mio_grpc/pb"
)

type greeterImpl struct {
}

// 调用链的处理函数
type HandlerFunc func(*Context)

// 重新包装的 Context;
type Context struct {
	req         interface{}            // 输入参数
	resp        interface{}            // 输出参数
	info        *grpc.UnaryServerInfo  // 服务的信息
	ctx         context.Context        // 服务方法的 上下文信息
	handler     grpc.UnaryHandler      // 对应服务的请求处理
	err         error                  // 错误
	reqJsData   string                 // 输入参数js格式的字符串
	respJsData  string                 // 输入参数的js格式的字符串
	handlerFunc []HandlerFunc          // 这里默认有一个处理 handler
	index       int                    // 当前回调所在的层数
	data        map[string]interface{} // 设置的 data
}

// 新建一个 context
// @param context.Context grpc 方法请求的上下文
// @param req             请求方法的输入参数
// @param resp            请求方法的输出参数
// @param info            grpc 方法名的信息, 里面的参数有 请求的方法名
// @param handler         这个参数是一个函数, 用来调用对应的 grpc 方法
//
// @ret   *Context        返回重新包装的 Context
func newContext(ctx context.Context, req, resp interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) *Context {
	data, _ := json.Marshal(req) // 序列化请求的参数
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
		validate = validator.New() // 校验器
		uni      = ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh") // 中午翻译器
	)
	// 关联校验器和翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	return func(c *Context) {
		// 参数校验
		if err := validate.Struct(c.req); err != nil {
			// 判断错误是否是校验字段的错误
			// 如果是参数校验不通过, 把错误信息翻译成对应的说明
			if transErr, ok := err.(validator.ValidationErrors); ok {
				translations := transErr.Translate(trans)
				var buf bytes.Buffer
				for _, s2 := range translations {
					buf.WriteString(s2)
				}
				// grpc 返回参数错误
				err = status.New(codes.InvalidArgument, buf.String()).Err()
				// 提前终止调用
				c.AbortWith(err)
				return
			}
			// 如果校验 grpc 输入参数的时候遇到错误, 但是错误不是翻译成错误相关的, 返回未知错误
			err = status.New(codes.Unknown, fmt.Sprintf("error%s", err)).Err()
			// 提前终止调用
			c.AbortWith(err)
			return
		}
	}

}

// 获取请求参数
func (c *Context) GetReq() interface{} {
	return c.req
}

// 获取请求参数的 js 格式
func (c *Context) GetReqJsData() string {
	if c == nil {
		return ""
	}
	return c.reqJsData
}

// 设置 JsData
func (c *Context) SetReqJsData(str string) {
	if c == nil {
		return
	}
	if json.Valid([]byte(str)) {
		c.reqJsData = str
	}
}

// 设置返回请求参数的 js 格式的字符串
func (c *Context) SetRespJsData(str string) {
	if c == nil {
		return
	}
	if json.Valid([]byte(str)) {
		c.respJsData = str
	}
}

// 获取当前 grpc 需要请求的方法的名称
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

// 当前调用链方法所在的层数的下一层
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

// 提前终止所有的调用,
func (c *Context) AbortWith(err error) {
	const (
		abortLevel = 1 << 32
	)
	c.err = err
	c.index = abortLevel
}

// 模拟日志输出到 es
func log2es() HandlerFunc {
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
		_, _ = file.WriteString(
			fmt.Sprintf(
				"method:%s, status:%v, latency:%v, req:%s, resp:%s\n", c.FullMethod(), status.Convert(c.err).Code().String(), time.Now().Sub(start), c.GetReqJsData(), ""))
		//fmt.Println("log2es after ", time.Now(), writeString, err2, c.FullMethod())
	}
}

// 调用 处理函数, 这是默认调用的, 这里默认是所有的 调用都完成的时候才调用的
func procHandler(ctx *Context) {
	ctx.resp, ctx.err = ctx.handler(ctx.ctx, ctx.req)
}

// 包裹处理多个 handler func
func WrapperHandler(handFunc ...HandlerFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 从 grpc 拦截器的回调用获取参数, 新生成一个 Context
		c := newContext(ctx, req, resp, info, handler)

		// 构造一个默认的 拦截 panic 的回调处理, 作为第一个业务处理
		c.handlerFunc = append(c.handlerFunc, func(c *Context) {
			defer func() {
				if err := recover(); err != nil {
					c.AbortWith(status.New(codes.Internal, fmt.Sprintf("errors:%v", err)).Err())
				}
			}()
			c.Next()
		})
		// 将用户输入的处理作为中间的处理
		c.handlerFunc = append(c.handlerFunc, handFunc...)
		// 用户需要处理的完成之后, 调用 真正的服务方法
		c.handlerFunc = append(c.handlerFunc, procHandler)

		// 开始调用服务
		for c.index = 0; c.index < len(c.handlerFunc); c.index++ {
			c.handlerFunc[c.index](c)
		}
		// 返回服务 resp 和 err
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
