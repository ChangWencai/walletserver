package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	orm.RegisterModel(new(Auth), new(Role), new(RoleAuth), new(Admin), new(Dapp), new(User))
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}