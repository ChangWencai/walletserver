package models

import (
	"github.com/astaxie/beego/orm"
)

type DappMenu struct {
	Id 				int    `json:"id, omitempty"`
	Catalog         int    `json:"catalog"`// 1: banner 2: 游戏 3: 交易所 4: 冒险 5: 其它
	Name 		    string `json:"name"`
	Status 			int    `json:"status"`// 0: 下架, 1:上架
}

func (u *DappMenu) TableName() string {
	return TableName("dapp_menu")
}

func init() {
	orm.RegisterModel(new(DappMenu))
}

func DappMenus() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(DappMenu))
}

//// 检测DappName是否存在
//func CheckDappName(dappName string) bool {
//	exist := Dapps().Filter("DappName", dappName).Exist()
//	return exist
//}
//
//// GetDappById retrieves Dapp by Id. Returns error if
//// Id doesn't exist
//func GetDappById(id int) (v *Dapp, err error) {
//	o := orm.NewOrm()
//	v = &Dapp{Id: id}
//	if err = o.QueryTable(new(Dapp)).Filter("Id", id).RelatedSel().One(v); err == nil {
//		return v, nil
//	}
//	return nil, err
//}
//
//// GetDappByUser retrieves Dapp by User. Returns error if
//// Id doesn't exist
//func GetDappsByUserId(userId int, fields []string, limit int, offset int) (dapps []*Dapp, err error) {
//	o := orm.NewOrm()
//	if _, err = o.QueryTable(new(Dapp)).Filter("user_id", userId).Limit(limit, offset).All(&dapps, fields...); err == nil {
//		return dapps, nil
//	}
//	return nil, err
//}
//
//// GetAllDapps retrieves all Dapp matches certain condition. Returns empty list if
//// no records exist
//func GetAllDapps(query map[string]string, fields []string, sortby []string, order []string,
//	offset int, limit int, userId int) (ml []interface{}, totalCount int64, err error) {
//	o := orm.NewOrm()
//	qs := o.QueryTable(new(Dapp))
//	// query k=v
//	for k, v := range query {
//		// rewrite dot-notation to Object__Attribute
//		k = strings.Replace(k, ".", "__", -1)
//		qs = qs.Filter(k, v)
//	}
//	// order by:
//	var sortFields []string
//	if len(sortby) != 0 {
//		if len(sortby) == len(order) {
//			// 1) for each sort field, there is an associated order
//			for i, v := range sortby {
//				orderby := ""
//				if order[i] == "desc" {
//					orderby = "-" + v
//				} else if order[i] == "asc" {
//					orderby = v
//				} else {
//					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//			qs = qs.OrderBy(sortFields...)
//		} else if len(sortby) != len(order) && len(order) == 1 {
//			// 2) there is exactly one order, all the sorted fields will be sorted by this order
//			for _, v := range sortby {
//				orderby := ""
//				if order[0] == "desc" {
//					orderby = "-" + v
//				} else if order[0] == "asc" {
//					orderby = v
//				} else {
//					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
//				}
//				sortFields = append(sortFields, orderby)
//			}
//		} else if len(sortby) != len(order) && len(order) != 1 {
//			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
//		}
//	} else {
//		if len(order) != 0 {
//			return nil, 0, errors.New("Error: unused 'order' fields")
//		}
//	}
//
//	var l []Dapp
//	qs = qs.OrderBy(sortFields...).RelatedSel()
//	totalCount, err = qs.Filter("UserId", userId).Count()
//	if _, err = qs.Filter("UserId", userId).Limit(limit, offset).All(&l, fields...); err == nil {
//		if len(fields) == 0 {
//			for _, v := range l {
//				ml = append(ml, v)
//			}
//		} else {
//			// trim unused fields
//			for _, v := range l {
//				m := make(map[string]interface{})
//				val := reflect.ValueOf(v)
//				for _, fname := range fields {
//					m[fname] = val.FieldByName(fname).Interface()
//				}
//				ml = append(ml, m)
//			}
//		}
//		return ml, totalCount, nil
//	}
//	return nil, 0, err
//}
//
//// UpdateDappById update feedback by Id
//func UpdateDappById(m *Dapp)  (err error) {
//	o := orm.NewOrm()
//	v := Dapp{Id: m.Id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		status := m.Status
//		*m = v
//		m.Status = status
//		if num, err = o.Update(m); err == nil {
//			logs.Info("Number of records updated in database: ", num)
//		}
//	}
//	return
//}
//
//// DeleteDapp deletes dapp by Id and returns error if
//// the record to be deleted doesn't exist
//func DeleteDapp(id int) (err error) {
//	o := orm.NewOrm()
//	v := Dapp{Id: id}
//	// ascertain id exists in the database
//	if err = o.Read(&v); err == nil {
//		var num int64
//		if num, err = o.Delete(&Dapp{Id: id}); err == nil {
//			logs.Info("Number of records deleted in database: ", num)
//		}
//	}
//	return
//}
//
//// GetBanner get banner by catalog and so on,return banner array
//
//func GetBanner(line int) (v []Banner, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").From("wt_dapp").
//		Where("status = 1 and catalog = 1").OrderBy("popularity").Desc().Limit(3).Offset(0)
//
//	banner := qb.String()
//	o := orm.NewOrm()
//	var userBanner [] Banner
//	_, errBanner := o.Raw(banner).QueryRows(&userBanner)
//	if errBanner == nil {
//		return userBanner, nil
//	} else {
//		return userBanner, errBanner
//	}
//}
//
//// GetGame get banner by catalog and so on,return banner array
//func GetGame(line int) (v []Games, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").From("wt_dapp").
//		Where("status = 1 and catalog = 2").OrderBy("popularity").Desc().Limit(100).Offset(0)
//
//	game := qb.String()
//	o := orm.NewOrm()
//	var userGame [] Games
//	_, errBanner := o.Raw(game).QueryRows(&userGame)
//	if errBanner == nil {
//		return userGame, nil
//	} else {
//		return userGame, errBanner
//	}
//}
//
//// GetExchange get banner by catalog and so on,return banner array
//func GetExchange(line int) (v []Games, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").From("wt_dapp").
//		Where("status = 1 and catalog = 3").OrderBy("popularity").Desc().Limit(100).Offset(0)
//
//	exchange := qb.String()
//	o := orm.NewOrm()
//	var userExchange [] Games
//	_, errBanner := o.Raw(exchange).QueryRows(&userExchange)
//	if errBanner == nil {
//		return userExchange, nil
//	} else {
//		return userExchange, errBanner
//	}
//}
//
//// GetDangerous get banner by catalog and so on,return banner array
//func GetDangerous(line int) (v []Games, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").From("wt_dapp").
//		Where("status = 1 and catalog = 4").OrderBy("popularity").Desc().Limit(100).Offset(0)
//
//	dangerous := qb.String()
//	o := orm.NewOrm()
//	var userDangerous [] Games
//	_, errBanner := o.Raw(dangerous).QueryRows(&userDangerous)
//	if errBanner == nil {
//		return userDangerous, nil
//	} else {
//		return userDangerous, errBanner
//	}
//}
//
//// GetOther get banner by catalog and so on,return banner array
//func GetOther(line int) (v []Games, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").From("wt_dapp").
//		Where("status = 1 and catalog = 5").OrderBy("popularity").Desc().Limit(100).Offset(0)
//
//	other := qb.String()
//	o := orm.NewOrm()
//	var userOther [] Games
//	_, errBanner := o.Raw(other).QueryRows(&userOther)
//	if errBanner == nil {
//		return userOther, nil
//	} else {
//		return userOther, errBanner
//	}
//}
//
//// GetFind get dapp by name ,return dapp array
//func GetFind(name string) (v []Games, err error){
//	qb, _ := orm.NewQueryBuilder("mysql")
//	qb.Select("description", "dapp_img", "dapp_host", "dapp_name").
//		From("wt_dapp").Where(" status = 1 and  dapp_name like '%%" + name + "%%'").OrderBy("popularity").Desc()
//	find := qb.String()
//	o := orm.NewOrm()
//	var userFind [] Games
//	_, errBanner := o.Raw(find).QueryRows(&userFind)
//	if errBanner == nil {
//		return userFind, nil
//	} else {
//		return userFind, errBanner
//	}
//}