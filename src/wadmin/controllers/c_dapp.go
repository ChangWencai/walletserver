package controllers

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"wadmin/models"
	"fmt"
)

type DappController struct {
	BaseController
}

func (self *DappController) List() {
	self.Data["pageTitle"] = "DApp设置"
	self.display()
}

func (self *DappController) Add() {
	self.Data["pageTitle"] = "新增DApp"
	self.display()
}

func (self *DappController) Edit() {

	self.Data["pageTitle"] = "编辑DApp"

	id, _ := self.GetInt("id", 0)
	dappObj, _ := models.GetDappById(id)
	row := make(map[string]interface{})
	row["id"] = dappObj.Id
	row["user_id"] = dappObj.UserId
	row["dapp_name"] = dappObj.DappName
	row["dapp_host"] = dappObj.DappHost
	row["dapp_author"] = dappObj.DappAuthor
	row["dapp_img"] = dappObj.DappImg
	row["description"] = dappObj.Description
	row["catalog"] = dappObj.Catalog
	row["status"] = dappObj.Status
	row["popularity"] = dappObj.Popularity
	row["create_time"] = dappObj.CreateTime
	row["update_time"] = dappObj.UpdateTime
	self.Data["dapp"] = row
    
	self.display()
}

func (self *DappController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}
	dappName := strings.TrimSpace(self.GetString("dappName"))
	self.pageSize = limit

	//查询条件
	filters := make([]interface{}, 0)
	if dappName != "" {
		filters = append(filters, "dapp_name__icontains", dappName)
	}
	result, count := models.GetDappList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["dapp_name"] = v.DappName
		row["dapp_host"] = v.DappHost
		row["dapp_author"] = v.DappAuthor
		row["dapp_img"] = v.DappImg
		row["description"] = v.Description
		//区分Dapp种类
		if 1 == v.Catalog {
			row["catalog"] = "banner"
		} else if 2 == v.Catalog {
			row["catalog"] = "游戏"
		} else if 3 == v.Catalog {
			row["catalog"] = "交易所"
		} else if 4 == v.Catalog {
			row["catalog"] = "冒险"
		} else if 5 == v.Catalog {
			row["catalog"] = "其它"
		}
		//区分Dapp状态
		if 1 == v.Status {
			row["status"] = "上架"
		} else {
			row["status"] = "下架"
		}
		row["popularity"] = v.Popularity
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *DappController) AjaxSave() {
	dappId, _ := self.GetInt("id")
	if dappId == 0 {
		dappObj := new(models.Dapp)
		dappObj.DappName = strings.TrimSpace(self.GetString("dapp_name"))
		dappObj.DappHost = strings.TrimSpace(self.GetString("dapp_host"))
		dappObj.DappAuthor = strings.TrimSpace(self.GetString("dapp_author"))
		dappObj.DappImg = strings.TrimSpace(self.GetString("dapp_img"))
		dappObj.Description = strings.TrimSpace(self.GetString("description"))
		dappObj.Catalog, _ =  self.GetInt("catalog")
		dappObj.Popularity, _ =  self.GetInt("popularity")
		dappObj.UserId = self.userId
		dappObj.CreateTime = time.Now().Unix()
		dappObj.UpdateTime = time.Now().Unix()
		dappObj.Status = 1
		if _, err := models.AddDapp(dappObj); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

    // 修改
	dappUpdate, _ := models.GetDappById(dappId)
	dappUpdate.DappName = strings.TrimSpace(self.GetString("dapp_name"))
	dappUpdate.DappHost = strings.TrimSpace(self.GetString("dapp_host"))
	dappUpdate.DappAuthor = strings.TrimSpace(self.GetString("dapp_author"))
	dappUpdate.DappImg = strings.TrimSpace(self.GetString("dapp_img"))
	dappUpdate.Description = strings.TrimSpace(self.GetString("description"))
	dappUpdate.Status, _ = self.GetInt("status")
	dappUpdate.Popularity, _ =  self.GetInt("popularity")
	dappUpdate.UpdateTime = time.Now().Unix()

	if err := dappUpdate.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *DappController) AjaxDel() {
	dappId, _ := self.GetInt("id")
	dappObj, _ := models.GetDappById(dappId)
	dappObj.UpdateTime = time.Now().Unix()
	fmt.Println("dappObj.Status  = ", dappObj.Status)
	if 0 == dappObj.Status {
		dappObj.Status = 1
	} else if 1 == dappObj.Status {
		dappObj.Status = 0
	}
	dappObj.Id = dappId

	if err := dappObj.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}
