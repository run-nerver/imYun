package Structs

type Login struct {
	Jscode		string		`json:"jscode" comment:"前端小程序所发送来的code"`
}

type UpdateUserInfo struct {
	Avatar 	string `json:"avatar"`
	City 	string `json:"city"`
	Gender	string `json:"gender"`
	NickName 	string `json:"nickName"`
	Province 	string `json:"province"`
}

// -----------------
// 		返回内容
// -----------------
type Code2Session struct {
	Openid 		string	`json:"openid"`
	Session_key string	`json:"session_key"`
	Unionid 	string	`json:"unionid"`
	Errcode 	int 	`json:"errcode"`
	Errmsg 		string 	`json:"errmsg"`
}