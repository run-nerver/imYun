package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	NickName 		string 	`gorm:"type:varchar(64);NOT NULL" json:"nickName" comment:"别名"`
	Introduction    string 	`gorm:"not null; type:varchar(512)" json:"introduction" comment:"简介"`
	Avatar 			string 	`gorm:"type:text" json:"avatar" comment:"头像链接"`
	City			string	`gorm:"type:varchar(32)" json:"city" comment:"城市"`
	Gender			string	`gorm:"type:int(8)" json:"gender" comment:"性别 0未知 1男 2女"`
	Province		string	`gorm:"type:varchar(32)" json:"province" comment:"省份"`
	Unionid			string	`gorm:"type:varchar(128)" json:"unionid" comment:"用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回"`
	Openid 			string	`gorm:"type:varchar(128)" json:"openid" comment:"用户唯一标识"`
	InviteCode		string	`gorm:"type:varchar(64);NOT NULL;UNIQUE'" json:"inviteCode" comment:"用户邀请码"`
	Orders			[]Order
}