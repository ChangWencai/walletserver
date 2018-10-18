package controllers

import (
	"strconv"
	"wapi/models"
	"encoding/json"

)

//  DappController operations for Device
type DappController struct {
	BaseController
}

// URLMapping ...
func (c *DappController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Update
// @Description update the user message
// @Param	id		path 	string	true   "The id you want to update"
// @Success 200 {object} models.Dapp
// @Failure 403 :id is empty
// @router /:id [put]
func (c *DappController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Return(errId)
		return
	}

	v := models.Dapp{Id: id, Status:1}
	if err := models.UpdateDappById(&v); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseUpdate)
	}
}

// GetOne ...
// @Title Get One
// @Description get Device by id
// @Param	id		path 	string	true		"The key for static block"
// @Success 200 {object} models.Dapp
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DappController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDappById(id)
	if err != nil {
		c.Return(errDataBaseSelect)
	} else {
		c.Return(&Response{0, 0, "ok", v})
	}
}

// GetAll ...
// @Title Get All
// @Description get Dapp
// @Param	id	path  int	true "The id you want to get"
// @Success 200 {object} models.Dapp
// @router /:id [get]
func (c *DappController) GetAll() {

}

// @Title DappList
// @Description  get dapp list
// @Param	body	body  models.DataLine 	 true  "The line you want to get"
// @Success 200 {array}  []models.Response
// @router / [post]
func (c *DappController) Post() {
	var v models.DataLine
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Return(errJson)
		return
	}

	line := v.Line
	var  value  models.ResponseDapp

	banner, errBanner := models.GetBanner(line)
	if errBanner != nil {
		c.Return(errDataBaseSelect)
		return
	}
	value.Banner = append(value.Banner, banner...)


	gameGame, errGame := models.GetGame(line)
	if errGame != nil {
		c.Return(errDataBaseSelect)
		return
	}
	objectGame := models.GameList{"游戏", gameGame}
	value.GameList = append(value.GameList, objectGame)


	gameExchange, errExchange := models.GetExchange(line)
	if errExchange != nil {
		c.Return(errDataBaseSelect)
		return
	}
	objectExchange := models.GameList{"交易所", gameExchange}
	value.GameList = append(value.GameList, objectExchange)


	gameDangerous, errDangerous := models.GetDangerous(line)
	if errDangerous != nil {
		c.Return(errDataBaseSelect)
		return
	}
	objectDangerous := models.GameList{"冒险", gameDangerous}
	value.GameList = append(value.GameList, objectDangerous)

	gameOther, errOther := models.GetOther(line)
	if errOther != nil {
		c.Return(errDataBaseSelect)
		return
	}
	objectOther := models.GameList{"其它", gameOther}
	value.GameList = append(value.GameList, objectOther)

	c.Return(&Response{0, 0, "ok", value})
}


// @Title GameList
// @Description  get game list
// @Param	body	body  models.GameName 	 true  "The game name"
// @Success 200 {array}  []models.Games
// @router /find [post]
func (c *DappController) Find() {
	var v models.GameName
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Return(errJson)
		return
	}

	gameList, err := models.GetFind(v.Name)
	if err != nil {
		c.Return(errDataBaseSelect)
	} else {
		c.Return(&Response{0, 0, "ok", gameList})
	}
}

// @Title Delete
// @Description delete dapp
// @Param	id   path 	string	true    "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty or id is not exist
// @router /:id [delete]
func (c *DappController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Return(errId)
		return
	}

	if err := models.DeleteDapp(id); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseDelete)
	}
}