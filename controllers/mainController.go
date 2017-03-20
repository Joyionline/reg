package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*
进入首页：登录页面
*/
func (this *MainController) Index() {
	this.TplName = "index.html"
}
