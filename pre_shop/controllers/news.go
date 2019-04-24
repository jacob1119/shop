package controllers

import (
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "news.tpl"
}
