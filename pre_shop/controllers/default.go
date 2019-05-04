package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()

	var category []orm.ParamsList
	_, err := o.QueryTable("goods_category").ValuesList(&category, "id", "name")

	if err == nil {
		//c.Data["Cat"] = category
		//fmt.Println("aaaaerror")
	}

	var cat [20] interface{}
	for index,value := range category {
		cat[index+1] = value[1]
	}

	c.Data["Cat"] = cat



	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}