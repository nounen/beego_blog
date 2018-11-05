package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int64      `orm:"auto" json:"id"`
	Name      string     `orm:"size(128)" json:"name"`
	Password  string     `orm:"size(255)" json:"password"`
	State     int64      `orm:"int" json:"state"`
	CreatedAt *time.Time `orm:"type(datetime)" json:"created_at"`
	DeletedAt *time.Time `orm:"type(datetime)" json:"deleted_at"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
