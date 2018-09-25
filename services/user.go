package services

import (
	"beego_blog/models"
	"github.com/astaxie/beego/orm"
)

// GetUserQuery user 模型查询
func GetUserQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(models.User))
}
