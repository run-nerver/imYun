package View

import (
	"PrintYun/control/WechatControler/Structs"
	"PrintYun/libs"
	"PrintYun/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"strconv"
)

func OrderView(order iris.Party)  {
	order.Use(middleware.JWTServer)
	order.Get("/orderlist", OrderList)
	//order.Post("/updata", UpdateOrder)
}

//func UpdateOrder(ctx *context.Context)  {
//	var (
//		UserID string
//		response interface{}
//		db *gorm.DB
//		orderForm Structs.OrderData
//	)
//	formErr := ctx.ReadForm(&orderForm)
//	if formErr != nil{
//		response = libs.MakeRespon(1001, "表单验证错误")
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
//	UserID = strconv.Itoa(int(CLaim["id"].(float64)))
//
//	db = libs.GetDB()
//	data := db.Model(&models.Order{}).
//		Where("id = ? AND user_id = ?", orderForm.OrderId, UserID).
//		Updates(map[string]interface{}{
//			"num":orderForm.Num,
//			"color":orderForm.Color,
//			"direction":orderForm.Direction,
//			"single_side":orderForm.SingleSide,
//			"remarks":orderForm.Remarks,
//			"code":libs.GetRandomString2(4),
//			"status" : 0,
//		})
//	if data.Error != nil{
//		response = libs.MakeRespon(1005, "出现异常错误")
//		_, _ = ctx.JSON(response)
//		return
//	}
//
//	response = libs.MakeRespon(1000, "操作成功")
//	_, _ = ctx.JSON(response)
//	return
//}

func OrderList(ctx *context.Context)  {
	var (
		UserId string
		GetOrderForm Structs.GetOrderForm
		db *gorm.DB
		response interface{}
		OrderList []Structs.OrderList
	)

	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		response = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(response)
		return
	}
	UserId = strconv.Itoa(int(CLaim["id"].(float64)))

	formErr := ctx.ReadQuery(&GetOrderForm)
	if formErr != nil {
		response = libs.MakeRespon(1001, "Form表单错误")
		_, _ = ctx.JSON(response)
		return
	}
	db = libs.GetDB()
	data := db.Table("orders o").
		Select([]string{
				"o.id, " +
				"o.created_at CreateTime",
				"o.updated_at UpdateTime, " +
				"o.deleted_at DeleteTime, " +
				"o.file_name FileName, " +
				"o.re_name ReFileName, " +
				"o.paper_format PaperFormat, "+
				"o.color FileColor, " +
				"o.direction Direction, " +
				"o.single_side SingleSide, " +
				"o.num FileNum, " +
				"o.remarks Remarks, " +
				"o.code Code",
		})
	if GetOrderForm.Key != "" {
		data = data.Where("o.file_name LIKE ? OR o.code LIKE ?", "%" + GetOrderForm.Key + "%")
	}
	if GetOrderForm.DateFrom != "" {
		data = data.Where("o.created_at >= ?", GetOrderForm.DateFrom)
	}
	if GetOrderForm.DateEnd != "" {
		data = data.Where("o.created_at <= ?", GetOrderForm.DateEnd)
	}
	data = data.Where("o.user_id = ?", UserId).
		Order("o.created_at").
		Limit(GetOrderForm.Limit).
		Offset(GetOrderForm.Offset).
		Find(&OrderList)
	reponseData := make(map[string]interface{})
	reponseData["items"] = OrderList
	response = libs.MakeResponData(1000, "查询成功", reponseData)
	ctx.JSON(response)
	return
}