package user

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"shop/pre_shop/models/class"
	"strings"
	"time"
)

type GoodsController struct {
	beego.Controller
}

func (g *GoodsController) Get() {

	g.TplName = "bad.tpl"
}

func (g *GoodsController) Sell() {
	title := g.GetString("title")
	cause := g.GetString("cause")
	price := g.GetString("price")
	desc := g.GetString("desc")
	category, _ := g.GetInt("category")
	deal, _ := g.GetInt("deal")
	origin_price := g.GetString("origin_price")

	// 参数验证
	valid := validation.Validation{}
	valid.Required(title, "title") // 校验是否为空值
	valid.Required(cause, "cause")
	valid.Required(price, "price")
	valid.Required(desc, "desc")
	valid.Required(category, "category")
	//valid.Required(deal, "deal")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			key := "标题"
			switch err.Key {
			case "cause":
				{
					key = "商品转卖原因"
				}
			case "price":
				{
					key = "商品出售价格"
				}
			case "desc":
				{
					key = "商品简介"
				}
			case "deal":
				{
					key = "交易方式"
				}
			}

			g.Data["Tips"] = key + " " + "不能为空"
			g.TplName = "bad.tpl"
			return
		}
	}
	username := g.GetSession("uid")
	fmt.Sprintf("%x", username)
	goods := &class.Goods{
		Title:        title,
		Cause:        cause,
		Price:        price,
		Desc:         desc,
		Category:     category,
		Deal:         deal,
		Add_time:     time.Now().Unix(),
		Modify_time:  time.Now().Unix(),
		Origin_price: origin_price,
		Username:     username.(string),
	}

	lastInsertId, err := goods.Create()
	if err != nil {
		fmt.Println(err)
		g.Data["Tips"] = "商品写入数据库失败"
		g.TplName = "bad.tpl"
		return
	} else {

		images1, h1, err1 := g.GetFile("images1")
		if err1 != nil {
			g.Ctx.WriteString("File retrieval failure")
			return
		}

		defer images1.Close()

		images1Ext := h1.Filename[strings.LastIndexAny(h1.Filename, ".")+1 : len(h1.Filename)]
		hash := md5.Sum([]byte(h1.Filename))
		images1Name := fmt.Sprintf("%x", hash)
		images1FinalUrl := "static/upload/images/" + images1Name + "." + images1Ext

		img := &class.Image{
			Img_url:  images1FinalUrl,
			Goods_id: lastInsertId,
		}
		//fmt.Println(images1FinalUrl)
		err2 := img.Create()
		if err2 == nil {
			err = g.SaveToFile("images1", images1FinalUrl)
		}

		g.Data["Tips"] = "商品发布成功"
		g.TplName = "bad.tpl"
		return
	}
}