package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type UserEmail struct {
	Id 			  int    `json:"id, omitempty"`
	SendId 		  int    `json:"send_id"`
	RecvId        int    `json:"recv_id"`
	Subject       string `json:"subject"`
	Catalog       int    `json:"catalog"`
	Address       string `json:"address"`
	Message       string `json:"message"`
	CreateTime    int    `json:"create_time, omitempty"`
	Status        int    `json:"status"`
}

type UserRequest struct {
	Page      int     `json:"page"`
	Pagesize  int     `json:"page_size"`
	Status    int     `json:"status"`
	Token     string  `json:"token"`
}

func init() {
	orm.RegisterModel(new(UserEmail))
}

const(
	MsgStatus = 1
)

func (u *UserEmail) TableName() string {
	return TableName("user_mail")
}

// GetUserMessageById get feedback by Id and returns feedback
func GetUserMessageById(id int) (v *UserEmail, err error){
	o := orm.NewOrm()
	v = &UserEmail{Id: id}
	if err = o.QueryTable(new(UserEmail)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserMessageByRecvIdAndStatus get user message by recv_id , status and returns user message array
func GetUserMessageByRecvIdAndStatus(recv_id int, status int, n int, m int) (v []UserEmail, err error){
	o := orm.NewOrm()
	var userMessage [] UserEmail
	if _, errQuery := o.QueryTable(new(UserEmail)).Filter("recv_id", recv_id).Filter("status", status).Limit(n, m).All(&userMessage); errQuery == nil {
		return userMessage, nil
	} else {
		return userMessage, errQuery
	}
}

// UpdateUserMessageById update user meaasge by Id
func UpdateUserMessageById(m *UserEmail)  (err error) {
	o := orm.NewOrm()
	v := UserEmail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		status := m.Status
		*m = v
		m.Status = status
		if num, err = o.Update(m); err == nil {
			logs.Info("Number of records updated in database: ", num)
		}
	}
	return
}

// DeleteUserMessage deletes user message by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserMessage(id int) (err error) {
	o := orm.NewOrm()
	v := UserEmail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserEmail{Id: id}); err == nil {
			logs.Info("Number of records deleted in database: ", num)
		}
	}
	return
}