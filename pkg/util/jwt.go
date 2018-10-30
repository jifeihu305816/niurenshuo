package util

import (
	"github.com/dgrijalva/jwt-go"
	"niurenshuo/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Identifier string `json:"identifier"` //标识:手机 邮箱 用户名或第三方的唯一标识
	Credential string `json:"credential"` //密码凭证:密码或TOKEN
	jwt.StandardClaims
}

//生成TOKEN
func GenerateToken(identifier, credential string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour)

	claims := Claims{
		identifier,
		credential,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "niurenshuo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//解析TOkEN
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
