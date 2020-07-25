package model

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type User struct {
	Id int `orm:"id"`
	Username string `orm:"username"`
	Nickname string `orm:"nickname"`
	Password string `orm:"password"`
	Salt string `orm:"salt"`
	Avatar string `orm:"avatar"`
	Phone string `orm:"phone"`
	Email string `orm:"email"`
	Createtime int `orm:"createtime"`
	Updatetime int `orm:"updatetime"`
}

//获取orm对象
func (o *User) GetInstance() *gdb.Model {
	m := g.DB().Table("user")
	return m
}