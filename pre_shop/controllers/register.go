package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"shop/pre_shop/models/class"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.Data["Website"] = "http://graduate.hjtvsl.cn:8080/"
	c.Data["Email"] = "1135030879@qq.com"
	c.TplName = "register.tpl"
}

func (c *RegisterController) Register() {
	username := c.GetString("username") // login.html中传过来的数据，这个GetInt返回两个值
	password := c.GetString("password")
	phone := c.GetString("phone")
	email := c.GetString("email")
	realName := c.GetString("real_name")

	// 参数验证
	valid := validation.Validation{}
	valid.Required(username, "username") // 校验是否为空值
	valid.Required(password, "password")
	valid.Required(phone, "phone")
	valid.Required(email, "email")
	valid.Required(realName, "real_name")

	// 操作数据库，查看手机号，邮箱，用户名是否重复
	var user class.User
	err := orm.NewOrm().QueryTable("user").Filter("Username", username).Limit(1).One(&user)
	if err == nil {
		c.Data["Tips"] = "用户名已经存在"
		c.TplName = "bad.tpl"
		return
	}

	err = orm.NewOrm().QueryTable("user").Filter("Phone", phone).Limit(1).One(&user)
	if err == nil {
		c.Data["Tips"] = "手机号已经存在"
		c.TplName = "bad.tpl"
		return
	}

	err = orm.NewOrm().QueryTable("user").Filter("Email", email).Limit(1).One(&user)
	if err == nil {
		c.Data["Tips"] = "邮箱已经存在"
		c.TplName = "bad.tpl"
		return
	}

	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.Data["Tips"] = "注册失败" + valid.Errors[0].Key + valid.Errors[0].Message + "请重新填写"
		c.TplName = "bad.tpl"
		return
	}

	hash := md5.Sum([]byte(password))
	password = fmt.Sprintf("%x", hash)
	timeUnix := time.Now().Unix()

	u := &class.User{
		Username:    username,
		Password:    password,
		Real_name:   realName,
		Phone:       phone,
		Email:       email,
		Add_time:    timeUnix,
		Modify_time: timeUnix,
		Active:      0,
	}

	err = u.Create()
	if err != nil {
		fmt.Println(err)
		c.Data["Tips"] = "写入数据库失败"
		c.TplName = "bad.tpl"
		return
	}
	c.Data["Tips"] = "注册成功"
	c.TplName = "welcome.tpl"
}
