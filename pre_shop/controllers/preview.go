package controllers

import (
	"github.com/astaxie/beego"
)

type PreviewController struct {
	beego.Controller
}

func (c *PreviewController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "preview.tpl"
}