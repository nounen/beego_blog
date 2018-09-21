package controllers

import (
	"beego_blog/models"
	"beego_blog/utils"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type ArticleController struct {
	BaseController
}

// Index 列表数据
func (c *ArticleController) Index() {
	Articles := utils.Paging(
		orm.NewOrm().QueryTable(new(models.Article)),
		[]string{
			"Id",
			"Name",
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

	c.Json["Articles"] = &Articles
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
	c.UnmarshalRequestJson(Article)
	Article.CreatedAt = time.Now()
	return Article
}

// checkArticleFromRequest 表单验证
func (c *ArticleController) checkArticleFromRequest(Article *models.Article) {
	valid := validation.Validation{}
	valid.Required(Article.Name, "Name")
	valid.MaxSize(Article.Name, 12, "Name")

	c.RespondIfBadEntityJson(&valid)
}
