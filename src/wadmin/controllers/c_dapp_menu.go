package controllers

import (
	"strings"
	"wadmin/models"
)

type DappMenuController struct {
	BaseController
}

func (self *DappMenuController) List() {
	self.Data["pageTitle"] = "DApp Menu 管理"
	self.display()
}

func (self *DappMenuController) Add() {
	self.Data["pageTitle"] = "新增DApp Menu"
	self.display()
}



func (self *DappMenuController) Edit() {
	self.Data["pageTitle"] = "编辑DApp Menu"

	id, _ := self.GetInt("id", 0)
	dappMenuObj, _ := models.GetDappMenuById(id)
	row := make(map[string]interface{})
	row["id"] = dappMenuObj.Id
	row["catalog"] = dappMenuObj.Catalog
	row["name"] = dappMenuObj.Name
	row["status"] = dappMenuObj.Status
	self.Data["dapp_menu"] = row
	self.display()
}

func (self *DappMenuController) AjaxSave() {
	dappMenuId, _ := self.GetInt("id")
	if dappMenuId == 0 {
		dappMenuObj := new(models.DappMenu)
		dappMenuObj.Catalog, _     = self.GetInt("catalog")
		dappMenuObj.Name = strings.TrimSpace(self.GetString("name"))
		dappMenuObj.Status = 1
		errCheckCatalog := models.CheckDappMenuCatalog(dappMenuObj.Catalog)
		if true == errCheckCatalog {
			self.ajaxMsg("The Catalog is  exist in the wt_dapp_menu table", MSG_ERR)
		}
		if _, err := models.AddDappMenu(dappMenuObj); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	dappMenuObj, _ := models.GetDappMenuById(dappMenuId)
	dappMenuObj.Catalog, _ = self.GetInt("catalog")
	dappMenuObj.Name       = strings.TrimSpace(self.GetString("name"))

	if err := dappMenuObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("更新成功", MSG_OK)
}

func (self *DappMenuController) AjaxDel() {
	dappMenuId, _ := self.GetInt("id")
	dappMenuObj, _ := models.GetDappMenuById(dappMenuId)
	dappMenuObj.Status = 0
	dappMenuObj.Id = dappMenuId

	if err := dappMenuObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *DappMenuController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	catalog, _ := self.GetInt("catalog")
	self.pageSize = limit

	//查询条件
	filters := make([]interface{}, 0)

	if catalog != 0 {
		filters = append(filters, "catalog__icontains", catalog)
	}
	result, count := models.GetDappMenuList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["catalog"] = v.Catalog
		row["name"] = v.Name
		//区分栏目状态
		if 1 == v.Status {
			row["status"] = "使用"
		} else {
			row["status"] = "停用"
		}
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
