package filters

import (
	"beego_blog/utils"
	"github.com/astaxie/beego/context"
	"net/http"
)

// TestFilter 中间件测试案例
func JwtFilter() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		_, err := utils.GetUserId(ctx)

		if !(err == nil) {
			// 用户登录提交信息
			type LoginFailed struct {
				Message string `json:"message"`
				Error   string `json:"error"`
			}

			loginFailed := LoginFailed{
				Message: "操作失败",
				Error:   err.Error(),
			}

			ctx.Output.SetStatus(http.StatusBadRequest)
			ctx.Output.JSON(loginFailed, false, false)
		}
	}
}
