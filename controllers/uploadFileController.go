package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io"
)

/**
 * 文件上传控制器
 */
type UploadFileController struct {
	beego.Controller
}


/**
 * 使用post方法处理文件的上传
 */
func (u *UploadFileController) Post(){
	//1、获取客户端上传的文件以及其他form表单的信息

	//标题
	fileTitle := u.Ctx.Request.PostFormValue("upload_title")
	//文件
	file, header, err :=u.GetFile("upload_file")
	if err != nil {
		u.Ctx.WriteString("抱歉，用户文件解析失败，请重试")
		return
	}
	fmt.Println("自定义的文件标题:",fileTitle)
	fmt.Println("文件名称:",header.Filename)
	fmt.Println("文件的大小:",header.Size)//字节大小

	fmt.Println(file)
	u.Ctx.WriteString("解析到上传文件，文件名是："+header.Filename)

	//2、将文件保存在本地的一个目录中
	//文件全路径： 路径 + 文件名 + "." + 扩展名
	//要的文件的路径
	uploadDir := "./static/img/" + header.Filename
	//创建一个writer: 用于向硬盘上写一个文件

	io.Copy(uploadDir,file)
	//3、关闭文件
	//4、根据文件保存结果，返回相应的提示信息或者页面跳转

}
