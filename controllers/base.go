package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
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

// getId 获取路由上的 :id 参数, 并做类型转换
func (c *BaseController) getId() int64 {
	// string 转 int64, 如何直接拿到想要类型呢
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)

	return id
}

// RespondJson 响应 json
func (c *BaseController) RespondJson() {
	c.Data["json"] = &c.Json
	c.ServeJSON()
}

// RespondBadJson 错误响应
func (c *BaseController) RespondBadJson(err error) {
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.Json["message"] = "操作失败"
	c.Json["error"] = err.Error()
	c.RespondJson()
	c.StopRun() // 防止出错后继续执行后面的代码
}

// RespondBadJson 操作成功响应
func (c *BaseController) RespondCreatedJson() {
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Json["message"] = "操作成功"
	c.RespondJson()
}

// UnmarshalRequestJson 解码json请求数据
func (c *BaseController) UnmarshalRequestJson(RequestBody interface{}) interface{} {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, RequestBody); err != nil {
		c.RespondBadJson(err)
	}

	return RequestBody
}
