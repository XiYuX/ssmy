package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RegisterContorller struct {
	beego.Controller
}

func (r *RegisterContorller) post() {
	//1、解析请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err!= nil {
		//返回错误页面
		r.TplName = "analysisErrorPage.html"
		return
	}
	//2、保存用户信息到数据库
	_,err =user.SaveUser()
	//3、返回前端结果
	if err !=nil{
		r.TplName = "userErrorPage.html"
		return
	}
	//用户注册成功
	r.TplName = "login.html"
}
