package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Cart struct {
	Id       int // 设置为主键，字段Id, Password首字母必须大写
	Username string
	Goods_id int
	Title    string
	Price    string
	Add_time int64
	Url      string
}

func (c *Cart) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(c)
	return id, err
}
