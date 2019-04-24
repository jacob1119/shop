package controllers

import (
	"github.com/astaxie/beego"
)

type WelcomeController struct {
	beego.Controller
}

func (c *WelcomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "welcome.tpl"
}
