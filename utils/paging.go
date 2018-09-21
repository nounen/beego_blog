package utils

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

// Filters 查询条件结构
type Filters struct {
	Equals   map[string]string
	Likes    map[string]string
	Betweens map[string][]string
	Ins      map[string][]string
	Orders   map[string]string
}

// Paging 分页数据结构
type Page struct {
	Page      int64
	PerPage   int64
	TotalPage int64
	PrevPage  int64
	NextPage  int64
	Total     int64
	List      interface{}
}

// @query orm.QuerySeter		 	// 查询初始化
// @fields []string	 			// 查询字段
// @fieldMap map[string]string	// 允许过滤字段与别名映射
// @filters *Filters	 			// 过滤条件
// @page int64			 			// 当前页码
// @perPage int64		 			// 每页数量
// TODO： per_page 不能超过 200， page 不能超过总页数等等处理
// Paging 分页处理工具
func Paging(
	query orm.QuerySeter,
	fields []string,
	fieldMap map[string]string,
	filters *Filters,
	page int64,
	perPage int64,
) *Page {
	// fieldMap 查询字段映射： 1.没有在这里的字段不允许查询; 2. 表别名时字段别名映射
	// TODO: 别名无效
	// query: =
	for equalKey, equalValue := range filters.Equals {
		fieldAlias, ok := fieldMap[equalKey]

		if ok {
			query = query.Filter(fieldAlias, equalValue)
		}
	}

	// query: like
	for likeKey, likeValue := range filters.Likes {
		fieldAlias, ok := fieldMap[likeKey]

		if ok {
			query = query.Filter(fieldAlias+"__icontains", likeValue)
		}
	}

	// query: between
	for betweenKey, betweenValue := range filters.Betweens {
		fieldAlias, ok := fieldMap[betweenKey]

		if ok {
			query = query.Filter(fieldAlias+"__gte", betweenValue[0])
			query = query.Filter(fieldAlias+"__lte", betweenValue[1])
		}
	}

	// query: in
	for inKey, inValue := range filters.Ins {
		fieldAlias, ok := fieldMap[inKey]

		if ok {
			query = query.Filter(fieldAlias+"__in", inValue)
		}
	}

	// 数量统计 total
	total, _ := query.Count()

	// query: order by
	for orderKey, orderValue := range filters.Orders {
		fieldAlias, ok := fieldMap[orderKey]

		if ok {
			if strings.ToLower(orderValue) == "desc" {
				query = query.OrderBy("-" + fieldAlias)
			} else {
				query = query.OrderBy(fieldAlias)
			}
		}
	}

	// 查询
	list := []orm.Params{}
	query.
		Limit(perPage, (page-1)*perPage).
		Values(&list, fields...)

	// 分页计算
	// 总页数
	totalPage := total / perPage
	if total%perPage > 0 {
		totalPage = total/perPage + 1
	}

	// 上一页, 下一页
	prevPage := int64(0)
	nextPage := int64(0)

	if totalPage > page {
		nextPage = page + 1
	}

	if page > 1 {
		prevPage = page - 1
	}

	// 分页数据
	Paging := Page{
		Page:      page,
		PerPage:   perPage,
		TotalPage: totalPage,
		PrevPage:  prevPage,
		NextPage:  nextPage,
		Total:     total,
		List:      list,
	}

	return &Paging
}
