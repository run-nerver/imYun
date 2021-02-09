package middleware

import (
	"PrintYun/libs"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
	"strings"
	"strconv"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte("PrintYun"), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})
	return jwtHandler
}

func JWTServer(ctx *context.Context)  {
	var (
		RDB *redis.Client
		reponse interface{}
		UrlPath []string
		UID string
		JWToken1 string
		JWToken2 string
		)
	JWToken1 = ctx.Request().Header.Get("Authorization")
	CLaim, err := libs.ParseHStoken(JWToken1)
	if err != nil {
		reponse = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(reponse)
		return
	}
	UID = strconv.Itoa(int(CLaim["id"].(float64)))

	UrlPath = strings.Split(ctx.Path(), "/")
	if UrlPath[2] == "webs" {
		RDB = libs.GetRedisDB0()
	} else if UrlPath[2] == "wechat" {
		RDB = libs.GetRedisDB1()
	}else {
		reponse = libs.MakeRespon(1002, "请重新登录")
		_, _ = ctx.JSON(reponse)
		return
	}


	JWToken2, err = RDB.HGet(libs.Ctx, UID, "JWT").Result()
	if err != nil {
		reponse = libs.MakeRespon(1002, "校验码已过期, 请重新登录")
		_, _ = ctx.JSON(reponse)
		return
	}

	if JWToken1  != JWToken2 {
		reponse = libs.MakeRespon(1002, "授权码不匹配, 请重新登录")
		_, _ = ctx.JSON(reponse)
		return
	}

	ctx.Next()
}