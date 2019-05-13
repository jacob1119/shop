package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	o := orm.NewOrm()

	var category []orm.ParamsList
	_, err := o.QueryTable("goods_category").ValuesList(&category, "id", "name")

	if err == nil {
		fmt.Println("aaaaerror")
	}

	var cat [20] interface{}
	for index, value := range category {
		cat[index+1] = value[1]
	}

	// 取出我的发售
	var lists []orm.ParamsList
	num, err := orm.NewOrm().Raw("select * from goods order by add_time desc limit 0,8  ").ValuesList(&lists)

	var g [4]class.Goods
	var g2 [4]class.Goods
	var i int64
	for i = 0; i < num; i++ {

		id, _ := strconv.Atoi(lists[i][0].(string))
		category, _ := strconv.Atoi(lists[i][6].(string))
		deal, _ := strconv.Atoi(lists[i][7].(string))
		atime, _ := strconv.ParseInt(lists[i][8].(string), 10, 64)
		mtime, _ := strconv.ParseInt(lists[i][9].(string), 10, 64)

		var imgUrl []orm.ParamsList
		_, err = orm.NewOrm().Raw("select img_url from image where goods_id= ? ", id).ValuesList(&imgUrl)
		url := imgUrl[0][0].(string)
		var categoryName string
		categoryName = cat[category].(string)

		goo := class.Goods{
			Id:           id,
			Title:        lists[i][1].(string),
			Cause:        lists[i][3].(string),
			Price:        lists[i][4].(string),
			Desc:         lists[i][5].(string),
			Category:     category,
			Deal:         deal,
			Add_time:     atime,
			Modify_time:  mtime,
			Username:     lists[i][10].(string),
			Origin_price: lists[i][11].(string),
			Image:        url,
			CategoryName: categoryName,
		}

		if i < 4 {
			g[i] = goo
		} else {
			g2[i-4] = goo
		}


	}
	//fmt.Println(g)

	c.Data["Goods1"] = g
	c.Data["Goods2"] = g2
	c.Data["Cat"] = cat
	c.Data["Website"] = "http://graduate.hjtvsl.cn:8080/"
	c.Data["Email"] = "1135030879@qq.com"
	c.TplName = "index.tpl"
}
