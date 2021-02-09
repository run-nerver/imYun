package models

import "gorm.io/gorm"

type Printer struct {
	gorm.Model

	NickName 		string 	`gorm:"type:varchar(64);NOT NULL" json:"nickName" comment:"别名"`
	Name 			string 	`gorm:"type:varchar(32);NOT NULL;UNIQUE" json:"name" comment:"用户名"`
	PassWord 		string 	`gorm:"type:varchar(64);NOT NULL;" json:"passWord" comment:"密码"`

	PrinterShopName	string	`gorm:"type:varchar(64)" json:"printerShopName" comment:"打印店名子"`
	Introduction    string 	`gorm:"not null; type:varchar(512)" json:"introduction" comment:"简介"`
	Avatar 			string 	`gorm:"type:text" json:"avatar" comment:"头像链接"`
	PrinterIcon		[]PrinterIcon
	Order			[]Order
}