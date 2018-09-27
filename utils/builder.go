package utils

import "github.com/astaxie/beego/orm"

func GetFirst() {

}

// GetQueryBuilder 获取 QueryBuilder 对象
func GetQueryBuilder() (qb orm.QueryBuilder) {
	queryBuilder, _ := orm.NewQueryBuilder("mysql")
	return queryBuilder
}
