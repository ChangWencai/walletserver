package controllers

import (
	"wapi/models"
	"encoding/json"
	"strconv"
	"wapi/utils"

)

//  FeedBackController operations for Device
type FeedBackController struct {
	BaseController
}

// URLMapping ...
func (c *FeedBackController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetAll ...
// @Title Get All
// @Description get feedbacks
// @Success 200 {object} models.Feedback
// @Failure 403
// @router / [get]
func (c *FeedBackController) GetAll() {
	userFeedback, err := models.GetUserFeedBack()
	if err == nil {
		c.Return(&Response{0, 0, "ok", userFeedback})
	} else {
		c.Return(errDataBaseSelect)
	}
}

// @Title Update
// @Description update the feedback
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} models.Feedback
// @Failure 403 :id is empty
// @router /:id [put]
func (c *FeedBackController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Return(errId)
		return
	}

	v := models.Feedback{Id: id, Status:models.HandleStatus}
	if err := models.UpdateFeedbackById(&v); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseUpdate)
	}
}

// GetOne ...
// @Title GetOne
// @Description get feedback by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.Feedback
// @Failure 403 :id is empty or id is not valid
// @router /:id [get]
func (c *FeedBackController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Return(errId)
		return
	}

	v, _ := models.GetFeedbackById(id)
	if v == nil {
		c.Return(errDataBaseSelect)
	} else {
		c.Return(&Response{0, 0, "ok", v})
	}
}

// @Title Delete
// @Description delete feedback
// @Param	id   path 	string 	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FeedBackController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if  id <= 0 {
		c.Return(errId)
		return
	}

	if err := models.DeleteFeedback(id); err == nil {
		c.Return(successReturn)
	} else {
		c.Return(errDataBaseDelete)
	}
}

// @Title FeedBack
// @Description  user feedback
// @Param	body	body 	models.FeedbackRequestStruct	 true	"body for user content"
// @Success 200 {array}  []models.Feedback
// @router / [post]
func (c *FeedBackController) Post() {
	var v models.FeedbackRequestStruct
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Return(errJson)
		return
	}

	token := v.Token
	et := utils.EasyToken{}
	validation, _ := et.ValidateToken(token)
	if !validation {
		c.Return(errExpired)
		return
	}

	if  err := models.AddFeedBack(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Return(&Response{0, 0, "ok", "用户反馈成功"})
	} else {
		c.Return(&Response{11030, 11030, "用户反馈失败", err.Error()})
	}
}
