package controllers

import (
	"reg/models"

	"fmt"

	"github.com/astaxie/beego"
)

//定义该控制器
type UserController struct {
	beego.Controller
}

/*
进入首页：登录页面
*/
func (this *UserController) Index() {
	this.TplName = "index.html"
}

func (this *UserController) Regist() {
	this.TplName = "regist.html"
}

//定义注册处理方法
func (this *UserController) RegHandle() {
	//接收参数
	username := this.GetString("username")
	pwd := this.GetString("pwd")

	user := new(models.User)
	user.Username = username
	user.Pwd = pwd
	_, err := models.AddUser(user)
	if err != nil {
		// this.Alert("插入数据库错误或者已经插入")
		beego.Error("insert the data err")
		return
	} else {
		this.Ctx.Redirect(302, "/success?username="+username+"&password="+pwd)
	}
}

//定义注册成功的方法
func (this *UserController) RegSuccess() {
	//this.Data["Username"] = this.GetString("username")
	// this.Ctx.WriteString("您的用户名是：" + this.GetString("username"))
	// this.Ctx.WriteString("您的密码是：" + this.GetString("pwd"))
	// user := new(models.User)
	// username := user.Username
	// pwd := user.Pwd
	// fmt.Println(username + pwd)
	this.Data["Username"] = this.GetString("username")
	this.Data["Pwd"] = this.GetString("pwd")
	this.TplName = "success.html"
}

//登录跳转
func (this *UserController) ToLogin() {
	this.TplName = "index.html"
}

//进行登录
func (this *UserController) LoginHandle() {
	//isAction := this.GetString("isAction")
	//if "0" == isAction {
	//接收页面参数
	username := this.GetString("username")
	password := this.GetString("pwd")
	// u := &models.User{
	// 	Username: username,
	// 	Pwd:      password,
	// }
	b := models.GetUser(username, password)
	if b {
		//用户存在
		fmt.Println("用户登录成功。。。")
		beego.Info("用户登录成功。。。。")
		this.Ctx.Redirect(302, "/loginsuccess")
	}
	this.Ctx.Redirect(302, "/error")
}

//登录失败
func (this *UserController) Error() {
	this.TplName = "error.html"
}

//登录成功
func (this *UserController) LoginSuccess() {
	this.TplName = "loginsuccess.html"
}
