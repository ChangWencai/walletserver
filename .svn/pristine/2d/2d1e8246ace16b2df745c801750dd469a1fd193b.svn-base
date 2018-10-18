package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id 				int
	UserName 		string
	Password 		string
	Avatar 			string
	Salt 			string
	Phone 			string
	Introduce 		string
	Token 			string
	Gender 			int// 0:Male, 1: Female, 2: undefined
	Age 			int
	Address 		string
	Email 			string
	LastLoginTime 	int64
	Status 			int  // 0: disabled, 1:enabled
	CreatedAt 		int64
	CreateIp 		string
	UpdatedAt 		int64
	Coin 			int
}


func (a *User) TableName() string {
	return TableName("user")
}

func GetUserById(id int) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetUserList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user"))
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

func (a *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
