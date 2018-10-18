package models

import "github.com/astaxie/beego/orm"

type Announce struct {
	Id 				int
	UserId 			int
	Subject 		string
	Catalog    	  	int
	Message			string
	CreateTime 		int64
	EndTime 		int64
	Status 			int
}

func (a *Announce) TableName() string {
	return TableName("announce")
}

func init() {
	orm.RegisterModel(new(Announce))
}

func AddAnnounce(a *Announce) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GetAnnounceById(id int) (*Announce, error) {
	r := new(Announce)
	err := orm.NewOrm().QueryTable(TableName("announce")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetAnnounceList(page, pageSize int, filters ...interface{}) ([]*Announce, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Announce, 0)
	query := orm.NewOrm().QueryTable(TableName("announce"))
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

func (a *Announce) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}