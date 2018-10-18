package controllers

import (
	//"errors"
	//"regexp"
	//"strings"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
	//"github.com/dgrijalva/jwt-go"
)

// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

// UserData definition.
type UserSuccessLoginData struct {
	Token    string  `json:"token"`
	UserID   int     `json:"user_id"`
	NickName string  `json:"nick_name"`
}

// CreateDevice definition.
type CreateObjectData struct {
	Id int `json:"id"`
}

// GetDevices definition.
type GetDeviceData struct {
	TotalCount int64 `json:"total_count"`
	Devices interface{} `json:"devices"`
}

// GetAirAds definition.
type GetAirAdData struct {
	TotalCount int64 `json:"total_count"`
	AirAds interface{} `json:"airads"`
}

// Predefined controller error/success values.
var (
	successReturn = &Response{200, 0, "ok", "ok"}
	err404 = &Response{404, 404, "找不到网页", "找不到网页"}
	errInputData = &Response{400, 10001, "数据输入错误", "客户端参数错误"}
	errDatabase = &Response{500, 10002, "服务器错误", "数据库操作错误"}
	errUserToken = &Response{500, 10002, "服务器错误", "令牌操作错误"}
	errDupUser = &Response{400, 10003, "用户信息已存在", "数据库记录重复"}
	errNoUser = &Response{400, 10004, "用户信息不存在", "数据库记录不存在"}
	errPass = &Response{400, 10005, "用户信息不存在或密码不正确", "密码不正确"}
	errNoUserOrPass = &Response{400, 10006, "用户不存在或密码不正确", "数据库记录不存在或密码不正确"}
	errNoUserChange = &Response{400, 10007, "用户不存在或数据未改变", "数据库记录不存在或数据未改变"}
	errInvalidUser = &Response{400, 10008, "用户信息不正确", "Session信息不正确"}
	errOpenFile = &Response{500, 10009, "服务器错误", "打开文件出错"}
	errWriteFile = &Response{500, 10010, "服务器错误", "写文件出错"}
	errSystem = &Response{500, 10011, "服务器错误", "操作系统错误"}
	errPermission = &Response{400, 10013, "没有权限", "没有操作权限"}

	errExpired = &Response{11001, 11001, "登录已过期", "验证token过期"}
	errToken = &Response{11002, 11002, "token不存在", "token不存在"}

	errId = &Response{11011, 11011, "输入的ID不合法", "ID只能是正整数"}
	errJson = &Response{11012, 11012, "Json文件解析出错", "Json文件解析出错"}
	errPage = &Response{11013, 11013, "分页数据有误", "分页数据有误"}
	errPhoneExist = &Response{11014, 11014, "手机号码已经注册了", "手机号称已经注册了"}
	errDataBaseSelect = &Response{11015, 11015, "数据库查询错误", "数据库查询错误"}
	errDataBaseUpdate = &Response{11016, 11016, "数据库更新错误", "数据库更新错误"}
	errDataBaseDelete = &Response{11017, 11017, "数据库删除错误", "数据库删除错误"}
	errDataBaseInsert = &Response{11018, 11018, "数据库增加错误", "数据库增加错误"}
	errCode = &Response{11019, 11019, "验证码输入有错", "验证码输入有错"}
	errVerificationCode = &Response{11020, 11020, "输入验证码超时", "输入验证码超时"}
	errUserNotExist = &Response{11021, 11021, "用户不存在", "用户不存在"}
	errOldPwd = &Response{11022, 11022, "输入旧密码错误", "输入旧密码错误"}
	errPhoneNotExist = &Response{11023, 11023, "该手机号码还没有注册", "该手机号码还没有注册"}//11024
)

// BaseController definition.
//type BaseController struct {
//	beego.Controller
//}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *Response) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.Data = ""
	}

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()
	base.StopRun()
}

var sqlOp = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}

