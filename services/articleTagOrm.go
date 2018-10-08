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

// DeleteByArticleId 根据文章删除关联标签
func DeleteByArticleId(o orm.Ormer, articleId int64) {
	o.QueryTable(new(models.ArticleTag)).
		Filter("article_id", articleId).
		Delete()
}
