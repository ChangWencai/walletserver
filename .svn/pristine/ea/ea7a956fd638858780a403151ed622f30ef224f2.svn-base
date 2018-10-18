package routers

import (
	"github.com/astaxie/beego"
	"wadmin/controllers"
)

func init() {

	// admin
	beego.Router("/login", &controllers.AdminLoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.AdminLoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.AdminLoginController{}, "*:NoAuth")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	//beego.Router("/message", &controllers.UserEmailController{}, "*:messsage")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.DappController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.UserMailController{})
	beego.AutoRouter(&controllers.AnnounceController{})
	beego.AutoRouter(&controllers.DappMenuController{})
	//beego.AutoRouter(&controllers.NoticeController{})
}
