package controllers

import (
	"beego_blog/models"
)

type TagController struct {
	BaseController
}

// Index 列表数据
func (c *TagController) Index() {
	Tags, _ := models.GetAllTag(
		map[string]string{},
		[]string{"Id", "Name", "CreatedAt", "DeletedAt"},
		[]string{},
		[]string{},
		0,
		10,
	)

	c.Json["Tags"] = &Tags
	c.RespondJson()
}
