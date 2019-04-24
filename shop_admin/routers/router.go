package routers

import (
	"shop/shop_admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/welcome", &controllers.WelcomeController{})
}
