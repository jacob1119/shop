package user

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}


func (u *UserController) Get() {

	//username := u.GetSession("uid")
	//if username == nil {
	//	fmt.Println(username)
	//	u.Data["Tips"] = "请登录"
	//	u.TplName = "bad.tpl"
	//	return
	//}

	// 数据库中取出商品分类
	o := orm.NewOrm()
	var category []orm.ParamsList
	_, err := o.QueryTable("goods_category").ValuesList(&category, "id", "name")

	if err == nil {
		//c.Data["Cat"] = category
		//fmt.Println("aaaaerror")
	}
	var cat [20] interface{}
	for index,value := range category {
		cat[index+1] = value[1]
	}
	u.Data["Cat"] = cat




	u.Data["Website"] = "beego.me"
	u.Data["Email"] = "astaxie@gmail.com"
	u.TplName = "user/userinfo.tpl"
}
