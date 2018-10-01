package services

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
)

// GetArticleTagQuery
func GetArticleTagQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.ArticleTag))
}

// InsertMultiArticleTag 文章标签关联
func InsertMultiArticleTag(articleId int64, tagIds []int64) {
	articleTags := []models.ArticleTag{}

	for _, tagId := range tagIds {
		articleTag := models.ArticleTag{
			ArticleId: articleId,
			TagId: tagId,
		}

		articleTags = append(articleTags,articleTag)
	}

	orm.NewOrm().InsertMulti(10, articleTags)
}
