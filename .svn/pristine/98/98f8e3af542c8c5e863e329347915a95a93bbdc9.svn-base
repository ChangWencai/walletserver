package models

import (
	"errors"
	"sort"
	"crypto/md5"
	"reflect"
	"strings"
	"time"
	"fmt"
	"wapi/utils"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type User struct {
	Id 				int `json:"id"`
	UserName 		string `json:"username"`
	Password 		string `json:"password"`
	Avatar 			string `json:"avatar, omitempty"`
	Salt 			string `json:"salt"`
	Phone 			string `json:"phone"`
	Introduce 		string `json:"introduce, omitempty"`
	Token 			string `json:"token"`
	Gender 			int `json:"gender"`  // 0:Male, 1: Female, 2: undefined
	Age 			int `json:"age"`
	Address 		string `json:"address"`
	Email 			string `json:"email"`
	LastLoginTime 	int64 `json:"last_login"`
	Status 			int `json:"status"`// 0: enabled, 1:disabled
	CreatedAt 		int64 `json:"created_at"`
	CreateIp 		string `json:"created_ip"`
	UpdatedAt 		int64 `json:"updated_at"`
	Coin 			int `json:"coin"`
}

type ProtoUser struct {
	UserName 		string `json:"username"`
	Password 		string `json:"password"`
	Avatar 			string `json:"avatar, omitempty"`
	Phone 			string `json:"phone"`
	Introduce 		string `json:"introduce, omitempty"`
	Gender 			int    `json:"gender"`  // 0:Male, 1: Female, 2: undefined
	Age 			int    `json:"age"`
	Address 		string `json:"address"`
	Email 			string `json:"email"`
	Coin 			int    `json:"coin"`
	Code            string `json:"code"`
}

type UpdatePwd struct {
	OldPwd    string  `json:"old_pwd, omitempty"`
	NewPwd    string  `json:"new_pwd, omitempty"`
	Phone     string  `json:"phone, omitempty"`
	Code      string  `json:"code, omitempty"`
}

type ForgetPwd struct {
	NewPwd    string  `json:"new_pwd, omitempty"`
	Phone     string  `json:"phone, omitempty"`
	Code      string  `json:"code, omitempty"`
}

type Verify struct {
	Id          int64    `json:"id, omitempty"`
	Phone       string   `json:"phone, omitempty"`
	VerifyCode  string   `json:"verify_code, omitempty"`
	Time        int64    `json:"time, omitempty"`
}

type NickName struct {
	Token    string  `json:"token, omitempty"`
	NickName string  `json:"nick_name, omitempty"`
}

type PhoneMsg struct {
	Phone  string  `json:"phone"`
}

func (u *User) TableName() string {
	return TableName("user")
}

func (u *Verify) TableName() string {
	return TableName("verify")
}

func init() {
	orm.RegisterModel(new(Verify), new(User))
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测用户是否存在
func CheckUserId(userId int) bool {
	exist := Users().Filter("Id", userId).Exist()
	return exist
}

// 检测用户是否存在
func CheckUserName(username string) bool {
	exist := Users().Filter("UserName", username).Exist()
	return exist
}

// 检测用户是否存在
func CheckPhone(phone string) bool {
	exist := Users().Filter("Phone", phone).Exist()
	return exist
}

// 检测用户是否存在
func CheckUserIdAndToken(userId int, token string) bool {
	exist := Users().Filter("Id", userId).Filter("Token", token).Exist()
	return exist
}


// 检测用户是否存在
func CheckEmail(email string) bool {
	exist := Users().Filter("Email", email).Exist()
	return exist
}

// CheckPass compare input password.
func (u *User) CheckPassword(password string) (ok bool, err error) {
	hash, err := utils.GeneratePassHash(password, u.Salt)
	if err != nil {
		return false, err
	}

	return u.Password == hash, nil
}

// 根据用户ID获取用户
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 根据用户ID获取用户
func GetUserByPhone(phone string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Phone: phone}
	if err = o.QueryTable(new(User)).Filter("Phone", phone).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetUserByToken get user by token and returns errors if
// the record to be got doesn't exist
func GetUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(phone string, password string) (bool, *User) {
	o := orm.NewOrm()
	user, err := GetUserByPhone(phone)
	if err != nil {
		fmt.Println()
		return false, nil
	}
	//passwordHash, err := utils.GeneratePassHash(password, user.Salt)
	if err != nil {
		return false, nil
	}
	err = o.QueryTable(user).Filter("phone", phone).Filter("password", password).One(user)
	return err != orm.ErrNoRows, user

}

func AddUser(m *ProtoUser, createIp string) (*User, error) {
	o := orm.NewOrm()
	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}

	//passwordHash, err := utils.GeneratePassHash(m.Password, salt)
	//if err != nil {
	//	return nil, err
	//}
	CreatedAt := time.Now().UTC().Unix()
	UpdatedAt := CreatedAt
	LastLogin := CreatedAt
	Status := 0

	et := utils.EasyToken{
		Phone: m.Phone,
		Uid: 0,
		Expires:  time.Now().Unix() + 2 * 3600,
	}
	token, err := et.GetToken()
	user := User{
		UserName:m.UserName,
		Password: m.Password,
		Avatar:m.Avatar,
		Salt:salt,
		Phone:m.Phone,
		Introduce:m.Introduce,
		Token:token,
		Gender:m.Gender,
		Age:m.Age,
		Address:m.Address,
		Email:m.Email,
		LastLoginTime:LastLogin,
		Status:Status,
		CreatedAt:CreatedAt,
		CreateIp:createIp,
		UpdatedAt:UpdatedAt,
		Coin:0,
	}
	_, err = o.Insert(&user)
	if err == nil{
		return &user, err
	}

	return nil, err
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

// UpdateDevice updates User by DeviceCount and returns error if
// the record to be updated doesn't exist
func UpdateUserDeviceCount(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.Coin += 1
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// updates User's Token and returns error if
// the record to be updated doesn't exist
func UpdateUserToken(m *User, token string) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.Token = token
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// updates User's LastLogin and returns error if
// the record to be updated doesn't exist
func UpdateUserLastLogin(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	lastLogin := time.Now().UTC().Unix()
	m.LastLoginTime = lastLogin
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			logs.Info("Number of records deleted in database:", num)
		}
	}
	return
}


func GetUserName(id int) string {
	var err error
	var username string

	err = utils.GetCache("GetUsername.id."+fmt.Sprintf("%d", id), &username)
	if err != nil {
		cacheExpire, _ := beego.AppConfig.Int("cache_expire")
		var user User
		o := orm.NewOrm()
		o.QueryTable(TableName("user")).Filter("Id", id).One(&user, "username")
		username = user.UserName
		utils.SetCache("GetRealname.id."+fmt.Sprintf("%d", id), username, cacheExpire)
	}
	return username
}

func Send_alidayu(phone, sms_param string) (string, error) {
	sms_url   := beego.AppConfig.String("sms_url")
	sms_type  := beego.AppConfig.String("sms_type")
	sms_free_sign_name := beego.AppConfig.String("sms_free_sign_name")
	sms_template_code := beego.AppConfig.String("sms_template_code")
	sms_AppKey := beego.AppConfig.String("sms_AppKey")
	sms_AppSecret := beego.AppConfig.String("sms_AppSecret")

	req := httplib.Post(sms_url)

	m := map[string]string{
		"app_key":                     sms_AppKey,
		"timestamp":                   time.Now().Format("2006-01-02 15:04:05"),
		"v":                           "2.0",
		"method":                      "alibaba.aliqin.fc.sms.num.send",
		"partner_id":                  "top-apitools",
		"format":                      "json",
		"sms_type":                    sms_type,
		"rec_num":                     phone,
		"sms_free_sign_name":          sms_free_sign_name,
		"sms_template_code":           sms_template_code,
		"force_sensitive_param_fuzzy": "true",
		"sign_method":                 "md5",
		"sms_param":                   sms_param,
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	singString := sms_AppSecret
	for _, k := range keys {
		req.Param(k, m[k])
		singString += k + m[k]
	}
	singString += sms_AppSecret

	signByte := md5.Sum([]byte(singString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	req.Param("sign", sign)

	result, err := req.String()
	return result, err
}