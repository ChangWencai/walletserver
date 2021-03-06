package models

import (
	"github.com/astaxie/beego/orm"
)

type Dapp struct {
	Id 				int
	UserId 			int
	DappName 		string
	DappHost    	string
	DappImg			string
	DappAuthor		string
	Description     string
	Catalog         int //1: banner 2: 游戏 3: 交易所 4: 冒险 5: 其它
	Popularity      int
	Status 			int // 0: 下架, 1:上架
	CreateTime 	    int64
	UpdateTime 		int64
}


func (a *Dapp) TableName() string {
	return TableName("dapp")
}

func AddDapp(a *Dapp) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GetDappByName(DappName string) (*Dapp, error) {
	a := new(Dapp)
	err := orm.NewOrm().QueryTable(TableName("dapp")).Filter("dapp_name", DappName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetDappList(page, pageSize int, filters ...interface{}) ([]*Dapp, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Dapp, 0)
	query := orm.NewOrm().QueryTable(TableName("dapp"))
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

func getDappKinds()  {
	
}

func GetDappByIds(ids string) ([]*Dapp, error) {
	list := make([]*Dapp, 0)
	sql := "SELECT * FROM wt_dapp WHERE id in(" + ids + ")"
	orm.NewOrm().Raw(sql).QueryRows(&list)

	return list, nil
}

func GetDappById(id int) (*Dapp, error) {
	r := new(Dapp)
	err := orm.NewOrm().QueryTable(TableName("dapp")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Dapp) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
