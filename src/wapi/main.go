package main


import (
	_ "wapi/routers"
	"wapi/utils"


	"github.com/astaxie/beego"
	"wapi/controllers"
)

func main() {
	utils.InitSql()
	utils.InitTemplate()
	utils.InitCache()
	utils.InitBootStrap()
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}