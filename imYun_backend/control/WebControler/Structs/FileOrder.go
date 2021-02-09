package Structs

import "time"

type FileOrderParams struct {
	Key 	string 	`json:"key" comment:"关键字"`
	Limit 	int 	`json:"limit"`
	Offset 	int 	`json:"offset"`
	DateFrom 		string 	`json:"dateFrom" comment:"订单起始时间"`
	DateEnd 		string 	`json:"dateEnd" comment:"订单结束时间"`
}

type FileOrderData struct {
	UserId		int 		`json:"userId"`
	OrderId 	int 		`json:"orderId"`
	Status 		int 		`json:"status"`
	FileName 	string		`json:"fileName"`
	ReFileName 	string		`json:"reFileName"`
	UploadDate 	time.Time 	`json:"updateTime"`
	DeleteDate	time.Time 	`json:"deleteDate"`
}

type DeleteFOData struct {
	OrderIds 	[]string 	`json:"orderIds"`
}