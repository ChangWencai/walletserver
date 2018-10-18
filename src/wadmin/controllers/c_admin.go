package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"wadmin/libs"
	"wadmin/models"
)

type AdminController struct {
	BaseController
}

func (self *AdminController) List() {
	self.Data["pageTitle"] = "管理员管理"
	self.display()
	//self.TplName = "admin/list.html"
}

func (self *AdminController) Add() {
	self.Data["pageTitle"] = "新增管理员"

	// 角色
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}

	self.Data["role"] = list

	self.display()
}

func (self *AdminController) Edit() {
	self.Data["pageTitle"] = "编辑管理员"

	id, _ := self.GetInt("id", 0)
	adminObj, _ := models.GetAdminById(id)
	row := make(map[string]interface{})
	row["id"] = adminObj.Id
	row["login_name"] = adminObj.LoginName
	row["real_name"] = adminObj.RealName
	row["phone"] = adminObj.Phone
	row["email"] = adminObj.Email
	row["role_ids"] = adminObj.RoleIds
	self.Data["admin"] = row

	roleIds := strings.Split(adminObj.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(roleIds); i++ {
			roleId, _ := strconv.Atoi(roleIds[i])
			if roleId == v.Id {
				row["checked"] = 1
			}
			fmt.Println(roleIds[i])
		}
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}
	self.Data["role"] = list
	self.display()
}

func (self *AdminController) AjaxSave() {
	adminId, _ := self.GetInt("id")
	if adminId == 0 {
		adminObj := new(models.Admin)
		adminObj.LoginName = strings.TrimSpace(self.GetString("login_name"))
		adminObj.RealName = strings.TrimSpace(self.GetString("real_name"))
		adminObj.Phone = strings.TrimSpace(self.GetString("phone"))
		adminObj.Email = strings.TrimSpace(self.GetString("email"))
		adminObj.RoleIds = strings.TrimSpace(self.GetString("roleids"))
		adminObj.UpdateTime = time.Now().Unix()
		adminObj.UpdateId = self.userId
		adminObj.Status = 1

		// 检查登录名是否已经存在
		_, err := models.GetAdminByName(adminObj.LoginName)

		if err == nil {
			self.ajaxMsg("登录名已经存在", MSG_ERR)
		}
		//新增
		pwd, salt := libs.Password(4, "")
		adminObj.Password = pwd
		adminObj.Salt = salt
		adminObj.CreateTime = time.Now().Unix()
		adminObj.CreateId = self.userId
		if _, err := models.AddAdmin(adminObj); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	adminObj, _ := models.GetAdminById(adminId)
	//修改
	adminObj.Id = adminId
	adminObj.UpdateTime = time.Now().Unix()
	adminObj.UpdateId = self.userId
	adminObj.LoginName = strings.TrimSpace(self.GetString("login_name"))
	adminObj.RealName = strings.TrimSpace(self.GetString("real_name"))
	adminObj.Phone = strings.TrimSpace(self.GetString("phone"))
	adminObj.Email = strings.TrimSpace(self.GetString("email"))
	adminObj.RoleIds = strings.TrimSpace(self.GetString("roleids"))
	adminObj.UpdateTime = time.Now().Unix()
	adminObj.UpdateId = self.userId
	adminObj.Status = 1

	resetPwd, _ := self.GetInt("reset_pwd")
	if resetPwd == 1 {
		pwd, salt := libs.Password(4, "")
		adminObj.Password = pwd
		adminObj.Salt = salt
	}
	if err := adminObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg(strconv.Itoa(resetPwd), MSG_OK)
}

func (self *AdminController) AjaxDel() {

	adminId, _ := self.GetInt("id")
	status := strings.TrimSpace(self.GetString("status"))
	if adminId == 1 {
		self.ajaxMsg("超级管理员不允许操作", MSG_ERR)
	}

	adminStatus := 0
	if status == "enable" {
		adminStatus = 1
	}
	adminObj, _ := models.GetAdminById(adminId)
	adminObj.UpdateTime = time.Now().Unix()
	adminObj.Status = adminStatus
	adminObj.Id = adminId

	if err := adminObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *AdminController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	adminId, _ := self.GetInt("admin_id")

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if adminId != 0 {
		filters = append(filters, "id__icontains", adminId)
	}
	result, count := models.GetAdminList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}


func (self *AdminController) Modify() {
	self.Data["pageTitle"] = "资料修改"
	id := self.userId
	adminObj, _ := models.GetAdminById(id)
	row := make(map[string]interface{})
	row["id"] = adminObj.Id
	row["login_name"] = adminObj.LoginName
	row["real_name"] = adminObj.RealName
	row["phone"] = adminObj.Phone
	row["email"] = adminObj.Email
	self.Data["admin"] = row
	self.display()
}

func (self *AdminController) AjaxModify() {
	adminId, _ := self.GetInt("id")
	adminObj, _ := models.GetAdminById(adminId)
	//修改
	adminObj.Id = adminId
	adminObj.UpdateTime = time.Now().Unix()
	adminObj.UpdateId = self.userId
	adminObj.LoginName = strings.TrimSpace(self.GetString("login_name"))
	adminObj.RealName = strings.TrimSpace(self.GetString("real_name"))
	adminObj.Phone = strings.TrimSpace(self.GetString("phone"))
	adminObj.Email = strings.TrimSpace(self.GetString("email"))

	resetPwd := self.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(self.GetString("password_old"))
		pwdOldMd5 := libs.Md5([]byte(pwdOld + adminObj.Salt))
		if adminObj.Password != pwdOldMd5 {
			self.ajaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(self.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(self.GetString("password_new2"))

		if pwdNew1 != pwdNew2 {
			self.ajaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := libs.Password(4, pwdNew1)
		adminObj.Password = pwd
		adminObj.Salt = salt
	}
	adminObj.UpdateTime = time.Now().Unix()
	adminObj.UpdateId = self.userId
	adminObj.Status = 1

	if err := adminObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
