package main

import (
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1、链接数据库
	db_mysql.ConnctDB()

	//设置静态资源文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

