package funcs

import (
	"errors"
	"fmt"
	"time"

	"dhui.com/configs"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uid      string
	Username string
	jwt.StandardClaims
}

func Get_Token(uid string, username string) string {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	// fmt.Println("Get_Token:", uid)
	claims := &Claims{
		Uid:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(configs.Secret_Key)
	if err != nil {
		Info("token 颁发出错。 ", err)
		fmt.Println(err)
	}
	return tokenString
}

// 解密
func ParseToken(tknStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tknStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return configs.Secret_Key, nil
	})
	if err != nil {
		return nil, err
	}
	// 令牌有效
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
			return claims, nil
		}
	}
	return nil, errors.New("invalid token")
}
