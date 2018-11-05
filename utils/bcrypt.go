package utils

import "golang.org/x/crypto/bcrypt"

// Encryption 字符串加密
func Encryption(str string) (hashStr string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash[:])
}

// @passwordNow 当前密码(用户输入明文密码)
// @passwordHashed 数据库里密码(加密后)
// CheckEncryption 密码验证
func CheckEncryption(passwordHashed, passwordNow string) error {
	// 正确密码验证
	result := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(passwordNow))
	return result
}
