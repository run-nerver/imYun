package View

import (
	Ws "PrintYun/control/WebControler"
	"PrintYun/control/WechatControler/Structs"
	"PrintYun/libs"
	"PrintYun/models"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/neffos"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Upload(ctx *context.Context)  {
	var (
		keys int
		response interface{}
		UploadForm Structs.UploadForm
		db *gorm.DB
		OrderData models.Order
		UserID int
		Count int64
		filename string
		)
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()
	err := ctx.ReadForm(&UploadForm)
	if err != nil {
		response = libs.MakeRespon(1001, "Form表单错误")
		_, _ = ctx.JSON(response)
		return
	}

	err = ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		response = libs.MakeRespon(1001, "Upload表单错误")
		_,_ = ctx.JSON(response)
		return
	}
	form := ctx.Request().MultipartForm
	files := form.File["file"]
	db = libs.GetDB()

	fmt.Println(time.Now().Format("20060102150405"))
	CLaim, err := libs.ParseHStoken(ctx.Request().Header.Get("Authorization"))
	if err != nil {
		response = libs.MakeRespon(1002, "验证异常")
		_, _ = ctx.JSON(response)
		return
	}
	UserID = int(CLaim["id"].(float64))
	_ = db.Where(&models.Order{UserID: UserID}).Find(&OrderData).Count(&Count)
	if int(Count) == 0 {
		keys = 0
	}else{
		keys=int(Count)+1
	}
	for _, file := range files {
		types := strings.Split(file.Filename, ".")

		filename = time.Now().Format("20060102150405") + strconv.Itoa(UserID) + strconv.Itoa(keys)+ "." + types[len(types)-1]

		OrderData = models.Order{
			UserID: UserID,
			PrinterID: UploadForm.PrinterID,
			FileName: UploadForm.FileName,
			ReName: filename,
			Num: UploadForm.Num,
			Color: UploadForm.Color,
			SingleSide: UploadForm.SingleSide,
			Direction: UploadForm.Direction,
			Remarks: UploadForm.Remarks,
			PaperFormat: UploadForm.PaperFormat,
			Status: 0,
			Code: libs.GetRandomString2(4),
		}
		creatErr := db.Create(&OrderData)
		if creatErr.Error != nil {
			response = libs.MakeRespon(1005, fmt.Sprintf("文件上传失败: %s\nErr : %s", file.Filename, creatErr.Error))
			_, _ = ctx.JSON(response)
			return
		}


		if GoEnv:=os.Getenv("GORUNENV");GoEnv!="docker"{
			_, err = saveUploadedFile(file, filename, "./UploadFile/" + time.Now().Format("200601"))
			fmt.Println(os.Getwd())
			if err != nil {
				response = libs.MakeRespon(1005, fmt.Sprintf("文件上传失败: %s\nErr : %s", file.Filename, err.Error()))
				_, _ = ctx.JSON(response)
				return
			}

			var c libs.YConfig
			data := c.GetConfig()
			if data.Oss.Enable != "false" {
				_, err = UploadOss(filename)
				if err != nil {
					response = libs.MakeRespon(1005, fmt.Sprintf("文件转储失败: %s\nErr : %s", file.Filename, err.Error()))
					_, _ = ctx.JSON(response)
					return
				}
			}
		}else{
			filePathDate := time.Now().Format("200601")
			if _, err = os.Stat("/dist/UploadFile/" + filePathDate); os.IsNotExist(err) {
				// 必须分成两步
				// 先创建文件夹
				os.Mkdir("/dist/UploadFile/" + filePathDate, 0777)
				// 再修改权限
				os.Chmod("/dist/UploadFile/" + filePathDate, 0777)
			}

			_, err = saveUploadedFile(file, filename, "/dist/UploadFile/" + filePathDate)
			fmt.Println(os.Getwd())
			if err != nil {
				response = libs.MakeRespon(1005, fmt.Sprintf("文件上传失败: %s\nErr : %s", file.Filename, err.Error()))
				_, _ = ctx.JSON(response)
				return
			}
		}

		// 这里 可删 : 删除本地数据
		//os.Remove("./UploadFile/"+filename)
		wsData := Structs.OrderList{}
		wsData.Id = int(OrderData.ID)
		wsData.FileName = OrderData.FileName
		wsData.ReFileName = OrderData.ReName
		wsData.Code = OrderData.Code
		wsData.PaperFormat = OrderData.PaperFormat
		wsData.Remarks = OrderData.Remarks
		wsData.SingleSide = OrderData.SingleSide
		wsData.FileNum = OrderData.Num
		wsData.NickName = CLaim["username"].(string)
		wsData.Direction = OrderData.Direction
		wsData.FileColor = OrderData.Color
		wsData.UpdateTime = OrderData.UpdatedAt
		data, _ := json.Marshal(wsData)
		Ns := Ws.GetNs()
		RDB := libs.GetRedisDB2()
		WSID := RDB.HGet(libs.Ctx,strconv.Itoa(UploadForm.PrinterID), "WSID")

		Ns.Broadcast(nil, neffos.Message{
			Namespace: "default",
			Event:     "Notify",
			Body:      data,
			To: WSID.Val(),
		})
		keys++
	}
	response = libs.MakeRespon(1000, "文件上传成功")
	_, _ = ctx.JSON(response)
	return
}

func saveUploadedFile(fh *multipart.FileHeader, filename string, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()
	out, err := os.OpenFile(filepath.Join(destDirectory, filename), os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, src)
}

//* 删除 这里 //

func UploadOss(filename string) (int, error) {
	var c libs.YConfig
	data := c.GetConfig()

	// 连接OSS
	client, err := oss.New(data.Oss.Endpoint, data.Oss.AccessKeyID, data.Oss.AccessKeySecret)
	if err != nil {
		return 1, err
	}

	// 获取bucket存储空间
	bucket, err := client.Bucket("al-yun-oss")
	if err != nil {
		return 1, err
	}

	err = bucket.PutObjectFromFile("./UploadFile/"+filename[0:6]+"/"+filename, "./UploadFile/"+filename[0:6]+"/"+filename)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

// 删除 这里 *//