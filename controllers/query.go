package controllers

import (
	"emmm/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type QueryController struct {
	beego.Controller
}

func (t *QueryController) Get() {
	t.TplName = "teml.html"
	all := []models.Info{}

	DB.Where("scan_time LIKE ?", fmt.Sprintf(`%s`, time.Now().Format("200601")+`%`)).Order("department desc").Find(&all)
	//DB.Find(&all)
	t.Data["Data"] = all
	return
}

func (t QueryController) Post() {
	fmt.Println(t.Input().Get("name"))
	t.Ctx.WriteString("OK")

}
