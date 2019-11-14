package main

import (
	_ "emmm/routers"
	"github.com/astaxie/beego"
	"emmm/controllers"
)

func main() {
	controllers.InitDb()
	beego.Run()
}

