package services

import (
	"beego_blog/utils"
	"github.com/astaxie/beego/orm"
)

// GetTagsByArticleId 获取文章 tags
func GetTagsByArticleId(articleId int64) []orm.Params {
	fields := []string{
		"tag.id",
		"tag.name",
	}

	queryString := utils.GetQueryBuilder().
		Select(fields...).
		From("tag").
		InnerJoin("article_tag").
		On("article_tag.tag_id = tag.id").
		Where("article_id = ?").
		String()

	// 执行SQL语句
	// TODO: 查询结果都是 【字符串】 如何处理？
	var list []orm.Params
	orm.NewOrm().
		Raw(queryString, articleId).
		Values(&list)

	return list
}
