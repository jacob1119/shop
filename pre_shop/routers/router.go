package routers

import (
	"github.com/astaxie/beego"
	"shop/pre_shop/controllers"
	"shop/pre_shop/controllers/user"
)

func init() {
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

    beego.Router("/goods/list",&controllers.ListController{})
	beego.Router("/goods/category",&controllers.ListController{},`get:Category`)


}
