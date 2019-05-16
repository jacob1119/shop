package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"strconv"
)

type CartController struct {
	beego.Controller
}

func (c *CartController) Get() {

	username := c.GetSession("username")
	//uid := u.GetSession("uid")
	if username == nil {
		//fmt.Println(username)
		c.Data["Tips"] = "请登录"
		c.TplName = "bad.tpl"
		return
	}

	// 获取商品总数
	var number []orm.ParamsList
	_, _ = orm.NewOrm().Raw("select count(*) from cart where is_order=0 and username= ? ", username).ValuesList(&number)

	count, _ := strconv.Atoi(number[0][0].(string))
	if count < 1 {
		c.Data["Tips"] = "Wows  购物车为空"
		c.TplName = "bad.tpl"
		return
	}
	size := 12
	page, _ := c.GetInt("page")
	if page == 0 || count/size < 1 {
		page = 1
	}

	// 取出我的发售
	var lists []orm.ParamsList
	num, err := orm.NewOrm().Raw("select * from cart where is_order=0 and username= ? order by add_time desc limit ?,12 ", username, (page-1)*size).ValuesList(&lists)

	if err != nil || num < 0 {
		c.Data["Tips"] = "Wows  商品走丢了"
		c.TplName = "bad.tpl"
		return
	}

	//fmt.Println(lists)
	//os.Exit(0)
	var g [4]class.Cart
	var g2 [4]class.Cart
	var g3 [4]class.Cart
	var i int64
	for i = 0; i < num; i++ {

		id, _ := strconv.Atoi(lists[i][0].(string))
		goods_id, _ := strconv.Atoi(lists[i][2].(string))

		var imgUrl []orm.ParamsList
		_, err = orm.NewOrm().Raw("select img_url from image where goods_id= ? ", goods_id).ValuesList(&imgUrl)
		url := imgUrl[0][0].(string)

		goo := class.Cart{
			Id:       id,
			Goods_id: goods_id,
			Title:    lists[i][3].(string),
			Price:    lists[i][4].(string),
			Username: lists[i][1].(string),
			Url:      url,
		}

		if i < 4 {
			g[i] = goo
		} else if i < 8 {
			g2[i-4] = goo
		} else {
			g3[i-8] = goo
		}

	}

	var count1 int64
	count1 = 0
	username = c.GetSession("username")
	// 取出购物车中商品
	if username != nil {
		var lists []orm.ParamsList
		num, _ := orm.NewOrm().Raw("select * from cart where is_order=0 and username= ? ", username).ValuesList(&lists)
		count1 = num

		c.Data["Is_login"] = true
		c.Data["Username"] = username
	}

	c.Data["Cart_count"] = count1
	c.Data["Goods1"] = g
	c.Data["Goods2"] = g2
	c.Data["Goods3"] = g3
	c.Data["Count"] = count
	c.Data["Page"] = page
	c.Data["Prev"] = page - 1
	c.Data["Next"] = page + 1
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmaic.com"
	c.TplName = "cart.tpl"
}
