package utils

import (
	"github.com/astaxie/beego"


)

func InitTemplate() {
	//beego.AddFuncMap("getUsername", models.GetUsername)

	beego.AddFuncMap("getDate", GetDate)
	beego.AddFuncMap("getDateMH", GetDateMH)

}
