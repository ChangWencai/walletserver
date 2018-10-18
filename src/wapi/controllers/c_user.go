package controllers

import (
	"wapi/models"
	"wapi/utils"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
)

// Operations about Users
type UserController struct {
	BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.ProtoUser	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.ProtoUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := utils.CheckNewUserPost(v.Phone, v.Password); errorMessage != "ok"{
			c.Data["json"] = Response{11024, 11024,errorMessage, errorMessage}
			c.ServeJSON()
			return
		}
		if models.CheckPhone(v.Phone){
			c.Return(errPhoneExist)
			return
		}

		response := c.VerifyCodeTime(v.Phone, v.Code)
		if response != nil {
			c.Return(response)
			return
		}

		if user, err := models.AddUser(&v, c.GetClientIp()); err == nil {
			c.Ctx.Output.SetStatus(201)
			var returnData = &UserSuccessLoginData{user.Token, user.Id, user.UserName}
			c.Data["json"] = &Response{0, 0, "ok", returnData}
		} else {
			c.Data["json"] = &Response{1, 1, "用户注册失败", err.Error()}
		}
	} else {
		c.Data["json"] = &Response{1, 1, "用户注册失败", err.Error()}
	}

	c.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 10
	var offset int

	token := c.Ctx.Input.Header("token")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	} else {
		fields = strings.Split("Username,Gender,Age,Address,Email,Token", ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	token := c.Ctx.Input.Header("token")
	//idStr := c.Ctx.Input.Param("id")
	idStr := c.Ctx.Input.Param(":id")
	//token := c.Ctx.Input.Param(":token")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(token)
	if !valido {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if v == nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			c.Data["json"] = successReturn
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = successReturn
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	phone		formData 	string	true		"The phone for login"
// @Param	password	formData 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [POST]
func (c *UserController) Login() {
	var token string
	var Phone = c.GetString("phone")
	var Password = c.GetString("password")

	if errorMessage := utils.CheckPhonePassword(Phone, Password); errorMessage != "ok"{
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403,errorMessage, ""}
		c.ServeJSON()
		return
	}

	if ok, user := models.Login(Phone, Password); ok {
		et := utils.EasyToken{}
		validation, err := et.ValidateToken(user.Token)

		if !validation {
			et = utils.EasyToken{
				Phone: 	user.Phone,
				Uid:    int64(user.Id),
				Expires: time.Now().Unix() + 2 * 3600,
			}
			token, err = et.GetToken()
			if token == "" || err != nil {
				c.Data["json"] = errUserToken
				c.ServeJSON()
				return
			} else {
				models.UpdateUserToken(user, token)
			}
		} else {
			token = user.Token
		}
		models.UpdateUserLastLogin(user)

		var returnData = &UserSuccessLoginData{token, user.Id, user.UserName}
		c.Data["json"] = &Response{0, 0, "ok", returnData}
	} else {
		c.Data["json"] = &errNoUserOrPass
	}
	c.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Param	token		header 	string	true		"token"
// @Success 200 {string} logout success
// @Failure 401 unauthorized
// @router /auth [get]
func (c *UserController) Auth() {
	et := utils.EasyToken{}
	token := strings.TrimSpace(c.Ctx.Request.Header.Get("token"))
	validation, _ := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(11001)
		c.Return(errExpired)
		return
	}

	c.Return(successReturn)
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Return(successReturn)
}

// @Title message verify
// @Description test message verify
// @Param	body	body  models.PhoneMsg  true	 "phone"
// @Success 200 {string} verify success
// @Failure 401 unauthorized
// @router /verify [post]
func (c *UserController) Verify() {
	var v models.PhoneMsg
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Return(errJson)
		return
	}

	o := orm.NewOrm()
	value := models.Verify{Phone: v.Phone}

	err := o.Read(&value, "phone")
	fmt.Println("err = ", err)
	now := time.Now().Unix()
	fmt.Println("phone: ", value.Phone, " time: ", value.Time)

	// 调用第三方接口获取验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	value.Time = now + 60
	value.VerifyCode = vcode

	if  err == nil {
		if _, errUpdate := o.Update(&value); errUpdate != nil {
			c.Return(errDataBaseUpdate)
		} else {
			c.Return(&Response{0, 0, "ok", value.VerifyCode})
		}
	} else {
		_, errInsert := o.Insert(&value)
		if errInsert != nil {
			c.Return(errDataBaseInsert)
		} else {
			c.Return(&Response{0, 0, "ok", value.VerifyCode})
		}
	}
}

// @Title update password
// @Description update user password
// @Param	body	body  models.UpdatePwd  true	 "phone"
// @Success 200 {string} update success
// @Failure 401 unauthorized
// @router /update [post]
func (c *UserController) Update() {
	var userUpdate models.UpdatePwd
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userUpdate); err != nil {
		c.Return(errJson)
		return
	}

	if exist := models.CheckPhone(userUpdate.Phone); !exist {
		c.Return(errPhoneNotExist)
		return
	}

	response := c.VerifyCodeTime(userUpdate.Phone, userUpdate.Code)
	if response != nil {
		c.Return(response)
		return
	}

	o := orm.NewOrm()
	var user  models.User
	user = models.User{Phone: userUpdate.Phone}
	err := o.Read(&user, "Phone")
	if err != nil {
		c.Return(errDataBaseSelect)
		return
	}

	if user.Password != userUpdate.OldPwd {
		c.Return(errOldPwd)
		return
	}

	user.Password = userUpdate.NewPwd
	_, errUpdate := o.Update(&user)
	if errUpdate != nil {
		c.Return(errDataBaseUpdate)
	} else {
		c.Return(successReturn)
	}
}

// @Title forget password
// @Description user forget password
// @Param	body	body  models.ForgetPwd  true	 "phone"
// @Success 200 {string} verify success
// @Failure 401 unauthorized
// @router /forget [post]
func (c *UserController) Forget() {
	var userForget models.ForgetPwd
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userForget); err != nil {
		c.Return(errJson)
		return
	}

	if exist := models.CheckPhone(userForget.Phone); !exist {
		c.Return(errPhoneNotExist)
		return
	}

	response := c.VerifyCodeTime(userForget.Phone, userForget.Code)
	if response != nil {
		c.Return(response)
		return
	}

	o := orm.NewOrm()
	var user  models.User
	user = models.User{Phone: userForget.Phone}
	err := o.Read(&user, "Phone")
	if err != nil {
		c.Return(errDataBaseSelect)
		return
	}

	user.Password = userForget.NewPwd
	_, errUpdate := o.Update(&user)
	if errUpdate != nil {
		c.Return(errDataBaseUpdate)
	} else {
		c.Return(successReturn)
	}
}

// @Title edit nickname
// @Description user edit nickname
// @Param	body	body  models.NickName  true	 "nickname"
// @Success 200 {string} edit success
// @Failure 401 unauthorized
// @router /edit [post]
func (c *UserController) Edit() {
	var userNickName models.NickName
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userNickName); err != nil {
		c.Return(errJson)
		return
	}

	token    := userNickName.Token
	nickName := userNickName.NickName

	et := utils.EasyToken{}
	validation, _ := et.ValidateToken(token)
	if !validation {
		c.Return(errExpired)
		return
	}

	o := orm.NewOrm()
	var user  models.User
	user = models.User{Token: token}
	err := o.Read(&user, "Token")
	if err != nil {
		c.Return(errDataBaseSelect)
		return
	}

	user.UserName = nickName

	_, errUpdate := o.Update(&user)
	if errUpdate != nil {
		c.Return(errDataBaseUpdate)
	} else {
		c.Return(successReturn)
	}
}