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
	Image        string
	CategoryName string
	Sold         int
	Time_string  string
	Order_id     int
	Buyer        string
}

// 完成User类型定义
type GoodShow struct {
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
	Image        string
	CategoryName string
	Sold         int
}

func (g *Goods) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(g)
	return id, err
}
func (g *Goods) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(g)
	return err
}

func (g *Goods) Del() (num int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	num, err = o.Delete(g)

	return num, err
}

func (g *Goods) Modify() (rows int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	num, err := o.Update(g, "sold")

	return num, err
}
