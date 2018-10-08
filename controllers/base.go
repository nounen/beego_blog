package controllers

import (
	"beego_blog/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"net/http"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	Json        map[string]interface{}
	RequestData map[string]interface{}
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
	c.StopRun() // 防止出错后继续执行后面的代码
}

// RespondBadJson 错误响应
func (c *BaseController) RespondBadJson(err error) {
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.Json["message"] = "操作失败"
	c.Json["error"] = err.Error()
	c.RespondJson()
}

// RespondBadEntityJson 如果表单验证失败响应
func (c *BaseController) RespondIfBadEntityJson(valid *validation.Validation) {
	if valid.HasErrors() {
		c.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)

		message := ""
		messages := map[string]string{}

		for _, error := range valid.Errors {
			message += error.Key + " " + error.Message + ";"
			messages[error.Key] = error.Message
		}

		c.Json["message"] = message
		c.Json["messages"] = messages
		c.RespondJson()
	}
}

// RespondCreatedJson 操作成功响应
func (c *BaseController) RespondCreatedJson() {
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Json["message"] = "操作成功"
	c.RespondJson()
}

// RespondNoContentJson 操作成功响应
func (c *BaseController) RespondNoContentJson() {
	c.Ctx.Output.SetStatus(http.StatusNoContent)
	c.Json["message"] = "操作成功"
	c.RespondJson()
}

// @isLast 是否为最后一个sql操作
// RespondByTransaction 事务响应
func (c *BaseController) RespondByTransaction(orm orm.Ormer, err error, isLast bool) {
	// 事务失败处理： 回滚事务，响应错误
	if err != nil {
		orm.Rollback()
		c.RespondBadJson(err)
	}

	// 事务成功且是最后一条数据操作: 提交事务，响应成功
	if err == nil && isLast == true {
		orm.Commit()
		c.RespondCreatedJson()
	}
}

// UnmarshalRequestJson 解码json请求数据
func (c *BaseController) UnmarshalRequestJson(RequestBody interface{}) interface{} {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, RequestBody); err != nil {
		c.RespondBadJson(err)
	}

	return RequestBody
}

// UnmarshalRequestData 解码json请求数据
func (c *BaseController) UnmarshalRequestData() {
	var requestData map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData); err != nil {
		c.RespondBadJson(err)
	}

	c.RequestData = requestData
}

// UnmarshalRequestData 解码json请求数据
func (c *BaseController) UnmarshalRequest(requestData interface{}) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData); err != nil {
		c.RespondBadJson(err)
	}
}

// getPage 当前页码
func (c *BaseController) getPage() int64 {
	page, _ := strconv.ParseInt(c.Ctx.Input.Query("page"), 10, 64)

	if page <= 0 {
		page = 1
	}

	return page
}

// getPerPage 列表
func (c *BaseController) getPerPage() int64 {
	perPage, _ := strconv.ParseInt(c.Ctx.Input.Query("per_page"), 10, 64)

	if perPage == 0 {
		perPage = 5
	}

	return perPage
}

// getFilters 从url里面解析出过滤条件 （分页用）
func (c *BaseController) getFilters(defaultOrder bool) *utils.Filters {
	filters := utils.Filters{}
	filters.Ins = map[string][]string{}
	filters.Betweens = map[string][]string{}

	c.Ctx.Input.Bind(&filters.Orders, "order")
	c.Ctx.Input.Bind(&filters.Equals, "equal")
	c.Ctx.Input.Bind(&filters.Likes, "like")

	betweens := map[string]string{}
	c.Ctx.Input.Bind(&betweens, "between")
	for key, value := range betweens {
		filters.Betweens[key] = strings.Split(value, ",")
	}

	ins := map[string]string{}
	c.Ctx.Input.Bind(&ins, "in")
	for key, value := range ins {
		filters.Ins[key] = strings.Split(value, ",")
	}

	// 默认 id 倒序
	if defaultOrder {
		if len(filters.Orders) == 0 {
			filters.Orders["id"] = "desc"
		}
	}

	return &filters
}
