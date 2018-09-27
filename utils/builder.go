package utils

import (
	"github.com/astaxie/beego/orm"
)

// GetFirst
// @item 必须传递结构体的地址 eg： &tag
func GetFirst(qb orm.QueryBuilder, item interface{}, id int64) error {
	queryStr := qb.
		Limit(1).
		String()

	err := orm.
		NewOrm().
		Raw(queryStr, id).
		QueryRow(item)

	return err
}

// GetQueryBuilder 获取 QueryBuilder 对象
func GetQueryBuilder() (qb orm.QueryBuilder) {
	queryBuilder, _ := orm.NewQueryBuilder("mysql")
	return queryBuilder
}
