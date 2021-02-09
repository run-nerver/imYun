package libs

func MakeRespon(code int, message string) interface{} {
	Response := make(map[string]interface{})
	Response["code"] = code
	Response["message"] = message
	return Response
}

func MakeResponData(code int, message string, data map[string]interface{}) interface{} {
	Response := make(map[string]interface{})
	Response["code"] = code
	Response["message"] = message
	Response["data"] = data
	//for index := range data {
	//	Response[index] = data[index]
	//}
	return Response
}