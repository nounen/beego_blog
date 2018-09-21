package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Tag struct {
	Id int64
	//Id        int64     `orm:"auto"`
	Name      string    `orm:"size(128)"`
	CreatedAt time.Time `orm:"type(datetime)"`
	DeletedAt time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Tag))
}

// Index orm 使用探索
func Index() interface{} {
	tags := []orm.Params{}

	fields := []string{
		"Id",
		"Name",
	}

	filterName := "tag"

	orderId := "desc"

	// 初始化查询
	query := orm.NewOrm().
		QueryTable(new(Tag))

	// 过滤条件
	if len(filterName) > 2 {
		query = query.Filter("name__icontains", filterName)
	}

	// 分页： total
	//total, _ := query.Count()

	// 排序
	if orderId == "desc" {
		query = query.OrderBy("-id")
	}

	// 分页 和 赋值
	query.
		Limit(10, 0).
		Values(&tags, fields...)

	return tags
}

// AddTag insert a new Tag into database and returns
// last inserted Id on success.
func AddTag(m *Tag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTagById retrieves Tag by Id. Returns error if
// Id doesn't exist
func GetTagById(id int64) (v *Tag, err error) {
	o := orm.NewOrm()
	v = &Tag{Id: id}
	if err = o.QueryTable(new(Tag)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateTag updates Tag by Id and returns error if
// the record to be updated doesn't exist
func UpdateTagById(m *Tag) (err error) {
	o := orm.NewOrm()
	v := Tag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTag deletes Tag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTag(id int64) (err error) {
	o := orm.NewOrm()
	v := Tag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
