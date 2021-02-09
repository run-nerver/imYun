package Structs

import "time"

type OrderParams struct {
	Key 	string 	`json:"key" comment:"关键字"`
	Limit 	int 	`json:"limit"`
	Offset 	int 	`json:"offset"`
	DateFrom 		string 	`json:"dateFrom" comment:"订单起始时间"`
	DateEnd 		string 	`json:"dateEnd" comment:"订单结束时间"`
}

type LocalFile struct {
	ReFileName string `json:"reFileName"`
}

type OrderList struct {
	Id         	int    	`json:"id"`
	CreatTime	time.Time 	`json:"creatTime"`
	DeleteTime	time.Time 	`json:"deleteTime"`
	NickName   	string 	`json:"nickName"`
	FileName   	string 	`json:"fileName"`
	ReFileName 	string  `json:"rename"`
	PaperFormat string 	`json:"paperFormat"`
	FileColor  	int   	`json:"fileColor"`
	Direction  	int    	`json:"direction"`
	SingleSide 	int    	`json:"singleSide"`
	Remarks	   	string 	`json:"remarks"`
	Code		string 	`json:"code"`
	FileNum    	int    	`json:"fileNum"`
}

type OrderData struct {
	OrderId 	[]string 	`json:"orderId"`
}