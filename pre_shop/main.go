package main

import (
	"github.com/astaxie/beego"
	_ "shop/pre_shop/routers"
	_ "shop/pre_shop/models"
)

func main() {
	beego.Run()
}

