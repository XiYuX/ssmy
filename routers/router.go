package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//用户注册的接口请求
	beego.Router("/user_register",&controllers.RegisterController{})
	//直接登录的页面请求接口
	beego.Router("/login.html",&controllers.LoginController{})
	//用户登录请求接口
	beego.Router("/user_login",&controllers.LoginController{})
	//用户登入转注册
	beego.Router("/register_register",&controllers.RegisterController{})
	//重置密码界面转登入
	beego.Router("/upload",&controllers.UploadFileController{})











}
