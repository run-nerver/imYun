package Structs

type UploadForm struct {
	PrinterID	int 	`json:"printerid"`
	FileName	string 	`json:"fileName"`

	PaperFormat string 	`json:"paperFormat"`
	Num 		int 	`json:"num"`
	Color		int 	`json:"color"`
	Direction	int 	`json:"direction"`
	SingleSide	int 	`json:"singleSide"`
	Remarks		string 	`json:"remarks"`
}