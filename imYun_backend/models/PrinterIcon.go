package models

import "gorm.io/gorm"

type PrinterIcon struct {
	gorm.Model

	PrinterID		int		`gorm:"type:int(6);NOT NULL" json:"userId" comment:"打印店老板的ID"`
	ImageAvater	string	`gorm:"type:text;NOT NULL" json:"imageAvater" comment:"打印店的相关图片"`
}
