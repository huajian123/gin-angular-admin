package common

import (
	"ant-admin/gin-angular-admin/model/entity/sysEntity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
token 由3部分组成，第一部分是加密协议
中间一部分储存的用户信息
第三部分是前面两部分，加上 加盐，hash的值

*/

// jwt加密迷药
var jwtKey = []byte("a_secret_crect ")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 发送token,再登陆成功后
func ReleaseToken(user sysEntity.User) (string, error) {
	// 过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 发放时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "huajian",
			Subject:   "user token",
		},
	}

	// 用JWT 秘药来生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	// 生成错误，则将错误返回
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
