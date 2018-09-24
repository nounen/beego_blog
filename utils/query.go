package utils

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

// GetTagById retrieves Tag by Id. Returns error if
// Id doesn't exist
func GetById(
	query orm.QuerySeter,
	fields []string,
	id int64,
) (item interface{}, err error) {
	var list []orm.Params

	count, err := query.
		Filter("Id", id).
		Limit(1, 0).
		Values(&list, fields...)

	// 查无数据
	if count == 0 {
		return nil, errors.New("数据不存在")
	}

	// 查询结果转 key 转 小写下划线
	m := map[string]interface{}{}

	for k, v := range list[0] {
		m[snakeString(k)] = v
	}

	return m, err
}
