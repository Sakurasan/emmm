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

	DB.Where("scan_time LIKE ?", fmt.Sprintf(`%s`, time.Now().Format("200601")+`%`)).Order("department asc").Find(&all)
	//DB.Find(&all)
	t.Data["Data"] = all
	return
}

func (t *QueryController) Post() {
	t.TplName = "teml.html"
	department := t.Input().Get("partment")
	scantime := t.GetString("scantime2")
	fmt.Println("partment:", department)
	fmt.Println("scantime:", t.GetString("scantime"))
	fmt.Println("scantime:", scantime)
	all := []models.Info{}
	if department != "" {
		if scantime != "" {
			fmt.Println("查询department，scantime")
			DB.Where("scan_time LIKE ? AND department = ? ", scantime+`%`, department).Order("department asc").Find(&all)
		} else {
			fmt.Println("查询department")
			// DB.Where(" department = ? ", department).Find(&all)
			DB.Where(" department = ? AND scan_time LIKE ?", department, time.Now().Format("200601")+"%").Find(&all)
			// DB.Exec("select * from info where scan_time like ? and  department = ?  order by department asc", time.Now().Format("200601")+"%", department).Scan(&all)
		}
	} else {
		if scantime != "" {
			fmt.Println("查询scantime")
			DB.Where("scan_time LIKE ? ", scantime+`%`).Order("department desc").Find(&all)
		} else {
			fmt.Println("默认")
			DB.Where("scan_time LIKE ?", fmt.Sprintf(`%s`, time.Now().Format("200601")+"%")).Order("department asc").Find(&all)
		}
	}
	t.Data["Data"] = all
	return
}
