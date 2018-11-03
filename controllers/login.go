package controllers

import (
	"beego_blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	BaseController
}

// Login 用户登录，得到 jwt token
func (c *LoginController) Login() {
	token, _ := utils.GenerateToken("admin", "111111", 1)

	claims, _ := utils.ParseToken(token)

	beego.Debug(token)
	beego.Debug(claims)
	logs.Debug(claims)
}

func (c *LoginController) Logout() {}
