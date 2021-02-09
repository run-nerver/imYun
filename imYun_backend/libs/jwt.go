package libs

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"time"
)


// 	加密
func CreateJwt(UserName string,UserID uint) string {
	token := jwtmiddleware.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": UserName,
		"id": UserID,
		"iss":"Iris",
		"iat":time.Now().Unix(),
		"exp":time.Now().Add(time.Second*60*60*24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("PrintYun"))
	return tokenString
}


// 解密
func ParseHStoken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("PrintYun"), nil
	})
	if err != nil {
		fmt.Println("HS256的token解析错误，err:", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ParseHStoken:claims类型转换失败")
		return nil, err
	}
	return claims, err
}