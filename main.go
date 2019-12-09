package main

import (
	"emmm/controllers"
	_ "emmm/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	dbconf, _ := beego.AppConfig.GetSection("database")
	fmt.Println(dbconf)
	fmt.Println("数据库类型:", dbconf["type"])
	if "mysql" == dbconf["type"] {
		controllers.InitMysqlDb(dbconf)
	} else {
		controllers.InitDb()
	}

	beego.Run()
}
