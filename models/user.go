package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int64      `orm:"auto" json:"id"`
	Name      string     `orm:"size(128)" json:"name"`
	Password  string     `orm:"size(255)" json:"password"`
	State     int64      `json:"state"`
	CreatedAt *time.Time `orm:"type(datetime)" json:"created_at"`
	DeletedAt *time.Time `orm:"type(datetime)" json:"deleted_at"`
}

func init() {
	orm.RegisterModel(new(User))
}
