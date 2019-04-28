package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id int         // 设置为主键，字段Id, Password首字母必须大写
	Username string
	Real_name string
	Password string
	Unique_key string
	Phone string
	Email string
	Add_time int64
	Modify_time int64
	Active int
}

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	o.Using("default")
	fmt.Println("Create success!")
	_, _ = o.Insert(u)
	return err
}

func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}

func (u *User) GetDa() (err error)  {
	o := orm.NewOrm()
	user := User{Username:"jacob"}

	err = o.Read(&user)

	return err
}

