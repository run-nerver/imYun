package View

import (
	"PrintYun/control/WebControler/Structs"
	"PrintYun/libs"
	"PrintYun/middleware"
	"PrintYun/models"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func OrderView(order iris.Party)  {
	order.Use(middleware.JWTServer)
	order.Get("/orderlist", OrderList)
	order.Delete("/odelete", OrderDelete)
	order.Post("/confirm", Confirm)
	order.Get("/getLocalFile", DownloadOrderFileLocal)
}

func OrderList(ctx *context.Context)  {
	var (
		db *gorm.DB
		OrderData []Structs.OrderList
		PrinterID int
		reponse interface{}
		orderParams Structs.OrderParams
		COUNT int64
		)

	formErr := ctx.ReadURL(&orderParams)
	if formErr != nil {
		reponse = libs.MakeRespon(1001, "请求参数错误")
		ctx.JSON(reponse)
		return
	}


	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		reponse = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(reponse)
		return
	}
	PrinterID = int(CLaim["id"].(float64))

	db = libs.GetDB()
	data := db.Table("orders o").
		Select([]string{
				"o.id, " +
				"o.created_at CreatTime, " +
				"o.deleted_at DeleteTime, " +
				"u.nick_name NickName, " +
				"o.file_name FileName, " +
				"o.re_name ReFileName, " +
				"o.paper_format PaperFormat, "+
				"o.color FileColor, " +
				"o.direction Direction, " +
				"o.single_side SingleSide, " +
				"o.num FileNum, " +
				"o.remarks Remarks, " +
				"o.code Code",
		}).Joins("inner join users u on u.id=o.user_id")
	counts := db.Table("orders o").
		Joins("inner join users u on u.id=o.user_id")
	if orderParams.Key != ""{
		data = data.Where("o.file_name LIKE ? OR o.code LIKE ?", "%" + orderParams.Key + "%", "%" + orderParams.Key + "%")
		counts = counts.Where("o.file_name LIKE ? OR o.code LIKE ?", "%" + orderParams.Key + "%", "%" + orderParams.Key + "%")
	}
	//nilTime := time.Time{}
	if orderParams.DateFrom != ""{
		data = data.Where("o.created_at > ?",  orderParams.DateFrom)
		counts = counts.Where("o.created_at > ?",  orderParams.DateFrom)
	}
	if orderParams.DateEnd != ""{
		data = data.Where("o.created_at < ?",  orderParams.DateEnd)
		counts = counts.Where("o.created_at < ?",  orderParams.DateEnd)
	}

	if PrinterID != 1{
		data = data.Where("o.status = ? AND o.printer_id = ? AND o.ds='0'", strconv.Itoa(0), strconv.Itoa(PrinterID)).
			Limit(orderParams.Limit).
			Offset(orderParams.Offset).
			Order("o.updated_at desc").
			Find(&OrderData)
		counts = counts.Where("o.status = ? AND o.printer_id = ? AND o.ds='0'", strconv.Itoa(0), strconv.Itoa(PrinterID)).Count(&COUNT)
		fmt.Println(COUNT)
		if data.RowsAffected == 0 {
			reponseData := make(map[string]interface{})
			reponseData["total"] = int(COUNT)
			reponseData["items"] = OrderData
			reponse = libs.MakeResponData(1000, "查询成功", reponseData)
			_, _ = ctx.JSON(reponse)
			return
		}
		reponseData := make(map[string]interface{})
		reponseData["items"] = OrderData
		reponseData["total"] = int(COUNT)
		reponse = libs.MakeResponData(1000, "查询成功", reponseData)
		ctx.JSON(reponse)
		return
	}else {
		data = data.Where("o.status = ? AND o.ds='0'", strconv.Itoa(0)).
			Limit(orderParams.Limit).
			Offset(orderParams.Offset).
			Order("o.created_at").
			Find(&OrderData)
		counts = counts.Where("o.status = ? AND o.ds='0'", strconv.Itoa(0)).Count(&COUNT)
		fmt.Println(COUNT)
		if data.RowsAffected == 0 {
			reponseData := make(map[string]interface{})
			reponseData["total"] = int(COUNT)
			reponseData["items"] = OrderData
			reponse = libs.MakeResponData(1000, "查询成功", reponseData)
			_, _ = ctx.JSON(reponse)
			return
		}
		reponseData := make(map[string]interface{})
		reponseData["items"] = OrderData
		reponseData["total"] = int(COUNT)
		reponse = libs.MakeResponData(1000, "查询成功", reponseData)
		_, _ = ctx.JSON(reponse)
		return
	}
}

func OrderDelete(ctx *context.Context)  {
	var (
		OrderData 	Structs.OrderData
		reponse 	interface{}
		db			*gorm.DB
		dbErr		*gorm.DB
		)

	formErr := ctx.ReadJSON(&OrderData)
	if formErr != nil {
		reponse = libs.MakeRespon(1001, "表单错误")
		_, _ = ctx.JSON(reponse)
		return
	}
	db = libs.GetDB()
	if len(OrderData.OrderId) != 1 {
		dbErr = db.Model(&models.Order{}).Where("id IN ?", OrderData.OrderId).Update("ds","1")
	}else{
		dbErr = db.Model(&models.Order{}).Where("id = ?", OrderData.OrderId).Update("ds","1")
	}

	if dbErr.Error != nil {
		reponse = libs.MakeRespon(1005, "数据溢出")
		_, _ = ctx.JSON(reponse)
		return
	}
	reponse = libs.MakeRespon(1000, "操作成功")
	_, _ = ctx.JSON(reponse)
	return
}

func Confirm(ctx *context.Context)  {
	var (
		OrderData 	Structs.OrderData
		reponse 	interface{}
		db			*gorm.DB
	)

	formErr := ctx.ReadForm(&OrderData)
	if formErr != nil {
		reponse = libs.MakeRespon(1001, "表单错误")
		_, _ = ctx.JSON(reponse)
		return
	}
	db = libs.GetDB()
	dbErr := db.Model(&models.Order{}).Where("id = ?", OrderData.OrderId).Update("status", 1)
	if dbErr.Error != nil {
		reponse = libs.MakeRespon(1005, "数据溢出")
		_, _ = ctx.JSON(reponse)
		return
	}
	reponse = libs.MakeRespon(1000, "操作成功")
	_, _ = ctx.JSON(reponse)
	return
}

func DownloadOrderFileLocal(ctx *context.Context)  {
	var (
		GetLocalPamars Structs.LocalFile
		db *gorm.DB
		OrderData models.Order
		response interface{}
		filePath string
	)

	formErr := ctx.ReadQuery(&GetLocalPamars)
	if formErr != nil {
		response = libs.MakeRespon(1001, "表单错误")
		_, _ = ctx.JSON(response)
		return
	}
	if GoEnv:=os.Getenv("GORUNENV");GoEnv!="docker"{
		filePath = "./UploadFile/" + GetLocalPamars.ReFileName[0:6] + "/" + GetLocalPamars.ReFileName
	}else {
		filePath = "/dist/UploadFile/" + GetLocalPamars.ReFileName[0:6] + "/" + GetLocalPamars.ReFileName
	}
	db = libs.GetDB()
	result := db.Where("re_name = ?", GetLocalPamars.ReFileName).Find(&OrderData)
	if result.Error != nil{
		response = libs.MakeRespon(1005, "异常错误")
		_, _ = ctx.JSON(response)
		return
	}
	ctx.SendFile(filePath, OrderData.FileName)
}