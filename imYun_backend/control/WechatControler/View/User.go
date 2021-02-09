package View

import (
	"PrintYun/control/WechatControler/Structs"
	"PrintYun/libs"
	"PrintYun/middleware"
	"PrintYun/models"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	//"gorm.io/gorm"
)


func UserView(user iris.Party)  {
	user.Post("/login", Login)
	{
		user.Use(middleware.JWTServer)
		user.Post("/updateinfo", UpdateUserInfo)
	}
}

func Login(ctx *context.Context)  {
	var (
		db *gorm.DB
		Login Structs.Login
		reponse interface{}
		UserData models.User
		JwtToken string
		URL string
		RDB *redis.Client
		c libs.YConfig
		)

	FormErr := ctx.ReadForm(&Login)
	if FormErr != nil {
		reponse = libs.MakeRespon(1001, "表单验证错误")
		_, _ = ctx.JSON(reponse)
		return
	}

	yc := c.GetConfig()
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	URL = fmt.Sprintf(url, yc.Wechat.AppID, yc.Wechat.Secret, Login.Jscode)
	resp, err := http.Get(URL)
	if err != nil{
		reponse = libs.MakeRespon(1005, err.Error())
		_, _ = ctx.JSON(reponse)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res Structs.Code2Session
	_ = json.Unmarshal(body, &res)
	if res.Errcode != 0 {
		reponse = libs.MakeRespon(1005 , "转换失败")
		_, _ = ctx.JSON(reponse)
		return
	}

	db = libs.GetDB()
	data := db.Where(models.User{Openid: res.Openid}).First(&UserData)
	RDB = libs.GetRedisDB1()
	if data.RowsAffected == 0{
		var inviteCode string
		var userData models.User
		for {
			inviteCode = libs.GetRandomString2(10)
			data := db.Model(&models.User{}).Where(models.User{InviteCode: inviteCode}).Find(&userData)
			if data.RowsAffected == 0 {
				break
			}
		}
		UserData = models.User{
			NickName: libs.GetRandomString2(8),
			Openid: res.Openid,
			InviteCode: inviteCode,
		}
		db.Create(&UserData)
		JwtToken = libs.CreateJwt(UserData.NickName, UserData.ID)
		_, err = RDB.HSet(libs.Ctx, strconv.Itoa(int(UserData.ID)), map[string]interface{}{"JWT": JwtToken,"OpenId":res.Openid}).Result()
		if err != nil {
			reponse = libs.MakeRespon(1005, "缓存服务出现异常")
			_, _ = ctx.JSON(reponse)
			return
		}
		reponseData := make(map[string]interface{})
		reponseData["token"] = JwtToken
		reponse = libs.MakeResponData(1000, "用户第一次访问", reponseData)
		_, _ = ctx.JSON(reponse)
		return
	} else {
		JwtToken = libs.CreateJwt(UserData.NickName, UserData.ID)
		_, err = RDB.HSet(libs.Ctx, strconv.Itoa(int(UserData.ID)), map[string]interface{}{"JWT": JwtToken,"OpenId":res.Openid}).Result()
		if err != nil{
			reponse = libs.MakeRespon(1005, "缓存出现异常")
			ctx.JSON(reponse)
			return
		}
		reponseData := make(map[string]interface{})
		reponseData["token"] = JwtToken
		reponse = libs.MakeResponData(1000, "欢迎再次访问", reponseData)
		_, _ = ctx.JSON(reponse)
		return
	}
}

func UpdateUserInfo(ctx *context.Context)  {
	var (
		UpdateUserInfo 	Structs.UpdateUserInfo
		UserId 			string
		db 				*gorm.DB
		response		interface{}
	)

	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		response = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(response)
		return
	}
	UserId = strconv.Itoa(int(CLaim["id"].(float64)))

	formErr := ctx.ReadForm(&UpdateUserInfo)
	if formErr != nil {
		response = libs.MakeRespon(1001, "Form表单错误")
		_, _ = ctx.JSON(response)
		return
	}
	db = libs.GetDB()
	Result := db.Model(&models.User{}).
		Where("id = ?", UserId).
		Updates(map[string]interface{}{
			"nick_name" : UpdateUserInfo.NickName,
			"city" : UpdateUserInfo.City,
			"province" : UpdateUserInfo.Province,
			"avatar" : UpdateUserInfo.Avatar,
			"gender" : UpdateUserInfo.Gender,
		})
	if Result.Error != nil {
		response = libs.MakeRespon(1005, "出现了异常错误")
		_, _ = ctx.JSON(response)
		return
	}
	response = libs.MakeRespon(1000, "操作成功")
	_, _ = ctx.JSON(response)
	return
}