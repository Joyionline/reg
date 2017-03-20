package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	dns := getConfig()
	beego.Info("数据库is", dns)
	//默认需要
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		beego.Error("\n 数据库连接失败", err)
	} else {
		beego.Info("\n 数据库连接sucess")
		orm.RunSyncdb("default", false, true)
		orm.Debug = true
	}
}

/*
* 获取配置
	flag ==1 表示 只链接
	==0 创建 加链接
*/
func getConfig() string {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	//打印信息
	beego.Info("链接数据库")
	//注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//给dns赋值
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)

	return dns
}
