package controllers

import (
	"beego_blog/models"
	"time"
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

// Store 创建数据
func (c *TagController) Store() {
	if _, err := models.AddTag(c.getTagFromRequest()); err == nil {
		c.RespondCreatedJson()
	} else {
		c.RespondBadJson(err)
	}
}

// getTagFromRequest 获取表单提交数据
func (c *TagController) getTagFromRequest() *models.Tag {
	Tag := &models.Tag{}
	c.UnmarshalRequestJson(Tag)
	Tag.CreatedAt = time.Now()
	return Tag
}
