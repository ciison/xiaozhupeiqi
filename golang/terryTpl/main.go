package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-grpc"
	"net/http"
	srv "xiaozhupeiqi/golang/terryTpl/service/proto/service"
)

func main() {

	router := gin.Default()
	router.Delims("[[", "]]") /// 设置模板的渲染标签
	router.LoadHTMLGlob("./golang/terryTpl/view/*")

	router.GET("/login", loginTpl)
	router.GET("/register", registerTpl)

	// 登录的实现
	router.POST("/login", loginImpl)

	// 注册的实现
	router.POST("/register", registerImpl)
	router.Run(":8080")
}

var (
	// 这里造一个假的标准
	GInput = Input{Phone: "13123456789", Password: "123456"}
)

var loginImpl = func(ctx *gin.Context) {
	var (
		input   Input
		retData = make(ResponseParams, 32)
		err     error
	)

	// 获取输入的参数
	if err = json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		beego.Error(err)
		retData["code"] = 2
		retData["msg"] = "invalid_input_params"
		ctx.JSON(http.StatusOK, &retData)
		return
	}

	fmt.Printf("input =%+v\n", input)
	service := grpc.NewService()
	service.Init()
	client := srv.NewService("terry.user.srv.service", service.Client())
	srvInput := srv.Input{}

	srvInput.Phone = input.Phone
	srvInput.Password = input.Password

	// 远程调用
	output, err := client.Login(context.TODO(), &srvInput)
	ctx.JSON(http.StatusOK, output)

}

var registerImpl = func(ctx *gin.Context) {
	var (
		input   Input
		retData = make(ResponseParams, 32)
		err     error
	)

	if err = json.NewDecoder(ctx.Request.Body).Decode(&input); err != nil {
		beego.Error(err)
		retData["code"] = 2
		retData["msg"] = "invalid_input_params"
		ctx.JSON(http.StatusOK, &retData)
		return
	}

	fmt.Printf("input = %+v\n", input)

	srvInput := srv.Input{}
	srvInput.Phone = input.Phone
	srvInput.Password = input.Password

	service := grpc.NewService()
	service.Init()
	client := srv.NewService("terry.user.srv.service", service.Client())
	output, _ := client.Register(context.TODO(), &srvInput)
	ctx.JSON(http.StatusOK, output)

}

// 输入参数
type Input struct {
	Phone    string `json:"phone"`    // 手机号码
	Password string `json:"password"` // 登录口令
}

// 输出参数类型
type ResponseParams map[string]interface{}

// 渲染登录页面
var loginTpl = func(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

// r渲染注册页面
var registerTpl = func(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}
