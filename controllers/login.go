package controllers

import (
	"beego_blog/models"
	"beego_blog/services"
	"beego_blog/utils"
	"errors"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}

// Login 用户登录，得到 jwt token
func (c *LoginController) Login() {
	loginInfo := c.getLoginInfoFromRequest()
	c.checkLoginInfoFromRequest(loginInfo)

	user, _ := models.GetUserByName(loginInfo.Name)
	if user == nil {
		err := errors.New("用户不存在")
		c.RespondBadJson(err)
	}

	result := utils.CheckEncryption(user.Password, loginInfo.Password)
	if !(result == nil) {
		err := errors.New("密码错误")
		c.RespondBadJson(err)
	}

	token, _ := utils.GenerateToken(user.Id, user.Name, user.Password)
	c.Json["token"] = &token
	c.RespondJson()
	//claims, _ := utils.ParseToken(token)
	//beego.Debug(claims)
}

func (c *LoginController) Logout() {

}

// getLoginInfoFromRequest 获取表单提交数据
func (c *LoginController) getLoginInfoFromRequest() *services.LoginInfo {
	loginInfo := &services.LoginInfo{}
	c.UnmarshalRequestJson(loginInfo)
	return loginInfo
}

// checkUserFromRequest 表单验证
func (c *LoginController) checkLoginInfoFromRequest(loginInfo *services.LoginInfo) {
	valid := validation.Validation{}

	valid.Required(loginInfo.Name, "name")
	valid.MaxSize(loginInfo.Name, 12, "Name")

	valid.Required(loginInfo.Password, "password")
	valid.MinSize(loginInfo.Password, 6, "password")
	valid.MaxSize(loginInfo.Password, 12, "password")

	c.RespondIfBadEntityJson(&valid)
}
