package control

import (
	WebRouter "PrintYun/control/WebControler/Router"
	WechatRouter "PrintYun/control/WechatControler/Router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func Service(application *iris.Application)  {
	application.PartyFunc("/v1", func(Route router.Party) {
		Route.PartyFunc("/webs", WebRouter.WebRouter)
		Route.PartyFunc("/wechat", WechatRouter.WechatRouter)
	})
}