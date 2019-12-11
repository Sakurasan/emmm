package controllers

import (
	"fmt"

	"emmm/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	if checkUserAgent(c.Ctx) {
		c.Data["Browser"] = true
		return
	}
	uid, err := getLocalAccount(c.Ctx)
	if err != nil {
		fmt.Println(err)
	} else if uid != "" && len(uid) > 9 {
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
	user := models.User{}
	err := DB.Where("name = ?", newinfo.Name).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("未找到")
		} else {
			fmt.Println("其他错:", err)
		}
		t.Redirect("/signup", 302)
		return
	}

	if user.Name != "" && len(user.Name) > 2 {
		findinfo := models.Info{}
		err := DB.Where("name = ? AND scan_time = ?", user.Name, time.Now().Format("20060102")).First(&findinfo).Error
		if err != nil {
			fmt.Println(err)
			if gorm.IsRecordNotFoundError(err) {
				fmt.Println("本月未找到，允许记录")
				newinfo.Department = user.Department
				newinfo.ComputerType = user.ComputerType
				newinfo.OsType = user.OsType
				newinfo.MacAddr = user.MacAddr
				newinfo.SecuritySoftWare = user.SecuritySoftWare
				DB.Create(&newinfo)
				t.Ctx.SetCookie("uid", fmt.Sprintf("%s", user.UserId), 3600*24*300, "/")
				t.Data["POST"] = true
				t.Data["Status"] = "完毕"
				return
			} else {
				t.Data["POST"] = true
				t.Data["Status"] = "请重试" + err.Error()
			}
		}
		if findinfo.Name != "" && len(findinfo.Name) >= 2 {
			t.Data["POST"] = true
			t.Data["Status"] = "本月重复提交"
		}

	}

}
