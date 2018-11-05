package utils

// IsNil: 是否 nil. TODO: 有 bug，又找不到，当传入查询结果 nil 时返回居然是 false
func IsNil(val interface{}) bool {
	return val == nil
}
