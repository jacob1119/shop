package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Talk struct {
	Id        int // 设置为主键，字段Id, Password首字母必须大写
	From_user string
	To_user   string
	Title     string
	Content   string
	Add_time  int64
	Str_time  string
	Is_read int
}

func (c *Talk) Create() (id int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	id, err = o.Insert(c)
	return id, err
}
func (c *Talk) Del() (num int64, err error) {
	o := orm.NewOrm()
	o.Using("default")
	num, err = o.Delete(c)

	return num, err
}
