package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SignUpController struct {
	beego.Controller
}

func (c *SignUpController) Get(){
	c.TplName = "signup.html"
}

func (c *SignUpController) Post(){
	c.TplName = "signup.html"
	fmt.Println(c.GetString("Department"),c.GetString("Name"),c.GetString("ComputerType"),c.GetString("OsType"),
		c.GetString("MacAddr"),c.GetString("SecuritySoftWare"))
}
