package controllers

import (
	"wadmin/models"
	"time"
	"strings"
	"github.com/astaxie/beego"
)

type UserMailController struct {
	BaseController
}

func (self *UserMailController) List() {
	self.Data["pageTitle"] = "消息管理"
	self.display()
}

func (self *UserMailController) Add() {
	self.Data["pageTitle"] = "新增消息"
	self.display()
}

func (self *UserMailController) Edit() {
	self.Data["pageTitle"] = "编辑消息"

	id, _ := self.GetInt("id", 0)
	messageObj, _ := models.GetUserMailById(id)
	row := make(map[string]interface{})
	row["id"] = messageObj.Id
	row["send_id"] = messageObj.SendId
	row["recv_id"] = messageObj.RecvId
	row["subject"] = messageObj.Subject
	row["catalog"] = messageObj.Catalog
	row["address"] = messageObj.Address
	row["message"] = messageObj.Message
	row["create_time"] = messageObj.CreateTime
	row["status"] = messageObj.Status
	self.Data["user_mail"] = row

	self.display()
}

func (self *UserMailController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	recvId, _ := self.GetInt("recvId")
	self.pageSize = limit

	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if recvId != 0 {
		filters = append(filters, "recv_id__icontains", recvId)
	}

	result, count := models.GetUserMailList(page, self.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["send_id"] = "admin"
		userMailObj, errUser := models.GetUserById(v.RecvId)
		if errUser == nil && "" == userMailObj.UserName {
			row["recv_id"] = userMailObj.Phone
		} else if errUser == nil && "" != userMailObj.UserName{
			row["recv_id"] = userMailObj.UserName
		} else  if errUser != nil {
			row["recv_id"] = v.RecvId
		}
		row["subject"] = v.Subject
		row["address"] = v.Address
		row["message"] = v.Message
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		//区分消息种类
		if 0 == v.Catalog {
			row["catalog"] = "全部"
		} else if 1 == v.Catalog {
			row["catalog"] = "系统"
		} else if 2 == v.Catalog {
			row["catalog"] = "资讯"
		}
		//区分消息状态
		if 1 == v.Status {
			row["status"] = "正常"
		} else {
			row["status"] = "冻结"
		}
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *UserMailController) AjaxSave() {
	userId, _ := self.GetInt("id")
	if userId == 0 {
		mailObj := new(models.UserMail)
		mailObj.SendId, _ = self.GetInt("send_id")
		mailObj.RecvId, _ = self.GetInt("recv_id")
		mailObj.Subject = strings.TrimSpace(self.GetString("subject"))
		mailObj.Catalog, _ = self.GetInt("catalog")
		mailObj.Address = strings.TrimSpace(self.GetString("address"))
		mailObj.Message = strings.TrimSpace(self.GetString("message"))
		mailObj.CreateTime = time.Now().Unix()
		mailObj.Status = 1
		if _, err := models.AddUserMail(mailObj); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("新增成功", MSG_OK)
	}

	// 修改
	userMailUpdate, _ := models.GetUserMailById(userId)
	userMailUpdate.SendId, _ = self.GetInt("send_id")
	userMailUpdate.RecvId, _ = self.GetInt("recv_id")
	userMailUpdate.Subject    = strings.TrimSpace(self.GetString("subject"))
	userMailUpdate.Address = strings.TrimSpace(self.GetString("address"))
	userMailUpdate.Message = strings.TrimSpace(self.GetString("message"))
	userMailUpdate.Status, _ = self.GetInt("status")
	//userMailUpdate.Catalog, _ = self.GetInt("catalog")
	if err := userMailUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("修改成功", MSG_OK)
}

func (self *UserMailController) AjaxDel() {
	userMailId, _ := self.GetInt("id")
	userMailObj, _ := models.GetUserMailById(userMailId)
	userMailObj.Status = 0
	userMailObj.Id = userMailId

	if err := userMailObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}