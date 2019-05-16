package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"strconv"
)

type ListController struct {
	beego.Controller
}

func (l *ListController) Get() {
	// 获取商品总数
	var number []orm.ParamsList
	_, _ = orm.NewOrm().Raw("select count(*) from goods where sold=0 ").ValuesList(&number)

	count, _ := strconv.Atoi(number[0][0].(string))
	if count < 1 {
		l.Data["Tips"] = "Wows  没有数据"
		l.TplName = "bad.tpl"
		return
	}
	size := 12
	page,_ := l.GetInt("page")
	if page == 0 || count/size < 1 {
		page = 1
	}

	// 取出我的发售
	var lists []orm.ParamsList
	num, err := orm.NewOrm().Raw( "select * from goods where sold=0 order by add_time desc limit ?,12 " ,(page-1)*size).ValuesList(&lists)

	if err != nil || num < 0 {
		l.Data["Tips"] = "Wows  商品走丢了"
		l.TplName = "bad.tpl"
		return
	}

	var g [4]class.Goods
	var g2 [4]class.Goods
	var g3 [4]class.Goods
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
		}

		if i < 4 {
			g[i] = goo
		} else if i < 8 {
			g2[i-4] = goo
		} else {
			g3[i-8] = goo
		}

	}

	l.Data["Goods1"] = g
	l.Data["Goods2"] = g2
	l.Data["Goods3"] = g3
	l.Data["Count"] = count
	l.Data["Page"] = page
	l.Data["Prev"] = page - 1
	l.Data["Next"] = page + 1
	l.Data["Website"] = "beego.me"
	l.Data["Email"] = "astaxie@gmail.com"
	l.TplName = "list.tpl"
}

func (l *ListController) Category() {
	category,_ := l.GetInt("category")
	// 获取商品总数
	var number []orm.ParamsList
	_, _ = orm.NewOrm().Raw("select count(*) from goods where sold=0 and category= ? " , category).ValuesList(&number)

	count, _ := strconv.Atoi(number[0][0].(string))
	if count < 1 {
		l.Data["Tips"] = "Wows  没有数据"
		l.TplName = "bad.tpl"
		return
	}
	size := 12
	page,_ := l.GetInt("page")
	if page == 0 || count/size < 1 {
		page = 1
	}

	// 取出我的发售
	var lists []orm.ParamsList
	num, err := orm.NewOrm().Raw( "select * from goods where category= ? order by add_time desc limit ?,12 " ,category,(page-1)*size).ValuesList(&lists)

	if err != nil || num < 0 {
		l.Data["Tips"] = "Wows  商品走丢了"
		l.TplName = "bad.tpl"
		return
	}

	var g [4]class.Goods
	var g2 [4]class.Goods
	var g3 [4]class.Goods
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
	username := l.GetSession("username")
	// 取出购物车中商品
	if username != nil {
		var lists []orm.ParamsList
		num,_ := orm.NewOrm().Raw("select * from cart where is_order=0 and username= ? ",username).ValuesList(&lists)
		count1 = num

		l.Data["Is_login"] = true
		l.Data["Username"] = username
	}

	l.Data["Cart_count"] = count1
	l.Data["Goods1"] = g
	l.Data["Goods2"] = g2
	l.Data["Goods3"] = g3
	l.Data["Count"] = count
	l.Data["Page"] = page
	l.Data["Prev"] = page - 1
	l.Data["Next"] = page + 1
	l.Data["Website"] = "beego.me"
	l.Data["Email"] = "astaxie@gmail.com"
	l.TplName = "list.tpl"
}