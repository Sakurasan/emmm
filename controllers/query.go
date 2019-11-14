package controllers

import (
	"github.com/astaxie/beego"
)

type QueryController struct {
	beego.Controller
}

func (t *QueryController) Get() {
	t.Data["Website"] = "beego.me"
	t.Data["Email"] = "astaxie@gmail.com"
	t.TplName = "teml.html"
}
