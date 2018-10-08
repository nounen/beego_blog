package controllers

import (
	"beego_blog/models"
	"beego_blog/services"
	"beego_blog/utils"
	"errors"
	"fmt"
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
		services.GetArticleQuery(),
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
	// 数据解析 验证
	article := c.getArticle()
	c.checkArticle(article)

	tagIds := c.getTagIds()
	c.checkTagIds(tagIds)

	// 事务
	orm := c.BeginTransaction()

	// 文章数据
	articleId, err1 := orm.Insert(article)
	c.RespondByTransaction(orm, err1, false)

	// 标签关联
	err2 := services.InsertMultiArticleTag(orm, articleId, tagIds)
	c.RespondByTransaction(orm, err2, true)
}

// Show 查看数据
func (c *ArticleController) Show() {
	fields := []string{
		"id",
		"title",
		"created_at",
		"user_id",
	}

	article, err := utils.GetById(services.GetArticleQuery(), fields, c.getId())

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
	article := c.getArticle()
	article.Id = c.getId()
	c.checkArticle(article)

	if err := models.UpdateArticleById(article); err == nil {
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

// getTagIds
func (c *ArticleController) getTagIds() []int64 {
	ti := &services.TagIds{}
	c.UnmarshalRequestJson(&ti)
	return ti.TagIds
}

// getArticle 获取表单提交数据
func (c *ArticleController) getArticle() *models.Article {
	article := &models.Article{}
	article.CreatedAt = utils.GetNow()
	c.UnmarshalRequestJson(article)
	return article
}

// checkArticle 表单验证
func (c *ArticleController) checkArticle(Article *models.Article) {
	valid := validation.Validation{}

	valid.Required(Article.Cover, "cover")
	valid.MaxSize(Article.Cover, 255, "cover")

	valid.Required(Article.State, "state")
	//valid.Range(Article.State, 1, 3, "state")

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
