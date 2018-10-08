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
func InsertMultiArticleTag(o orm.Ormer, articleId int64, tagIds []int64) error {
	articleTags := []models.ArticleTag{}

	for _, tagId := range tagIds {
		articleTag := models.ArticleTag{
			ArticleId: articleId,
			TagId:     tagId,
		}

		articleTags = append(articleTags, articleTag)
	}

	_, err := o.InsertMulti(10, articleTags)

	return err
}
