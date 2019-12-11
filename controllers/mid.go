package controllers

import (
	"github.com/astaxie/beego/context"
	"strings"
)

func getLocalAccount(ctx *context.Context) (string, error) {
	name, err := ctx.Request.Cookie("uid")
	if err != nil {
		return "", err
	}
	return name.Value, nil
}

func checkUserAgent(ctx *context.Context) bool {
	return strings.Contains(ctx.Request.UserAgent(), "Wechat")
}
