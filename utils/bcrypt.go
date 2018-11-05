package utils

import "golang.org/x/crypto/bcrypt"

// encryption 字符串加密
func Encryption(str string) (hashStr string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash[:])
}
