package View

import (
	"PrintYun/control/WebControler/Structs"
	"PrintYun/libs"
	"PrintYun/middleware"
	"PrintYun/models"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func FileOrderView(fo iris.Party)  {
	fo.Use(middleware.JWTServer)
	fo.Get("/list", ListFO)
	fo.Delete("/delete", DeleteFO)
}

// 列表出来 文件和数据库的信息
func ListFO(ctx *context.Context)  {
	var (
		FileOrderParams Structs.FileOrderParams
		FileOrderData []Structs.FileOrderData
		PrinterID string
		reponse interface{}
		db *gorm.DB
		)

	formErr := ctx.ReadURL(&FileOrderParams)
	if formErr != nil {
		reponse = libs.MakeRespon(1001, "请求参数错误")
		_, _ = ctx.JSON(reponse)
		return
	}

	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		reponse = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(reponse)
		return
	}
	PrinterID = strconv.Itoa(int(CLaim["id"].(float64)))
	db = libs.GetDB()
	data := db.Table("orders o").
		Select([]string{
			"u.id UserId",
			"o.id OrderId",
			"o.status Status",
			"o.file_name FileName",
			"o.re_name ReFileName",
			"o.updated_at UploadDate",
			"o.deleted_at DeleteDate",
		}).Joins("inner join users u on u.id=o.user_id")

	if FileOrderParams.Key != ""{
		data = data.Where("o.file_name LIKE ? OR o.code LIKE ?", "%" + FileOrderParams.Key + "%", "%" + FileOrderParams.Key + "%")
	}
	//nilTime := time.Time{}
	if FileOrderParams.DateFrom != ""{
		data = data.Where("o.created_at > ?",  FileOrderParams.DateFrom)
	}
	if FileOrderParams.DateEnd != ""{
		data = data.Where("o.created_at < ?",  FileOrderParams.DateEnd)
	}

	if PrinterID != "1" {
		data = data.
			Where("o.status = ? OR o.ds = '1' AND o.printer_id = ? OR o.deleted_at is not NULL", "1", PrinterID).
			Limit(FileOrderParams.Limit).
			Offset(FileOrderParams.Offset).
			Find(&FileOrderData)
	}else{
		data = data.
			Where("o.status = ? OR o.ds = '1' AND o.deleted_at is not NULL", "1").
			Limit(FileOrderParams.Limit).
			Offset(FileOrderParams.Offset).
			Find(&FileOrderData)
	}
	if data.RowsAffected == 0 {
		reponse = libs.MakeRespon(1000, "查询成功,暂时没有完成的订单")
		_, _ = ctx.JSON(reponse)
		return
	}
	reponseData := make(map[string]interface{})
	reponseData["items"] = FileOrderData
	reponseData["total"] = len(FileOrderData)
	reponse = libs.MakeResponData(1000, "查询成功", reponseData)
	_, _ = ctx.JSON(reponse)
	return
}

func DeleteFO(ctx *context.Context)  {
	var (
		DeleteFOData Structs.DeleteFOData
		OrderData	[]models.Order
		FileName	[]string
		reponse interface{}
		PrinterID string
		db *gorm.DB
	)
	formErr := ctx.ReadJSON(&DeleteFOData)
	if formErr != nil {
		reponse = libs.MakeRespon(1002, "请求参数错误")
		ctx.JSON(reponse)
		return
	}
	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		reponse = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(reponse)
		return
	}
	PrinterID = strconv.Itoa(int(CLaim["id"].(float64)))

	db = libs.GetDB()
	data := db.Where("id IN ? AND printer_id = ?",DeleteFOData.OrderIds,PrinterID).Find(&OrderData)
	if data.RowsAffected == 0 {
		reponse = libs.MakeRespon(1002, "检索结果不匹配, 请校验后重新提交")
		_, _ = ctx.JSON(reponse)
		return
	}else {
		for _, orders := range OrderData{
			FileName = append(FileName, orders.ReName)
		}
	}

	data = db.Where("id IN ? AND printer_id = ?", DeleteFOData.OrderIds, PrinterID).Delete(&models.Order{})
	if data.RowsAffected == 0 {
		reponse = libs.MakeRespon(1002, "当前用户下未筛查到相应的信息")
		_, _ = ctx.JSON(reponse)
		return
	}
	err = DeleteOss(FileName)
	if err != nil{
		reponse = libs.MakeRespon(1005, "删除文件失败")
		_, _ = ctx.JSON(reponse)
		return
	}
	reponse = libs.MakeRespon(1000, fmt.Sprintf("删除%s条数据成功", data.RowsAffected))
	_, _ = ctx.JSON(reponse)
	return
}

// 删除文件 OSS 非视图函数
func DeleteOss(filenames []string) error {
	var c libs.YConfig
	data := c.GetConfig()
	// 连接OSS
	client, err := oss.New(data.Oss.Endpoint, data.Oss.AccessKeyID, data.Oss.AccessKeySecret)
	if err != nil {
		return err
	}

	// 获取bucket存储空间
	bucket, err := client.Bucket("al-yun-oss")
	if err != nil {
		return err
	}

	_, err = bucket.DeleteObjects(filenames)
	if err != nil {
		return err
	}
	return nil
}

// 刪除文件 本地 非視圖函數
func DeleteFile(filenames []string)  {
	for _, filename := range filenames{
		err := os.Remove("./UploadFile/"+filename)
		if err != nil{
			log.Fatalln(err.Error())
		}
	}
}