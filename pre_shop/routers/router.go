package routers

import (
	"github.com/astaxie/beego"
	"shop/pre_shop/controllers"
	"shop/pre_shop/controllers/user"
)

func init() {
	/* 首页显示界面路由 */
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/index",&controllers.IndexController{})
	beego.Router("/about",&controllers.AboutController{})
	beego.Router("/about/contact",&controllers.AboutController{},`get:Contact`)
	beego.Router("/contact",&controllers.ContactController{})
	beego.Router("/delivery",&controllers.DeliveryController{})
	beego.Router("/preview",&controllers.PreviewController{})


    // 用户模块
    beego.Router("/login",&controllers.LoginController{})
	beego.Router("/user/login",&controllers.LoginController{},`post:Login`)
	beego.Router("/register",&controllers.RegisterController{})
	beego.Router("/user/register",&controllers.RegisterController{},`post:Register`)


	//用户个人中心模块
	beego.Router("/user/info",&user.UserController{})


	//商品模块
	beego.Router("/user/goods",&user.GoodsController{},`post:Sell`)
	beego.Router("/goods/cart",&user.GoodsController{},`get:Cart`)
	beego.Router("/cart/show",&user.CartController{})
	beego.Router("/goods/del",&user.GoodsController{},`get:Del`)

    beego.Router("/goods/list",&controllers.ListController{})
	beego.Router("/goods/category",&controllers.ListController{},`get:Category`)


    //订单
	beego.Router("/user/pay",&user.OrderController{},`get:Pay`)
	beego.Router("/order/del",&user.OrderController{},`get:Del`)


    // 通知公告
	beego.Router("/user/contact",&user.TalkController{},`post:Contact`)

    // 搜索
	beego.Router("/goods/search",&controllers.ListController{},`get:Search`)

}
