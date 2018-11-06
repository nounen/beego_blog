package utils

import (
	"errors"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 来源: https://github.com/EDDYCJY/blog/blob/master/golang/gin/2018-02-16-Gin%E5%AE%9E%E8%B7%B5-%E8%BF%9E%E8%BD%BD%E4%BA%94-%E4%BD%BF%E7%94%A8JWT%E8%BF%9B%E8%A1%8C%E8%BA%AB%E4%BB%BD%E6%A0%A1%E9%AA%8C.md

var jwtSecret = []byte("jwt-secret-xxx")

type Claims struct {
	Id   int64  `json:"id"`
	Name string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成 jwt token
func GenerateToken(id int64, name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
		name,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken token解析
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetUserId 根据token 解析出登录的用户 id
func GetUserId(ctx *context.Context) (id int64, err error) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		return 0, errors.New("请先登录(没有token)")
	}

	claims, _ := ParseToken(token)
	if claims == nil {
		return 0, errors.New("登录失效(请重新登录)")
	}

	return claims.Id, err
}
