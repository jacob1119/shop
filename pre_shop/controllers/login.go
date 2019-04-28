package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"shop/pre_shop/models/class"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "http://graduate.hjtvsl.cn:8080/"
	c.Data["Email"] = "1135030879@qq.com"
	c.TplName = "login.tpl"
}

func (c *LoginController) Login() {

	username := c.GetString("username") // login.html中传过来的数据，这个GetInt返回两个值
	password := c.GetString("password")

	// 参数验证
	valid := validation.Validation{}
	valid.Required(username, "username") // 校验是否为空值
	valid.Required(password, "password")

	hash := md5.Sum([]byte(password))
	password = fmt.Sprintf("%x", hash)
	fmt.Println(username,password)
	// 操作数据库，查看手机号，邮箱，用户名是否重复
	var user class.User
	err := orm.NewOrm().QueryTable("user").Filter("Username", username).Filter("Password",password).Limit(1).One(&user)
	if err == nil {

		// 设置session
		c.SetSession(username,username)


		u := c.GetSession(username)
		fmt.Sprintf("%x", u)
		c.Data["Tips"] = "尊敬的"+ u.(string) +"您好，您已经成功登录"
		c.TplName = "welcome.tpl"
		return
	} else {
		c.Data["Tips"] = "账号或用户名错误，请重新登录或者注册账号"
		c.TplName = "welcome.tpl"
		return
	}
}
