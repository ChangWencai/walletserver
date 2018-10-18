package models

import (
	"github.com/astaxie/beego/orm"
)

type UserMail struct {
	Id 				int
	SendId 			int
	RecvId 			int
	Subject 		string
	Catalog    	  	int
	Address         string
	Message			string
	CreateTime 		int64
	Status 			int
}

func (a *UserMail) TableName() string {
	return TableName("user_mail")
}

func init() {
	orm.RegisterModel(new(UserMail))
}

func AddUserMail(a *UserMail) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GetUserMailById(id int) (*UserMail, error) {
	r := new(UserMail)
	err := orm.NewOrm().QueryTable(TableName("user_mail")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetUserMailList(page, pageSize int, filters ...interface{}) ([]*UserMail, int64) {
	offset := (page - 1) * pageSize
	list := make([]*UserMail, 0)
	query := orm.NewOrm().QueryTable(TableName("user_mail"))
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

func (a *UserMail) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}