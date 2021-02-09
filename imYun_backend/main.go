package main

import (
	"PrintYun/control"
	"PrintYun/libs"
	"os"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
)


var app *iris.Application

func main()  {
	var GoEnv string
	if GoEnv=os.Getenv("GORUNENV");GoEnv=="docker"{
		app = CreateApp()
		app.Run(iris.Addr(os.Getenv("GORUNADDR")+":"+os.Getenv("GORUNPORT")))
	}else{
		var c libs.YConfig
		data := c.GetConfig()
		app = CreateApp()
		app.Run(iris.Addr(data.Host+":"+data.Port))
	}
}

func CreateApp() *iris.Application {
	app = iris.New()


	// 使用内置的文档生成器
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       false,
		DocTitle: "Iris",
		DocPath:  "ApiList/apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New()) 		// 开启接口生成文档
	//数据库 + 注册路由
	RegistALlAboutDB(app)

	return app
}

func RegistALlAboutDB(application *iris.Application)  {
	libs.CreatDB()
	libs.CreateRedisDB()
	control.Service(application)
}
