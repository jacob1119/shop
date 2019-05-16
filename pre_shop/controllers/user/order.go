package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shop/pre_shop/models/class"
	"time"
)

type OrderController struct {
	beego.Controller
}

func (o *OrderController) Get() {

	o.Data["Website"] = "beego.me"
	o.Data["Email"] = "astaxie@gmaio.com"
	o.TplName = "cart.tpl"
}

func (o *OrderController) Pay() {
	username := o.GetSession("username")
	//uid := u.GetSession("uid")
	if username == nil {
		//fmt.Println(username)
		o.Data["Tips"] = "您还未登陆"
		o.TplName = "bad.tpl"
		return
	}

	title := o.GetString("title")
	price := o.GetString("price")
	goodsId, _ := o.GetInt("id")
	cid, _ := o.GetInt("cid")
	if goodsId < 1 {
		o.Data["Tips"] = "系统出现错误"
		o.TplName = "bad.tpl"
		return
	}

	var lists []orm.ParamsList
	num, _ := orm.NewOrm().Raw("select * from order where  goods_id= ?", goodsId).ValuesList(&lists)
	fmt.Println(num)

	if num > 0 {
		o.Data["Tips"] = "此商品已经被其他用户购买"
		o.TplName = "bad.tpl"
		return
	}

	//fmt.Println(title,price,goodsId)

	order := class.Order{
		Goods_id: goodsId,
		Username: username.(string),
		Title:    title,
		Price:    price,
		Add_time: time.Now().Unix(),
	}

	_, err := order.Create()
	if err != nil {
		fmt.Println(err)
		o.Data["Tips"] = "商品下单失败、请稍后重试"
		o.TplName = "bad.tpl"
		return
	} else {
		// 订单已经生成 将cart内的is_order置为1  从购物车中删除商品
		or := class.Cart{
			Id:       cid,
		}

		num1, err1 := or.Del()
		fmt.Println(num1, err1)

		// 订单生成后将goods表中sold状态置为1
		goods := class.Goods{
			Id:   goodsId,
			Sold: 1,
		}
		_,_ = goods.Modify()

		o.Data["Tips"] = "订单已经生成，待卖方确认订单并联系交易方式。您可以在个人中心查看订单和卖家的通知"
		o.TplName = "bad.tpl"
		return
	}
}

func (o *OrderController) Del() {
	username := o.GetSession("username")
	//uid := u.GetSession("uid")
	if username == nil {
		//fmt.Println(username)
		o.Data["Tips"] = "您还未登陆"
		o.TplName = "bad.tpl"
		return
	}

	id, _ := o.GetInt("id")
	goods_id, _ := o.GetInt("gid")
	title := o.GetString("title")

	if id < 0 {
		o.Data["Tips"] = "撤单失败，稍后重试"
		o.TplName = "bad.tpl"
		return
	}

	order := &class.Order{
		Id: id,
	}

	num, _ := order.Del()
	if num > 0 {
		// 订单生成后将goods表中sold状态置为1
		goods := class.Goods{
			Id:   goods_id,
			Sold: 0,
		}
		_,_ = goods.Modify()

		o.Data["Tips"] = "商品名称为" + title + " 的订单已经撤单 "
		o.TplName = "bad.tpl"
		return

	} else {
		o.Data["Tips"] = "商品名称为" + title + " 的订单撤单失败 "
		o.TplName = "bad.tpl"
		return
	}

}
