package Structs

import "time"

//接收
type GetOrderForm struct {
	Key 	string 	`json:"key" comment:"关键字"`
	Limit 	int 	`json:"limit"`
	Offset 	int 	`json:"offset"`
	DateFrom 		string 	`json:"dateFrom" comment:"订单起始时间"`
	DateEnd 		string 	`json:"dateEnd" comment:"订单结束时间"`
}


//转出
type OrderData struct {
	OrderId		string 	`json:"orderId"`
	Num 		int 	`json:"num"`
	Color		int 	`json:"color"`
	Direction	int 	`json:"direction"`
	SingleSide	int 	`json:"single_side"`
	Remarks		string 	`json:"remarks"`
}

type UOrderData struct {
	OrderId		string 	`json:"orderId"`
	FileName 	string  `json:"fileName"`

}

type OrderList struct {
	Id         	int    	`json:"id"`
	Status 		int 	`json:"status"`
	NickName 	string `json:"nickName"`
	UpdateTime	time.Time 	`json:"creatTime"`
	DeleteTime	time.Time 	`json:"deleteTime"`
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