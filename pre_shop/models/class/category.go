package class

import (
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type Goods_category struct {
	Id int         // 设置为主键，字段Id, Password首字母必须大写
	Name string
	Upload_time string
}

func (u *User) ReadDB1() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

//func (u *User) GetCategoryName (res string) {
//	var categoryList []orm.ParamsList
//	num,err := orm.NewOrm().Raw("select * from goods_category").ValuesList(&categoryList)
//
//	if err == nil || num >0 {
//		fmt.Println(categoryList)
//	}
//
//
//
//}