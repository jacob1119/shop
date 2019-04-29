package user

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


func (u *UserController) Get() {

	username := u.GetSession("uid")
	if username == nil {
		fmt.Println(username)
		u.Data["Tips"] = "请登录"
		u.TplName = "bad.tpl"
		return
	}


	u.Data["Website"] = "beego.me"
	u.Data["Email"] = "astaxie@gmail.com"
	u.TplName = "user/userinfo.tpl"
}
