package controllers

import (
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "contact.tpl"
}