package services

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
)

// GetArticleTagQuery
func GetArticleTagQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.ArticleTag))
}

// InsertMultiArticleTag
func InsertMultiArticleTag(articleId int64, tagIds []int64) {

}
