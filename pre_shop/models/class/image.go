package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Image struct {
	Id       int // 设置为主键，字段Id, Password首字母必须大写
	Img_url  string
	Goods_id int64
}


func (i *Image) Create() (err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	 _,err = o.Insert(i)
	return err
}