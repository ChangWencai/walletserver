package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

const(
	sendId = 0
	cataLog = 1
	handleStatus = 0
	HandleStatus = 1
)

type Feedback struct {
	Id            int     `json:"id, omitempty"`
	SendId        int     `json:"sendid"`
	Catalog       int     `json:"catalog"`
	Message       string  `json:"message"`
	CreateTime 	  int64   `json:"create_time, omitempty"`
	Status        int     `json:"status"`//0,未处理 1,已处理
}

type FeedbackRequestStruct struct {
	Message    string `json:"message"`
	Token      string `json:"token"`
}

func (u *Feedback) TableName() string {
	return TableName("feedback")
}

func init() {
	orm.RegisterModel(new(Feedback))
}

func Feedbacks() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Feedback))
}

// AddFeedBack insert a new feedback data into database
func AddFeedBack(m *FeedbackRequestStruct) (err error) {
	o := orm.NewOrm()
	createdAt := time.Now().UTC().Unix()
	msg := m.Message

	feedBack := Feedback{
		SendId: sendId,
		Catalog: cataLog,
		Message: msg,
		CreateTime: createdAt,
		Status: handleStatus,
	}

	_, err = o.Insert(&feedBack)

	return  err
}

// GetFeedbackById get feedback by Id and returns feedback
func GetFeedbackById(id int) (v *Feedback, err error) {
	o := orm.NewOrm()
	v = &Feedback{Id: id}
	if err = o.QueryTable(new(Feedback)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateFeedbackById update feedback by Id
func UpdateFeedbackById(m *Feedback)  (err error) {
	o := orm.NewOrm()
	v := Feedback{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		status := m.Status
		*m = v
		m.Status = status
		var num int64
		if num, err = o.Update(m); err == nil {
			logs.Info("Number of records updated in database: ", num)
		}
	}
	return
}

// DeleteFeedback deletes feedback by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFeedback(id int) (err error) {
	o := orm.NewOrm()
	v := Feedback{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Feedback{Id: id}); err == nil {
			logs.Info("Number of records deleted in database: ", num)
		}
	}
	return
}

// GetUserFeedBack get user feedback and returns user feedback array
func GetUserFeedBack() (v []Feedback, err error){
	o := orm.NewOrm()
	var userFeedback [] Feedback
	if _, errQuery := o.QueryTable(new(Feedback)).OrderBy("create_time").Limit(0, 100).All(&userFeedback); errQuery == nil {
		return userFeedback,nil
	} else {
		return userFeedback, err
	}
}