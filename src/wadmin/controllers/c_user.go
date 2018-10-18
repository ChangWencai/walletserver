package controllers

import (
	"time"
	"strings"
	"wadmin/models"
	"github.com/astaxie/beego"
	"fmt"
)

type UserController struct {
	BaseController
}

func (self *UserController) List() {
	self.Data["pageTitle"] = "用户管理"
	self.display()
}

func (self *UserController) Edit() {
	self.Data["pageTitle"] = "编辑用户"

	id, _ := self.GetInt("id", 0)
	userObj, _ := models.GetUserById(id)
	row := make(map[string]interface{})
	row["id"] = userObj.Id
	row["user_name"] = userObj.UserName
	row["phone"] = userObj.Phone
	row["email"] = userObj.Email
	self.Data["user"] = row
	self.display()
}

func (self *UserController) AjaxDel() {
	userId, _ := self.GetInt("id")
	status := strings.TrimSpace(self.GetString("status"))
	User_status := 1
	if status == "enable" {
		User_status = 0
	}
	userObj, _ := models.GetUserById(userId)
	userObj.UpdatedAt = time.Now().Unix()

	userObj.Status = User_status
	userObj.Id = userId

	if err := userObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *UserController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	userName := strings.TrimSpace(self.GetString("userName"))
	self.pageSize = limit

	//查询条件
	filters := make([]interface{}, 0)
	if userName != "" {
		filters = append(filters, "user_name__icontains", userName)
	}
	result, count := models.GetUserList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["user_name"] = v.UserName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["create_time"] = beego.Date(time.Unix(v.CreatedAt, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdatedAt, 0), "Y-m-d H:i:s")
		if 1 == v.Status {
			row["status"] = "禁用"
		} else if 0 == v.Status {
			row["status"] = "启用"
		}
		//row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *UserController) Modify() {
	fmt.Println("Modify")
	self.Data["pageTitle"] = "资料修改"
	id := self.userId
	userObj, _ := models.GetUserById(id)
	row := make(map[string]interface{})
	row["id"] = userObj.Id
	row["user_name"] = userObj.UserName
	row["phone"] = userObj.Phone
	row["email"] = userObj.Email
	self.Data["user"] = row
	self.display()
}

func (self *UserController) AjaxModify() {
	userId, _ := self.GetInt("id")
	userObj, _ := models.GetUserById(userId)
	//修改
	userObj.Id = userId
	userObj.UpdatedAt = time.Now().Unix()
	userObj.UserName = strings.TrimSpace(self.GetString("user_name"))
	userObj.Phone = strings.TrimSpace(self.GetString("phone"))
	userObj.Email = strings.TrimSpace(self.GetString("email"))
	userObj.UpdatedAt = time.Now().Unix()
	//userObj.Status = 1

	if err := userObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}