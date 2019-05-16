package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "about.tpl"
}


func (c *AboutController) Contact() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Tips"] = "您好，感谢您的关注。您的反馈信息我会尽快阅读并给您反馈，希望您能持续关注我们的平台。"
	c.TplName = "bad.tpl"
}