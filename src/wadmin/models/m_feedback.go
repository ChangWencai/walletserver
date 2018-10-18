package models

import "github.com/astaxie/beego/orm"

type Feedback struct {
	Id 				int
	SendId 			int
	Catalog    	  	int
	Message			string
	CreateTime 		int64
	Status 			int
}

func (a *Feedback) TableName() string {
	return TableName("feedback")
}

func GetFeedbackById(id int) (*Feedback, error) {
	r := new(Feedback)
	err := orm.NewOrm().QueryTable(TableName("feedback")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func GetFeedbackList(page, pageSize int, filters ...interface{}) ([]*Feedback, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Feedback, 0)
	query := orm.NewOrm().QueryTable(TableName("feedback"))
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

func (a *Feedback) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
