package routers

import (
	"emmm/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.QueryController{})
	beego.Router("/signup",&controllers.SignUpController{})
}
