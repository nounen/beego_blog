package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// TestFilter 中间件测试案例
func TestFilter() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		beego.Debug("当前控制器和方法")
		beego.Debug(ctx.Input.RunController)
		beego.Debug(ctx.Input.RunMethod)
		beego.Debug(ctx.Input.Params())
	}
}
