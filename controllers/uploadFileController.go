package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
)

type UploadFileController struct {
	beego.Controller
}

//使用post方法上传

func (u *UploadFileController)Post(){
	//1、获取客户端上传的文件、以及其他from表单信息
	//获取标题
	fileTitle :=u.Ctx.Request.PostFormValue("upload_title")
	//获取文件
	file,header,err :=u.GetFile("upload_file")
	if err != nil{
		u.TplName = "fileErrorPage.hyml"
		return
	}
	fmt.Println("自定义的文件名称：",fileTitle)
	fmt.Println("文件名称",header.Filename)
	fmt.Println("文件大小",header.Size)//字节大小

	fmt.Println(file)

	u.Ctx.WriteString("解析到上传文件，文件名是："+header.Filename)
	//2、将文件保存在本地的一个目录当中
	//文件全路径：路径 +文件名 +"."+扩展名
	//要的文件的路径
	uploadDir := "./static/img" + header.Filename
	//创建一个writer：用于向硬盘上写一个文件
	io.Copy(,file)
	//3、关闭文件
	//4、根据文件保存结果，返回相应的提示信息或者页面跳转
}
