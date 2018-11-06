package main

import (
	_ "beego_blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbuser := beego.AppConfig.String("dbuser")
	dbpass := beego.AppConfig.String("dbpass")
	dbname := beego.AppConfig.String("dbname")
	dburl := beego.AppConfig.String("dburl")
	dbport := beego.AppConfig.String("dbport")
	dsn := dbuser + ":" + dbpass + "@tcp(" + dburl + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true // 调试模式打印查询语句
	}

	beego.Run()
}
