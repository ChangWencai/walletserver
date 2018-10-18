package controllers

import (
	"wadmin/models"
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type FeedbackController struct {
	BaseController
}


func (self *FeedbackController) List() {
	self.Data["pageTitle"] = "用户反馈管理"
	self.display()
}

func (self *FeedbackController) AjaxDel() {

	feedbackId, _ := self.GetInt("id")
	status := strings.TrimSpace(self.GetString("status"))

	feedbackStatus := 0
	if status == "enable" {
		feedbackStatus = 1
	}
	fbObj, _ := models.GetFeedbackById(feedbackId)
	fbObj.Status = feedbackStatus
	fbObj.Id = feedbackId

	if err := fbObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *FeedbackController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	sendId, _ := self.GetInt("send_id")

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if sendId != 0 {
		filters = append(filters, "send_id__icontains", sendId)
	}
	result, count := models.GetFeedbackList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["send_id"] = v.SendId
		row["catalog"] = v.Catalog
		row["message"] = v.Message
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		if 1 == v.Status {
			row["status"] = "正常"
		} else {
			row["status"] = "冻结"
		}
		//row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}