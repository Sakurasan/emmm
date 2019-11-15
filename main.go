package main

import (
	"emmm/controllers"
	_ "emmm/routers"
	"github.com/astaxie/beego"
)

func main() {
	controllers.InitDb()
	beego.Run()
}
