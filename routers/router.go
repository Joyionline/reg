package routers

import (
	"reg/controllers"
	"reg/models"

	"github.com/astaxie/beego"
)

func init() {
	models.Connect()
	router()
}

//定义路由的函数
func router() {
	beego.Info("init routers start ... ")
	beego.Router("/", &controllers.UserController{}, "*:Index")
	beego.Router("/regist", &controllers.UserController{}, "*:Regist")
	beego.Router("/success", &controllers.UserController{}, "*:RegSuccess")
	beego.Router("/reghandle", &controllers.UserController{}, "*:RegHandle")
	beego.Router("/tologin", &controllers.UserController{}, "*:ToLogin")
	beego.Router("/loginhandle", &controllers.UserController{}, "*:LoginHandle")
	beego.Router("/error", &controllers.UserController{}, "*:Error")
	beego.Router("/loginsuccess", &controllers.UserController{}, "*:LoginSuccess")

}
