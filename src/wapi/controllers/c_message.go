package controllers

import (
	"strconv"
	"encoding/json"
	"wapi/models"
	"wapi/utils"
)

// MessageController operations for User message
type MessageController struct {
	BaseController
}

// URLMapping ...
func (c *MessageController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetOne ...
// @Title GetOne
// @Description get user message by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.UserEmail
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MessageController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Ctx.ResponseWriter.WriteHeader(11011)
		c.Return(errId)
		return
	}

	v, _ := models.GetUserMessageById(id)
	if v == nil {
		c.Return(errDataBaseSelect)
	} else {
		c.Return(&Response{0, 0, "ok", v})
	}
}

// @Title Update
// @Description update user message
// @Param	id		path 	string	true  "The id you want to update"
// @Success 200 {object} models.UserEmail
// @Failure 403 :id is empty
// @router /:id [put]
func (c *MessageController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Ctx.ResponseWriter.WriteHeader(11011)
		c.Return(errId)
		return
	}

	v := models.UserEmail{Id: id, Status:models.MsgStatus}
	if err := models.UpdateUserMessageById(&v); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseUpdate)
	}
}

// @Title Delete
// @Description delete user message
// @Param	id   path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty or is not valid
// @router /:id [delete]
func (c *MessageController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	if err := models.DeleteUserMessage(id); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseDelete)
	}
}

// @Title UserMessage
// @Description  users message rownum
// @Param	body	body 	models.UserRequest	 true	"body for user content"
// @Success 200 {array}  []models.UserEmail
// @router / [post]
func (c *MessageController) Post() {
	var v models.UserRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Return(errJson)
		return
	}

	page     := v.Page
	pageSize := v.Pagesize
	status   := v.Status
	token    := v.Token

	et := utils.EasyToken{}
	validation, _ := et.ValidateToken(token)
	if !validation {
		c.Return(errExpired)
		return
	}

	if page <= 0 || pageSize <= 0{
		c.Return(errPage)
		return
	}

	yes, user := models.GetUserByToken(token)
	if !yes {
		c.Return(errToken)
		return
	}

	userMessage, errQuery := models.GetUserMessageByRecvIdAndStatus(user.Id, status, (page - 1) * pageSize, pageSize)
	if errQuery == nil {
		c.Return(&Response{0, 0, "ok", userMessage})
	} else {
		c.Return(errDataBaseSelect)
	}
}