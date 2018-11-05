package controllers

import (
	"beego_blog/models"
	"beego_blog/utils"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

// Store 创建用户
func (c *UserController) Store() {
	user := c.getUserFromRequest()
	c.checkUserFromRequest(user)

	user.Password = utils.Encryption(user.Password)

	if _, err := models.AddUser(user); err == nil {
		c.RespondCreatedJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Update 更新用户
func (c *UserController) Update() {
	user := c.getUserFromRequest()
	c.checkUserFromRequest(user)

	user.Id = c.getId()
	user.Password = utils.Encryption(user.Password)

	if err := models.UpdateUserById(user); err == nil {
		c.RespondNoContentJson()
	} else {
		c.RespondBadJson(err)
	}
}

// getUserFromRequest 获取表单提交数据
func (c *UserController) getUserFromRequest() *models.User {
	user := &models.User{}
	user.CreatedAt = utils.GetNow()
	c.UnmarshalRequestJson(user)
	return user
}

// checkUserFromRequest 表单验证
func (c *UserController) checkUserFromRequest(User *models.User) {
	valid := validation.Validation{}

	valid.Required(User.Name, "name")
	valid.MaxSize(User.Name, 12, "Name")

	valid.Required(User.Password, "password")
	valid.MinSize(User.Password, 6, "password")
	valid.MaxSize(User.Password, 12, "password")

	valid.Required(User.State, "state")
	valid.Range(User.State, 1, 3, "state")

	c.RespondIfBadEntityJson(&valid)
}
