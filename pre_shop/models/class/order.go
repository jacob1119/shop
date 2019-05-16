package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Order struct {
	Id       int // 设置为主键，字段Id, Password首字母必须大写
	Goods_id int
	Username string
	Title    string
	Price    string
	Add_time int64
	Url      string
	Status int
	Str_time string
	Boss string
}

func (c *Order) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(c)
	return id, err
}
func (c *Order) Del() (num int64,err error) {
	o := orm.NewOrm()
	o.Using("default")
	num,err = o.Delete(c)

	return num,err
}