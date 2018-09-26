package models

import (
	"github.com/astaxie/beego/orm"
)

type ArticleTag struct {
	Id        int64 `orm:"auto" json:"id"`
	ArticleId int64 `json:"article_id"`
	TagId     int64 `json:"tag_id"`
}

func init() {
	orm.RegisterModel(new(ArticleTag))
}
