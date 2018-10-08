package services

import (
	"beego_blog/utils"
	"github.com/astaxie/beego/orm"
)

// GetTagsByArticleId 获取文章 tags
func GetTagsByArticleId(articleId int64) []ArticleTag {
	fields := []string{
		"tag.id AS id",
		"tag.name AS name",
	}

	queryString := utils.GetQueryBuilder().
		Select(fields...).
		From("tag").
		InnerJoin("article_tag").
		On("article_tag.tag_id = tag.id").
		Where("article_id = ?").
		String()

	// 执行SQL语句
	var tags []ArticleTag
	orm.NewOrm().
		Raw(queryString, articleId).
		QueryRows(&tags)

	return tags
}

// GetTagIdsByArticleId 获取文章 tagIds
func GetTagIdsByArticleId(articleId int64) []int {
	fields := []string{
		"tag_id AS id",
	}

	queryString := utils.GetQueryBuilder().
		Select(fields...).
		From("article_tag").
		Where("article_id = ?").
		String()

	// 执行SQL语句
	var ids []int
	orm.NewOrm().
		Raw(queryString, articleId).
		QueryRows(&ids)

	return ids
}
