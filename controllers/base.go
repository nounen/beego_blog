package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Json map[string]interface{}
}

// Prepare 优先执行于其他方法
func (c *BaseController) Prepare() {
	// 用于 json 数据返回
	c.Json = map[string]interface{}{}
}

// RespondJson 响应 json
func (c *BaseController) RespondJson() {
	c.Data["json"] = &c.Json
	c.ServeJSON()
}
