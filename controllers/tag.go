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

// Show 查看数据
func (c *TagController) Show() {
	if Tag, err := models.GetTagById(c.getId()); err == nil {
		c.Json["Tag"] = &Tag
		c.RespondJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Update 更新数据
func (c *TagController) Update() {
	Tag := c.getTagFromRequest()
	Tag.Id = c.getId()

	if err := models.UpdateTagById(Tag); err == nil {
		c.RespondNoContentJson()
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
