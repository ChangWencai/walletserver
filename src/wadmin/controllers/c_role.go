package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"wadmin/models"
)

type RoleController struct {
	BaseController
}

func (self *RoleController) List() {
	self.Data["pageTitle"] = "角色管理"
	self.display()
}

func (self *RoleController) Add() {
	self.Data["zTree"] = true //引入ztreecss
	self.Data["pageTitle"] = "新增角色"
	self.display()
}
func (self *RoleController) Edit() {
	self.Data["zTree"] = true //引入ztreecss
	self.Data["pageTitle"] = "编辑角色"

	id, _ := self.GetInt("id", 0)
	roleObj, _ := models.RoleGetById(id)
	row := make(map[string]interface{})
	row["id"] = roleObj.Id
	row["role_name"] = roleObj.RoleName
	row["detail"] = roleObj.Detail
	self.Data["roleObj"] = row

	//获取选择的树节点
	roleAuth, _ := models.RoleAuthGetById(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	self.Data["auth"] = authId
	fmt.Println(authId)
	self.display()
}

func (self *RoleController) AjaxSave() {
	roleObj := new(models.Role)
	roleObj.RoleName = strings.TrimSpace(self.GetString("role_name"))
	roleObj.Detail = strings.TrimSpace(self.GetString("detail"))
	roleObj.CreateTime = time.Now().Unix()
	roleObj.UpdateTime = time.Now().Unix()
	roleObj.Status = 1
	auths := strings.TrimSpace(self.GetString("nodes_data"))
	roleId, _ := self.GetInt("id")
	if roleId == 0 {
		//新增
		roleObj.CreateTime = time.Now().Unix()
		roleObj.UpdateTime = time.Now().Unix()
		roleObj.CreateId = self.userId
		roleObj.UpdateId = self.userId
		if id, err := models.RoleAdd(roleObj); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			ra := new(models.RoleAuth)
			authsSlice := strings.Split(auths, ",")
			for _, v := range authsSlice {
				aid, _ := strconv.Atoi(v)
				ra.AuthId = aid
				ra.RoleId = id
				models.RoleAuthAdd(ra)
			}
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	roleObj.Id = roleId
	roleObj.UpdateTime = time.Now().Unix()
	roleObj.UpdateId = self.userId
	if err := roleObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		// 删除该角色权限
		models.RoleAuthDelete(roleId)
		ra := new(models.RoleAuth)
		authsSlice := strings.Split(auths, ",")
		for _, v := range authsSlice {
			aid, _ := strconv.Atoi(v)
			ra.AuthId = aid
			ra.RoleId = int64(roleId)
			models.RoleAuthAdd(ra)
		}

	}
	self.ajaxMsg("", MSG_OK)
}

func (self *RoleController) AjaxDel() {

	roleId, _ := self.GetInt("id")
	roleObj, _ := models.RoleGetById(roleId)
	roleObj.Status = 0
	roleObj.Id = roleId
	roleObj.UpdateTime = time.Now().Unix()

	if err := roleObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	// 删除该角色权限
	//models.RoleAuthDelete(roleId)
	self.ajaxMsg("", MSG_OK)
}

func (self *RoleController) Table() {
	fmt.Println("lurunze1")
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	roleName := strings.TrimSpace(self.GetString("roleName"))
	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if roleName != "" {
		filters = append(filters, "role_name__icontains", roleName)
	}
	result, count := models.RoleGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
