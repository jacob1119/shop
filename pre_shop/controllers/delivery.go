package controllers

import (
	"github.com/astaxie/beego"
)

type DeliveryController struct {
	beego.Controller
}

func (c *DeliveryController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "delivery.tpl"




}