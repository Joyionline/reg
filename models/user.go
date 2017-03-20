package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//定义用户信息表
type User struct {
	Id       int64
	Username string `orm:"unique;size(64)"`
	Pwd      string `orm:"size(64)`
}

//定义表名
func (u *User) TableName() string {
	return "user"
}

//模型注册
func init() {
	orm.RegisterModel(new(User))
}

// func GetUser(uname,pwd string) bool  {
// 	o := orm.NewOrm()
// 	b := o.QueryTable("user")
// }

/*
添加用户
*/
func AddUser(u *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	if err == nil {
		beego.Debug("id", id)
	}
	return id, err
}

/*
查询用户信息
*/
func (u *User) ReadUserInfo() (err error) {
	o := orm.NewOrm()
	//进行查询
	err = o.Read(u)
	return err
}

//查询用户是否存在
func GetUser(username, password string) bool {
	o := orm.NewOrm()
	//查询
	b := o.QueryTable("user").Filter("Username", username).Filter("Pwd", password).Exist()
	return b
}
