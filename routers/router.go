// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego_blog/controllers"
	"beego_blog/filters"
	"github.com/astaxie/beego"
)

func init() {
	login := beego.NewNamespace("/login",
		beego.NSRouter("/login", &controllers.LoginController{}, "post:Login"),
		beego.NSRouter("/logout", &controllers.LoginController{}, "post:Logout"),
	)

	v1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/tag",
			beego.NSRouter("/", &controllers.TagController{}, "get:Index"),
			beego.NSRouter("/", &controllers.TagController{}, "post:Store"),
			beego.NSRouter("/?:id:int", &controllers.TagController{}, "get:Show"),
			beego.NSRouter("/?:id:int", &controllers.TagController{}, "put:Update"),
			beego.NSRouter("/?:id:int", &controllers.TagController{}, "delete:Delete"),
		),

		beego.NSNamespace("/article",
			beego.NSRouter("/", &controllers.ArticleController{}, "get:Index"),
			beego.NSRouter("/", &controllers.ArticleController{}, "post:Store"),
			beego.NSRouter("/?:id:int", &controllers.ArticleController{}, "get:Show"),
			beego.NSRouter("/?:id:int", &controllers.ArticleController{}, "put:Update"),
			beego.NSRouter("/?:id:int", &controllers.ArticleController{}, "delete:Delete"),
		),

		beego.NSNamespace("/user",
			//beego.NSRouter("/", &controllers.UserController{}, "get:Index"),
			beego.NSRouter("/", &controllers.UserController{}, "post:Store"),
			//beego.NSRouter("/?:id:int", &controllers.UserController{}, "get:Show"),
			beego.NSRouter("/?:id:int", &controllers.UserController{}, "put:Update"),
			//beego.NSRouter("/?:id:int", &controllers.UserController{}, "delete:Delete"),
		),
	)

	beego.AddNamespace(login, v1)

	// 中间件
	beego.InsertFilter("/v1/*", beego.BeforeRouter, filters.JwtFilter())
}
