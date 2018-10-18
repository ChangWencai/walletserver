package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/orm"
	"wapi/models"
	"time"
	"fmt"
)

type BaseController struct {
	beego.Controller
}

// Controller Response is controller error info struct.
type Response struct {
	Status         int         `json:"status"`
	ErrorCode      int         `json:"error_code"`
	ErrorMessage   string      `json:"error_message"`
	Data           interface{} `json:"data"`
}

func (c *BaseController) GetClientIp() string {
	s := c.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

func (c *BaseController) Return(value *Response) {
	c.Data["json"] = value
	c.ServeJSON()
}

func (c *BaseController) VerifyCodeTime(phone string, verifyCode string) *Response {
	o := orm.NewOrm()

	value := models.Verify{Phone: phone}
	errVerify := o.Read(&value, "phone")
	if errVerify != nil {
		fmt.Println("errVerify = ", errVerify)
		return errDataBaseSelect
	}

	now := time.Now().Unix()
	if now > value.Time {
		return errVerificationCode
	}

	if value.VerifyCode != verifyCode {
		return errCode
	}

	return nil
}