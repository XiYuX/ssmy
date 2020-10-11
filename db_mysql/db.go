package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func ConnrctDB(){
	//1、读取conf配置信息
	config := beego.AppConfig
	dbDriver := config.String("driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//2、组织链接数据库字符串
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	//3、链接数据库
	db, err := sql.Open(dbDriver, connUrl)
	//4、获取数据库链接对象，处理链接结果
	if err != nil {
		panic("数据库连接失败，请重试")
	}

	DB = db//为全局的数据库操作对象赋值
}


