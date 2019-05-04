package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Goods struct {
	Id           int // 设置为主键，字段Id, Password首字母必须大写
	Title        string
	Cause        string
	Price        string
	Desc         string
	Category     int
	Deal         int
	Add_time     int64
	Modify_time  int64
	Username     string
	Origin_price string
}

func (g *Goods) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(g)
	return id, err
}
