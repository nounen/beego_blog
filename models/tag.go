package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id        int64     `orm:"auto" json:"id"`
	Name      string    `orm:"size(128)" json:"name"`
	CreatedAt time.Time `orm:"type(datetime)" json:"created_at"`
	DeletedAt time.Time `orm:"type(datetime)" json:"deleted_at"`
}

func init() {
	orm.RegisterModel(new(Tag))
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
