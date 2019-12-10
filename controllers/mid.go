package controllers

import (
	"github.com/astaxie/beego/context"
)

func getLocalAccount(ctx *context.Context) (string, error) {
	name, err := ctx.Request.Cookie("uid")
	if err != nil {
		return "", err
	}
	return name.Value, nil
}
