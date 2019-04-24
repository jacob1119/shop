package routers

import (
	"github.com/astaxie/beego"
	"shop/pre_shop/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index",&controllers.IndexController{})
	beego.Router("/news",&controllers.NewsController{})
	beego.Router("/about",&controllers.AboutController{})
	beego.Router("/contact",&controllers.ContactController{})
	beego.Router("/delivery",&controllers.DeliveryController{})
	beego.Router("/preview",&controllers.PreviewController{})
}
