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
	g *gorm.DB
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

	g.AutoMigrate(&User{})
}

// 通过手机号码和密码查找用户
func GetUserByPhoneAndPassword(phone, pwd string) (user User, err error) {

	err = g.Where("phone = ? AND password = ?", phone, pwd).Limit(1).Find(&user).Error
	return user, err
}

func CreateUser(user *User) (err error) {
	if user == nil {
		err = errors.New("invalid params")
		return err
	}

	err = g.Create(user).Error
	return err

}

func GetUserByPhone(phone string) (user User, err error) {
	err = g.Where("phone = ?", phone).Limit(1).Find(&user).Error
	return user, err
}
