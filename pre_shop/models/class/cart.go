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
	Is_order int
}

func (c *Cart) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(c)
	return id, err
}

func (c *Cart) Modify() (rows int64,err error) {
	o := orm.NewOrm()
	o.Using("default")
	num, err := o.Update(c,"is_order")

	return num,err

}
func (c *Cart) Del() (num int64,err error) {
	o := orm.NewOrm()
	o.Using("default")
	num,err = o.Delete(c)

	return num,err
}