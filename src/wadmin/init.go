package main

import (
	"github.com/astaxie/beego"
	"net/url"
	"github.com/astaxie/beego/orm"
	"runtime"

	"wadmin/models"
)

func InitLog() {
	var ostype = runtime.GOOS

	// log
	if ostype == "windows"{
		beego.SetLogger("file", `{"filename":"./log/server.log"}`)
	}else if ostype == "linux" {
		beego.SetLogger("file", `{"filename":"./log/server.log"}`)
	}
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
}

func InitDB() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	// fmt.Println(dsn)

	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	models.Init()
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}