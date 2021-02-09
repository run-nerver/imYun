package libs

import (
	"PrintYun/models"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func CreatDB()  {
	var (
		err error
		dialector gorm.Dialector
		conn string
	)


	if GoEnv:=os.Getenv("GORUNENV");GoEnv=="docker"{
		conn = "root:123456@tcp(mysql:3306)/PrintYun?parseTime=True&loc=Local"
	}else{
		var c YConfig
		data := c.GetConfig()
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", data.DB.UserName, data.DB.PassWord, data.DB.Host, data.DB.Port, data.DB.Name)
	}

	//conn = "root:123456@tcp(172.16.144.131:3306)/PrintYun?parseTime=True&loc=Local"
	dialector = mysql.Open(conn)

	db, err = gorm.Open(dialector, &gorm.Config{
	})
	if err != nil {
		color.Blue(fmt.Sprintf("Mysql connect Err : %v", err))
	}
	db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})

	Migrate()
}

func GetDB() *gorm.DB {
	return db
}

func Migrate()  {
	//var Nums int64
	err := db.AutoMigrate(
		&models.User{},
		&models.Printer{},
		&models.Order{},
		&models.PrinterIcon{},
	)
	if err != nil {
		color.Blue("AutoMigrate err : %v", err)
	}
	//db.Table("printers").Find(&models.Printer{}).Count(&Nums)
	//UserNum = int(db.Model(&models.User{}).Find(&models.User{}).RowsAffected)
	//if Nums == 0 {
	//	db.Create(&models.Printer{
	//		NickName		:	"admin",
	//		Name			:	"admin",
	//		PassWord		:	HashAndSalt([]byte("123456")),
	//		Introduction 	:	"This is admin",
	//	})
	//	db.Create(&models.Printer{
	//		NickName		:	"admin",
	//		Name			:	"admin1",
	//		PassWord		:	HashAndSalt([]byte("123456")),
	//		Introduction 	:	"This is admin",
	//	})
	//
	//	for a:=0; a<=5; a++ {
	//		db.Create(&models.PrinterIcon{
	//			PrinterID: 1,
	//		})
	//		db.Create(&models.PrinterIcon{
	//			PrinterID: 2,
	//		})
	//	}
		//
		//
		//for a:=0; a<=100; a++ {
		//	db.Create(&models.Printer{
		//		NickName		:	GetRandomString2(8),
		//		Name			:	GetRandomString2(8),
		//		PassWord		:	HashAndSalt([]byte(GetRandomString2(8))),
		//		Introduction 	:	GetRandomString2(16),
		//	})
		//}
		//
		//for a:=0; a<=100; a++ {
		//	db.Create(&models.User{
		//		NickName		:	GetRandomString2(8),
		//		Introduction 	:	GetRandomString2(16),
		//		InviteCode		: 	GetRandomString2(10),
		//	})
		//}
		//
		//for a:=0; a<=1000; a++ {
		//	db.Create(&models.Order{
		//		PrinterID		:	GenerateRangeNum(1,100),
		//		UserID			:	GenerateRangeNum(1,100),
		//		Status			:	GenerateRangeNum(0,2),
		//		FileName		:	GetRandomType(8),
		//		ReName			:	GetRandomType(8),
		//		Color 			:	GenerateRangeNum(0,2),
		//		Direction 		:	GenerateRangeNum(0,2),
		//		SingleSide 		:	GenerateRangeNum(0,2),
		//		Num				:	GenerateRangeNum(1,100),
		//		Remarks			: 	GetRandomString2(8),
		//		Code			: 	GetRandomString2(4),
		//	})
		//}
	//}

	return
}