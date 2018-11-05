package services

// 用户登录提交信息
type LoginInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
