package controllers

import (
	"beego_blog/models"
	"beego_blog/services"
	"beego_blog/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type ArticleController struct {
	BaseController
}

// Index 列表数据
func (c *ArticleController) Index() {
	fields := []string{
		"id",
		"title",
		"content",
		"cover",
		"state",
		"created_at",
		"user_id",
	}

	filtersMap := map[string]string{
		"id":         "id",
		"title":      "title",
		"created_at": "created_at",
	}

	articles := utils.Paging(
		c.getArticleQuery(),
		fields,
		filtersMap,
		c.getFilters(true),
		c.getPage(),
		c.getPerPage(),
	)

	c.Json["articles"] = &articles
	c.RespondJson()
}

// Store 创建数据
func (c *ArticleController) Store() {
	article := c.getArticleFromRequest()
	article.UserId = 1
	c.checkArticleFromRequest(article)

	tagIds := c.getTagIdsFromRequest()
	c.checkTagIds(tagIds)

	if articleId, err := models.AddArticle(article); err == nil {
		services.InsertMultiArticleTag(articleId, tagIds)
		c.RespondCreatedJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Show 查看数据
func (c *ArticleController) Show() {
	fields := []string{
		"id",
		"title",
		"created_at",
		"user_id",
	}

	article, err := utils.GetById(c.getArticleQuery(), fields, c.getId())

	if err == nil {
		userId := article["user_id"].(int64)
		articleId := article["id"].(int64)

		article["user"], _ = utils.GetById(services.GetUserQuery(), []string{"id", "name"}, userId)
		article["tags"] = services.GetTagsByArticleId(articleId)
		article["tag_ids"] = services.GetTagIdsByArticleId(articleId)

		c.Json["article"] = &article
		c.RespondJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Update 更新数据
func (c *ArticleController) Update() {
	Article := c.getArticleFromRequest()
	c.checkArticleFromRequest(Article)
	Article.Id = c.getId()

	if err := models.UpdateArticleById(Article); err == nil {
		c.RespondNoContentJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Delete 删除数据
func (c *ArticleController) Delete() {
	if err := models.DeleteArticle(c.getId()); err == nil {
		c.RespondNoContentJson()
	} else {
		c.RespondBadJson(err)
	}
}

// getTagIdsFromRequest
func (c *ArticleController) getTagIdsFromRequest() []int64 {
	ti := &services.TagIds{}
	c.UnmarshalRequestJson(&ti)
	return ti.TagIds
}

// getArticleFromRequest 获取表单提交数据
func (c *ArticleController) getArticleFromRequest() *models.Article {
	article := &models.Article{}
	article.CreatedAt = utils.GetNow()
	c.UnmarshalRequestJson(article)
	return article
}

// checkArticleFromRequest 表单验证
func (c *ArticleController) checkArticleFromRequest(Article *models.Article) {
	valid := validation.Validation{}

	valid.Required(Article.Cover, "cover")
	valid.MaxSize(Article.Cover, 255, "cover")

	valid.Required(Article.State, "state")
	valid.Range(Article.State, 1, 3, "state")

	valid.Required(Article.Title, "title")
	valid.MaxSize(Article.Title, 200, "title")

	valid.Required(Article.Content, "content")

	c.RespondIfBadEntityJson(&valid)
}

// checkTagIds
func (c *ArticleController) checkTagIds(tagIds []int64) {
	for _, tagId := range tagIds {
		_, err := models.GetTagById(tagId)

		if err != nil {
			msg := fmt.Sprintf("标签ID %d 不存在", tagId)
			c.RespondBadJson(errors.New(msg))
		}
	}
}

// getArticleQuery
func (c *ArticleController) getArticleQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.Article))
}
