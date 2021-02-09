package Router

import (
	WechatView "PrintYun/control/WechatControler/View"
	"PrintYun/middleware"
	"github.com/kataras/iris/v12"
)

func WechatRouter(wechat iris.Party)  {
	wechat.UseRouter(middleware.Cors)
	wechat.PartyFunc("/user", WechatView.UserView)
	wechat.PartyFunc("/order", WechatView.OrderView)
	wechat.Use(middleware.JWTServer)
	wechat.Post("/Upload", WechatView.Upload)
}