// ParseQueryParm parse query parameters.
//   query=col1:op1:val1,col2:op2:val2,...
//   op: one of eq, ne, gt, ge, lt, le
//func (base *BaseController) ParseQueryParameter() (v map[string]string, o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	queryVal := make(map[string]string)
//	queryOp := make(map[string]string)
//
//	query := base.GetString("query")
//	if query == "" {
//		return queryVal, queryOp, nil
//	}
//
//	for _, cond := range strings.Split(query, ",") {
//		kov := strings.Split(cond, ":")
//		if len(kov) != 3 {
//			return queryVal, queryOp, errors.New("Query format != k:o:v")
//		}
//
//		var key string
//		var value string
//		var operator string
//		if !nameRule.MatchString(kov[0]) {
//			return queryVal, queryOp, errors.New("Query key format is wrong")
//		}
//		key = kov[0]
//		if op, ok := sqlOp[kov[1]]; ok {
//			operator = op
//		} else {
//			return queryVal, queryOp, errors.New("Query operator is wrong")
//		}
//		value = strings.Replace(kov[2], "'", "\\'", -1)
//
//		queryVal[key] = value
//		queryOp[key] = operator
//	}
//
//	return queryVal, queryOp, nil
//}

// ParseOrderParameter parse order parameters.
//   order=col1:asc|desc,col2:asc|esc,...
//func (base *BaseController) ParseOrderParameter() (o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	order := make(map[string]string)
//
//	v := base.GetString("order")
//	if v == "" {
//		return order, nil
//	}
//
//	for _, cond := range strings.Split(v, ",") {
//		kv := strings.Split(cond, ":")
//		if len(kv) != 2 {
//			return order, errors.New("Order format != k:v")
//		}
//		if !nameRule.MatchString(kv[0]) {
//			return order, errors.New("Order key format is wrong")
//		}
//		if kv[1] != "asc" && kv[1] != "desc" {
//			return order, errors.New("Order val isn't asc/desc")
//		}
//
//		order[kv[0]] = kv[1]
//	}
//
//	return order, nil
//}

// ParseLimitParameter parse limit parameter.
//   limit=n
//func (base *BaseController) ParseLimitParameter() (l int64, err error) {
//	if v, err := base.GetInt64("limit"); err != nil {
//		return 10, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 10, nil
//	}
//}

// ParseOffsetParameter parse offset parameter.
//   offset=n
//func (base *BaseController) ParseOffsetParameter() (o int64, err error) {
//	if v, err := base.GetInt64("offset"); err != nil {
//		return 0, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 0, nil
//	}
//}

// VerifyForm use validation to verify input parameters.
//func (base *BaseController) VerifyForm(obj interface{}) (err error) {
//	valid := validation.Validation{}
//	ok, err := valid.Valid(obj)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		str := ""
//		for _, err := range valid.Errors {
//			str += err.Key + ":" + err.Message + ";"
//		}
//		return errors.New(str)
//	}
//
//	return nil
//}

// ParseToken parse JWT token in http header.
//func (base *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
//	authString := base.Ctx.Input.Header("Authorization")
//	beego.Debug("AuthString:", authString)
//
//	kv := strings.Split(authString, " ")
//	if len(kv) != 2 || kv[0] != "Bearer" {
//		beego.Error("AuthString invalid:", authString)
//		return nil, errInputData
//	}
//	tokenString := kv[1]
//
//	// Parse token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte("secret"), nil
//	})
//	if err != nil {
//		beego.Error("Parse token:", err)
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				// That's not even a token
//				return nil, errInputData
//			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
//				// Token is either expired or not active yet
//				return nil, errExpired
//			} else {
//				// Couldn't handle this token
//				return nil, errInputData
//			}
//		} else {
//			// Couldn't handle this token
//			return nil, errInputData
//		}
//	}
//	if !token.Valid {
//		beego.Error("Token invalid:", tokenString)
//		return nil, errInputData
//	}
//	beego.Debug("Token:", token)
//
//	return token, nil
//}
