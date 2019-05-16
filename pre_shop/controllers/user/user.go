package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {

	username := u.GetSession("username")
	//uid := u.GetSession("uid")
	if username == nil {
		fmt.Println(username)
		u.Data["Tips"] = "请登录"
		u.TplName = "bad.tpl"
		return
	}

	// 数据库中取出商品分类
	o := orm.NewOrm()
	var category []orm.ParamsList
	_, err := o.QueryTable("goods_category").ValuesList(&category, "id", "name")

	if err == nil {
		//c.Data["Cat"] = category
		//fmt.Println("aaaaerror")
	}
	var cat [20] interface{}
	for index, value := range category {
		cat[index+1] = value[1]
	}

	// 取出我的发售
	var lists []orm.ParamsList
	num, err := orm.NewOrm().Raw("select * from goods where username= ? order by add_time desc", username).ValuesList(&lists)

	var g [10]class.Goods
	var i int64
	for i = 0; i < num; i++ {

		id, _ := strconv.Atoi(lists[i][0].(string))
		category, _ := strconv.Atoi(lists[i][6].(string))
		deal, _ := strconv.Atoi(lists[i][7].(string))
		atime, _ := strconv.ParseInt(lists[i][8].(string), 10, 64)
		mtime, _ := strconv.ParseInt(lists[i][9].(string), 10, 64)

		timeLayout := "2006-01-02 15:04:05"
		dataTimeStr := time.Unix(atime, 0).Format(timeLayout)

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
			Time_string:  dataTimeStr,
		}

		g[i] = goo
	}
	//fmt.Println(g)

	// 取出通知消息




	// 取出我的购买
	var myShop []orm.ParamsList
	num, err = orm.NewOrm().Raw("select * from `order` where username= ? ", username).ValuesList(&myShop)

	var myOrder [10]class.Order

	if num > 0 {
		var i int64
		for i = 0; i < num; i++ {

			id, _ := strconv.Atoi(myShop[i][0].(string))
			goods_id,_ := strconv.Atoi(myShop[i][1].(string))
			status,_ := strconv.Atoi(myShop[i][7].(string))
			atime, _ := strconv.ParseInt(myShop[i][5].(string), 10, 64)
			timeLayout := "2006-01-02 15:04:05"
			dataTimeStr := time.Unix(atime, 0).Format(timeLayout)

			var imgUrl []orm.ParamsList
			_, err = orm.NewOrm().Raw("select img_url from image where goods_id= ? ", goods_id).ValuesList(&imgUrl)
			url := imgUrl[0][0].(string)

			var bossUsername []orm.ParamsList
			_, err = orm.NewOrm().Raw("select username from goods where id= ? ", goods_id).ValuesList(&bossUsername)

			//fmt.Println(bossUsername)
			bossName := bossUsername[0][0].(string)
			//bossName := "jacob"

			myOrd := class.Order{
				Id :      id,
				Goods_id :goods_id,
				Username :myShop[i][2].(string),
				Title   : myShop[i][3].(string),
				Price   : myShop[i][4].(string),
				Url     : url,
				Status :status,
				Str_time:dataTimeStr,
				Boss:bossName,
			}

			myOrder[i] = myOrd
		}
	}

	//fmt.Println(myOrder)

	u.Data["MyOrder"] = myOrder // 我的购买
	u.Data["Goods"] = g         // 我的发售
	u.Data["Cat"] = cat			// 分类
	u.Data["Username"] = username
	u.Data["Title"] = "个人中心"
	u.Data["Website"] = "beego.me"
	u.Data["Email"] = "astaxie@gmail.com"
	u.TplName = "user/userinfo.tpl"
}
