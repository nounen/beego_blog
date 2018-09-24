package controllers

import (
	"beego_blog/models"
	"beego_blog/utils"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type TagController struct {
	BaseController
}

// Index 列表数据
func (c *TagController) Index() {
	tags := utils.Paging(
		orm.NewOrm().QueryTable(new(models.Tag)),
		[]string{
			"id",
			"name",
			"created_at",
			"deleted_at",
		},
		map[string]string{
			"id":         "id",
			"name":       "name",
			"created_at": "created_at",
		},
		c.getFilters(true),
		c.getPage(),
		c.getPerPage(),
	)

	c.Json["tags"] = &tags
	c.RespondJson()
}

// Store 创建数据
func (c *TagController) Store() {
	tag := c.getTagFromRequest()
	c.checkTagFromRequest(tag)

	if _, err := models.AddTag(tag); err == nil {
		c.RespondCreatedJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Show 查看数据
func (c *TagController) Show() {
	fields := []string{
		"id",
		"name",
		"created_at",
	}

	tag, err := utils.GetById(c.getTagQuery(), fields, c.getId())

	if err == nil {
		c.Json["tag"] = &tag
		c.RespondJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Update 更新数据
func (c *TagController) Update() {
	tag := c.getTagFromRequest()
	tag.Id = c.getId()
	c.checkTagFromRequest(tag)

	if err := models.UpdateTagById(tag); err == nil {
		c.RespondNoContentJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Delete 删除数据
func (c *TagController) Delete() {
	if err := models.DeleteTag(c.getId()); err == nil {
		c.RespondNoContentJson()
	} else {
		c.RespondBadJson(err)
	}
}

// getTagFromRequest 获取表单提交数据
func (c *TagController) getTagFromRequest() *models.Tag {
	tag := &models.Tag{}
	tag.CreatedAt = utils.GetNow()
	c.UnmarshalRequestJson(tag)
	return tag
}

// checkTagFromRequest 表单验证
func (c *TagController) checkTagFromRequest(Tag *models.Tag) {
	valid := validation.Validation{}
	valid.Required(Tag.Name, "Name")
	valid.MaxSize(Tag.Name, 12, "Name")

	c.RespondIfBadEntityJson(&valid)
}

// getTagQuery
func (c *TagController) getTagQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.Tag))
}
