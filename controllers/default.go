package controllers

import (
	"fmt"

	"emmm/models"
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	uid, err := getLocalAccount(c.Ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		user := models.User{}
		DB.Where("user_id = ?", uid).Find(&user)
		fmt.Println("user:", uid, user.Name)
		c.Data["Name"] = user.Name
	}

	c.Data["GET"] = true
}

func (t *MainController) Post() {
	t.TplName = "index.html"

	newinfo := models.Info{}

	fmt.Println("看这里:", t.GetString("Department"), t.GetString("Name"), t.GetString("ComputerType"), t.GetString("OsType"))

	newinfo.Department = t.GetString("Department")
	newinfo.Name = t.GetString("Name")
	newinfo.ComputerType = t.GetString("ComputerType")
	newinfo.OsType = t.GetString("OsType")
	newinfo.MacAddr = t.GetString("MacAddr")
	newinfo.SecuritySoftWare = t.GetString("SecuritySoftWare")
	newinfo.ScanTime = time.Now().Format("20060102")
	newinfo.ScanResult = t.GetString("ScanResult")
	newinfo.VirusNum = t.GetString("VirusNum")
	newinfo.AbnormalNum = t.GetString("AbnormalNum")
	newinfo.HowDealWith = t.GetString("HowDealWith")
	newinfo.DealName = t.GetString("DealName")

	DB.Create(&newinfo)
	t.Data["POST"] = true
	t.Data["Status"] = "OK"
}
