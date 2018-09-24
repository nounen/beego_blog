package controllers

import (
	"beego_blog/models"
	"beego_blog/utils"
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
	Article := c.getArticleFromRequest()
	c.checkArticleFromRequest(Article)

	if _, err := models.AddArticle(Article); err == nil {
		c.RespondCreatedJson()
	} else {
		c.RespondBadJson(err)
	}
}

// Show 查看数据
func (c *ArticleController) Show() {
	if Article, err := models.GetArticleById(c.getId()); err == nil {
		c.Json["Article"] = &Article
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

// getArticleFromRequest 获取表单提交数据
func (c *ArticleController) getArticleFromRequest() *models.Article {
	Article := &models.Article{}
	Article.CreatedAt = utils.GetNow()
	c.UnmarshalRequestJson(Article)
	return Article
}

// checkArticleFromRequest 表单验证
func (c *ArticleController) checkArticleFromRequest(Article *models.Article) {
	valid := validation.Validation{}
	//valid.Required(Article.Title, "Title")
	//valid.MaxSize(Article.Title, 12, "Title")

	c.RespondIfBadEntityJson(&valid)
}

// getArticleQuery
func (c *ArticleController) getArticleQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.Article))
}
