# gin&go-micro

## html 模板设置
> 某个标签有问题, 😭
``` html
<!DOCTYPE html>
<html lang="zh_cn">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0,minimal-ui:ios">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script src="https://cdn.jsdelivr.net/npm/vue"></script>
    <script src="https://cdn.bootcss.com/axios/0.19.0/axios.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <link rel="stylesheet"
          href="https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u"
          crossorigin="anonymous">
    <title>Title</title>
</head>
<body style="width: 100%;">
<div id="my-app" class="container-fluid" style="width: 100%!important;">
    <div class="container-fluid nav navbar"
         style="background-color: wheat!important; margin-left: -28px; margin-right: -28px; padding-left:16px;">
        <h3>{{message}}</h3>
    </div>
    <div class="container-fluid">
        <div class="row">
            <div class="col-sm-4">
                <p style="font-size: 16px;"> {{leftContent}}</p>
            </div>
            <div class="col-sm-4">
                <p style="font-size: 16px;">
                    {{centerContent}}
                </p>
            </div>
            <div class="col-sm-4" style="padding: 100px;">
                <div>
                    <h3 style="text-align: center">{{rightTitle}}</h3>
                    <form>
                        <div class="form-group">
                            <label>phone</label>
                            <input type="email" class="form-control" id=""
                                   v-model="phone" placeholder="phone">
                        </div>
                        <div class="form-group">
                            <label for="">Password</label>
                            <input type="password" class="form-control" id=""
                                   v-model="password" placeholder="Password">
                        </div>

                        <button type="submit" class="form-control btn btn-primary"
                                v-on:click="demo">Submit
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
    <div class="foot"></div>
</div>
</body>
<script>
    // 这是 vue 的一个实例对象
    let v = new Vue({
        el: "#my-app",
        data: {
            message: "你好, 我是 xiao 成!",
            leftContent: "这是左边栏",
            centerContent: "这是中间的内容区域",
            rightTitle: "登陆",
            password: "",
            phone: "",
        },
        methods: {
            demo: function () {
                console.log("helloWorld")
            }
        },
    })
</script>
</html>
```

## gin 服务器搭建

``` go
package  main 

import (
    "github.com/gin-gonic/gin"
)

func main() {
 
	router:=gin.Default()
	router.Delims("[[", "]]") /// 设置模板的渲染标签
	router.LoadHTMLGlob("./golang/terryTpl/view/*")
	
    // loginTpl 是渲染模板的函数
    router.Get("/login", loginTpl)

    router.Run(":8080") // 设置服务监听的端口
}

```

## gorm 操作 MySQL 数据库

``` go

package model

import (
	"errors"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Phone    string `json:"phone" gorm:"index_unique"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

var (
	g *gorm.DB // 数据库操作句柄
)

func init() {
	var (
		err error
	)

	// 注册数据库

	// 这个是我实体机的数据库地址
	if g, err = gorm.Open("mysql", "test:123456@tcp(192.168.255.127:3306)/test?charset=utf8&parseTime=true&loc=Local"); err != nil {
		panic(err)
	}

	beego.Info("注册数据库成功")
	// 输出日志模式
	g.LogMode(true)

	// 使用数据库迁移
	g.AutoMigrate(&User{})
}

// 通过手机号码和密码查找用户
func GetUserByPhoneAndPassword(phone, pwd string) (user User, err error) {

	// 通过手机号码查找用户信息
	err = g.Where("phone = ? AND password = ?", phone, pwd).Limit(1).Find(&user).Error
	return user, err
}

func CreateUser(user *User) (err error) {
	if user == nil {
		err = errors.New("invalid params")
		return err
	}

	// 创建一个新的用户信息
	err = g.Create(user).Error
	return err

}

func GetUserByPhone(phone string) (user User, err error) {
	// 通过手机号码查找用户信息, 
	err = g.Where("phone = ?", phone).Limit(1).Find(&user).Error
	return user, err
}


```





