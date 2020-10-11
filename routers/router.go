package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//用户注册的接口请求
	beego.Router("/user_register",&controllers.RegisterContorller{})

}
