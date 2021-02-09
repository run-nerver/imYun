package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	PrinterID	int			`gorm:"int(32);NOT NULL" json:"printerId" comment:"打印店的id"`
	UserID     	int			`gorm:"int(32);NOT NULL" json:"userId" comment:"用户的id"`
	FileName	string		`gorm:"type:varchar(128);NOT NULL" json:"fileName" comment:"文件名"`
	ReName 		string		`gorm:"type:varchar(128);NOT NULL" json:"reName" comment:"存储文件名"`
	DS			string		`gorm:"type:int(8);defaulr:0" json:"ds" comment:"是否已被打印方删除 0为未删除 1为已删除"`
	Status 		int			`gorm:"type:int(8);NOT NULL;default:0" json:"status" comment:"打印的的状态 -1为订单上传 0为未打印 1为打印完成"`
	PaperFormat string		`gorm:"type:varchar(64);default:''" json:"paperFormat" comment:"打印的纸张大小"`
	Num			int			`gorm:"type:int(8);default:0" json:"num" comment:"打印的数量"`
	Color 		int 		`gorm:"type:int(8);check:Color_Checker, color in (0, 1);default:0" json:"color" comment:"颜色 0为黑白 1为彩色"`
	Direction 	int			`gorm:"type:int(8);check:Direction_Checker, direction in (0, 1);default:0" json:"direction" comment:"方向 0为纵向翻页 1为横向翻页"`
	SingleSide 	int			`gorm:"type:int(8);check:SingleSide_Checker, singleside in (0, 1);default:0" json:"singleSide" comment:"单双面 0为单面 1为双面"`
	Remarks		string		`gorm:"type:text" json:"remarks" comment:"备注"`
	Code 		string		`gorm:"type:varchar(128);NOT NULL" json:"code" comment:"编码"`
}