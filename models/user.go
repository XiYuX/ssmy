package models

import (
	"BeegoPackage0922/db_myssql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int	`from:"id"`
	Phone string	`from:"phone"`
	Password string	`from:"password"`

}


//定义一个方法

func (u User) SaveUser() (int64,error) {
	//1、密码脱敏处理
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	//2、执行数据库操作
	row,err :=db_myssql.Db.Exec("insert into user(phone,password)"+"values(?,?)",u.Phone,u.Password)
	if err != nil{
		return -1, err
	}
	id,err := row.RowsAffected()
	if err != nil{
		return -1,err
	}
	return id ,err
	//3、
	//
}