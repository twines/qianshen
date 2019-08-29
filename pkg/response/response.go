/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/8/29 9:17
*/
package response

func Success(data interface{}, msg ...string) map[string]interface{} {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	return map[string]interface{}{"code": 0, "status": "success", "msg": message, "data": data}
}
func Error(msg string, code ...int) map[string]interface{} {
	errorCode := 1
	if len(code) > 0 {
		errorCode = code[0]
	}
	return map[string]interface{}{"code": errorCode, "status": "error", "msg": msg}
}
