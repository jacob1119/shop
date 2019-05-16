package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"strconv"
)

type PreviewController struct {
	beego.Controller
}

func (c *PreviewController) Get() {
	var id int
	id, _ = c.GetInt("id")

	// 取出我的发售
	var g []orm.ParamsList
	num, err := orm.NewOrm().Raw("select * from goods where id= ? ", id).ValuesList(&g)
	if num < 1 {
		c.Data["Tips"] = "Wows  商品走丢了"
		c.TplName = "bad.tpl"
		return
	}
	// 取出分类
	o := orm.NewOrm()
	var category []orm.ParamsList
	_, err = o.QueryTable("goods_category").ValuesList(&category, "id", "name")

	if err == nil {
		//fmt.Println("aaaaerror")
	}

	var cat [20] interface{}
	for index, value := range category {
		cat[index+1] = value[1]
	}

	/*   商品详情   */
	category1, _ := strconv.Atoi(g[0][6].(string))
	deal, _ := strconv.Atoi(g[0][7].(string))
	atime, _ := strconv.ParseInt(g[0][8].(string), 10, 64)
	mtime, _ := strconv.ParseInt(g[0][9].(string), 10, 64)

	var imgUrl []orm.ParamsList
	_, err = orm.NewOrm().Raw("select img_url from image where goods_id= ? ", id).ValuesList(&imgUrl)
	url := imgUrl[0][0].(string)
	var categoryName string
	categoryName = cat[category1].(string)

	goods := class.Goods{
		Id:           id,
		Title:        g[0][1].(string),
		Cause:        g[0][3].(string),
		Price:        g[0][4].(string),
		Desc:         g[0][5].(string),
		Category:     category1,
		Deal:         deal,
		Add_time:     atime,
		Modify_time:  mtime,
		Username:     g[0][10].(string),
		Origin_price: g[0][11].(string),
		Image:        url,
		CategoryName: categoryName,
	}

	fmt.Println(category1)
	/* 相同分类的商品推荐 */
	var lists []orm.ParamsList
	num, err = orm.NewOrm().Raw("select * from goods where category= ? order by add_time desc limit 0,4 ", category1).ValuesList(&lists)

	var relatedGoods [4]class.Goods
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

		relatedGoods[i] = goo
	}

	var count int64
	count = 0
	username := c.GetSession("username")
	// 取出购物车中商品
	if username != nil {
		var lists []orm.ParamsList
		num, _ := orm.NewOrm().Raw("select * from cart where username= ? ", username).ValuesList(&lists)
		count = num

		c.Data["Is_login"] = true
		c.Data["Username"] = username
	}

	fmt.Println(username)


	c.Data["Cart_count"] = count
	c.Data["Related"] = relatedGoods
	c.Data["Goods"] = goods
	c.Data["Cat"] = cat
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "preview.tpl"
}
