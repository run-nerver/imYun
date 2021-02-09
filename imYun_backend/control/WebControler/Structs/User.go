package Structs

import "gorm.io/gorm"

// 表单
type Login struct {
	UserName string `json:"username" validate:"required"`
	PassWord string `json:"password" validate:"required"`
}

type PrinterParams struct {
	ImageAvatar 	string 	`json:"ImagerAvatar"`
	Avatar 			string 	`json:"avatar"`
	Introduction 	string 	`json:"introduce"`
}

// 数据库写入结构体
type PrinterInfo struct {
	gorm.Model
	NickName 		string 	`json:"nickName" comment:"别名"`
	Name 			string 	`json:"name" comment:"用户名"`
	PassWord 		string 	`json:"passWord" comment:"密码"`
	Introduction    string 	`json:"introduction" comment:"简介"`
	Avatar 			string 	`json:"avatar" comment:"头像链接"`
}