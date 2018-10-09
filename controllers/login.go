package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type LoginController struct {
	BaseController
}

// Login 用户登录，得到 jwt token
func (c *LoginController) Login() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}

func (c *LoginController) Logout() {}
