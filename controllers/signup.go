package controllers

import (
	"emmm/models"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"time"
)

type SignUpController struct {
	beego.Controller
}

func (c *SignUpController) Get() {
	c.TplName = "signup.html"
	c.Data["GET"] = true
}

func (c *SignUpController) Post() {
	c.TplName = "signup.html"
	Department := c.GetString("Department")
	Name := c.GetString("Name")
	ComputerType := c.GetString("ComputerType")
	OsType := c.GetString("OsType")
	MacAddr := c.GetString("MacAddr")
	SecuritySoftWare := c.GetString("SecuritySoftWare")

	fmt.Println(Department, Name, ComputerType, OsType, MacAddr, SecuritySoftWare)

	user := models.User{}
	user.Department = Department
	user.Name = Name
	if ComputerType == "" {
		user.ComputerType = "WinPC"
	} else {
		user.ComputerType = ComputerType
	}
	user.OsType = OsType
	user.MacAddr = MacAddr
	user.SecuritySoftWare = SecuritySoftWare

	vcode := fmt.Sprintf("%v", rand.Int31n(10000))
	user.UserId = time.Now().Format("20060102150405") + vcode
	DB.Create(&user)
	DB.Where("name = ?", user.Name).Find(&user)
	fmt.Println(user.UserId, user.ID)
	c.Ctx.SetCookie("uid", fmt.Sprintf("%s", user.UserId), 3600*24*300, "/")

	c.Data["POST"] = true
	return
}
