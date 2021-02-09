package Router

import (
	"PrintYun/control/WebControler"
	WebView "PrintYun/control/WebControler/View"
	"PrintYun/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func WebRouter(web iris.Party)  {
	web.UseRouter(middleware.Cors)
	web.PartyFunc("/user", WebView.UserView)
	web.PartyFunc("/order", WebView.OrderView)
	web.PartyFunc("/fo", WebView.FileOrderView)
	//web.Use(middleware.JWT().Serve)
	web.Get("/websocket", websocket.Handler(WebControler.SocketServer(), WebControler.IdGen))  // , WebControler.IdGen
}

