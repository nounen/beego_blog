package services

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
)

// GetArticleQuery
func GetArticleQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.Article))
}
