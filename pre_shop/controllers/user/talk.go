package user

import (
	"github.com/astaxie/beego"
	"shop/pre_shop/models/class"
	"time"
)

type TalkController struct {
	beego.Controller
}

func (t *TalkController) Get() {

	t.Data["Website"] = "beego.me"
	t.Data["Email"] = "astaxie@gmaio.com"
	t.TplName = "cart.tpl"
}

func (t *TalkController) Contact() {
	toUser := t.GetString("touser")
	fromUser := t.GetString("fromuser")
	title := t.GetString("title")
	content := t.GetString("content")

	talk := &class.Talk{
		From_user: fromUser,
		To_user:   toUser,
		Title:     title,
		Content:   content,
		Add_time:  time.Now().Unix(),
	}

	_, err := talk.Create()
	if err == nil {
		t.Data["Tips"] = "消息发送成功"
		t.TplName = "bad.tpl"
		return
	} else {
		t.Data["Tips"] = "消息发送失败，请稍后再试"
		t.TplName = "bad.tpl"
		return
	}
}
