package models

import (
	"github.com/astaxie/beego/orm"
)

type DappMenu struct {
	Id 				int
	Catalog    	  	int
	Name            string
	Status          int
}

func (a *DappMenu) TableName() string {
	return TableName("dapp_menu")
}

func init() {
	orm.RegisterModel(new(DappMenu))
}

func DappMenus() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(DappMenu))
}

func AddDappMenu(a *DappMenu) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GetDappMenuById(id int) (*DappMenu, error) {
	r := new(DappMenu)
	err := orm.NewOrm().QueryTable(TableName("dapp_menu")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetDappMenuList(page, pageSize int, filters ...interface{}) ([]*DappMenu, int64) {
	offset := (page - 1) * pageSize
	list := make([]*DappMenu, 0)
	query := orm.NewOrm().QueryTable(TableName("dapp_menu"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (a *DappMenu) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// 检测Catalog是否存在
func CheckDappMenuCatalog(Catalog int) bool {
	exist := DappMenus().Filter("Catalog", Catalog).Exist()
	return exist
}
