package View

import (
	WebStruct "PrintYun/control/WebControler/Structs"
	"PrintYun/libs"
	"PrintYun/middleware"
	"PrintYun/models"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"strconv"
)


func UserView(user iris.Party)  {
	user.Get("/", func(context *context.Context) {
		context.Writef("Hello world")
	})
	user.Post("/login", Login)
	{
		//user.Use(middleware.JWT().Serve)
		user.Use(middleware.JWTServer)
		user.Get("/userinfo", UserInfo).Name = "用户信息"
		user.Post("/logout", Logout).Name = "退出登录"
	}
}

func Login(ctx *context.Context)  {
	var (
		loginForm    WebStruct.Login
		formErr      error
		PrinterModel WebStruct.PrinterInfo
		checkBool    bool
		response     interface{}
		data         *gorm.DB
		JwtToken     string
	)
	db := libs.GetDB()

	formErr = ctx.ReadJSON(&loginForm)
	if formErr != nil {
		response = libs.MakeRespon(1001, fmt.Sprintf("Form check err : %v", formErr))
		ctx.JSON(response)
		return
	}

	data = db.Model(&models.Printer{}).Select([]string{"printers.id", "printers.name", "printers.pass_word"}).Where(&models.Printer{Name: loginForm.UserName}).First(&PrinterModel)

		//	db.Model(&model.User{}).
		//Select([]string{"users.id", "users.name", "users.pass_word", "ur.role_id"}).
		//Joins("left join user_role ur on ur.user_id=users.id").
		//Where(&model.User{Name: loginForm.UserName}).
		//First(&userModel)

	if int(data.RowsAffected) == 0{
		response = libs.MakeRespon(1002, "帐号或密码错误, 请检查重试")
		ctx.JSON(response)
		return
	}
	checkBool = libs.ComparePasswords(PrinterModel.PassWord, []byte(loginForm.PassWord))
	if checkBool == false{
		response = libs.MakeRespon(1002, "帐号或密码错误, 请检查重试")
		ctx.JSON(response)
		return
	}
	JwtToken = libs.CreateJwt(PrinterModel.Name,PrinterModel.ID)
	RDB := libs.GetRedisDB0()
	baseCmd := RDB.HSet(libs.Ctx, strconv.Itoa(int(PrinterModel.ID)), []string{"JWT", JwtToken})
	if baseCmd.Err() != nil {
		response = libs.MakeRespon(1005, fmt.Sprintf("缓存出现异常 %s", baseCmd.Err()))
		_, _ = ctx.JSON(response)
		return
	}
	reponseData := make(map[string]interface{})
	reponseData["token"] = JwtToken
	response = libs.MakeResponData(1000, "验证通过", reponseData)
	ctx.JSON(response)

	return
}


func Logout(ctx *context.Context)  {
	var (
		response interface{}
		RDB      *redis.Client
		UID      string
		)
	RDB = libs.GetRedisDB0()
	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		response = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(response)
		return
	}
	UID = strconv.Itoa(int(CLaim["id"].(float64)))
	baseCmd := RDB.Del(libs.Ctx, UID)
	if baseCmd.Err() != nil {
		response = libs.MakeRespon(1002, "操作失败")
		ctx.JSON(response)
		return
	}
	response = libs.MakeRespon(1000, "操作成功")
	ctx.JSON(response)
	return
}

func UserInfo(ctx *context.Context)  {
	var (
		response     interface{}
		PrinterInfo  map[string]interface{}
		PrinterModel models.Printer
		db           *gorm.DB
		roles        []string
		)
	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		response = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(response)
		return
	}
	PrinterID := int(CLaim["id"].(float64))
	db = libs.GetDB()
	PrinterInfo = make(map[string]interface{})
	data := db.Model(&models.Printer{}).First(&PrinterModel, PrinterID).Scan(&PrinterInfo)
	if data.RowsAffected == 0{
		response = libs.MakeRespon(1005, "出现了莫名其妙的问题, 请反馈")
		ctx.JSON(response)
		return
	}
	PrinterInfo["pass_word"] = ""
	roles = append(roles, "admin")
	PrinterInfo["roles"] = roles
	response = libs.MakeResponData(1000, "查询成功", PrinterInfo)
	ctx.JSON(response)

	//Ns := WebControler.GetNs()
	//
	//fmt.Println(Ns.GetConnections())
}

//func UpdataInfo(ctx *context.Context)  {
//	var (
//		PrinterParams WebStruct.PrinterParams
//		response      interface{}
//		db *gorm.DB
//	)
//
//	formErr := ctx.ReadJSON(&PrinterParams)
//	if formErr != nil {
//		response = libs.MakeRespon(1001, fmt.Sprintf("Form check err : %v", formErr))
//		_, _ = ctx.JSON(response)
//		return
//	}
//
//	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
//	if err != nil {
//		response = libs.MakeRespon(1002, "验证异常")
//		_, _ = ctx.JSON(response)
//		return
//	}
//	PrinterID := int(CLaim["id"].(float64))
//
//	db = libs.GetDB()
//
//	updatas := make(map[string]string)
//	result := db.Table("printers p").
//		Joins("inner join printer_icons pi on p.id=pi.printer_id").w
//	if PrinterParams.Avatar != "" {
//		updatas["avatar"] = PrinterParams.Avatar
//	}
//	if PrinterParams.ImageAvatar != "" {
//		updatas["image_avater"] = PrinterParams.ImageAvatar
//	}
//	if PrinterParams.Introduction != "" {
//		updatas["introduction"] = PrinterParams.Introduction
//	}
//}