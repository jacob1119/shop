package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
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
	c.Data["Website"] = "http://graduate.hjtvsl.cn:8080/"
	c.Data["Email"] = "1135030879@qq.com"
	c.TplName = "index.tpl"
}
