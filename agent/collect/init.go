//获取进程列表，服务列表，端口列表，登录日志
package collect

var allInfo = make(map[string]interface{})

func GetAllInfo() map[string]interface{} {
	allInfo["loginlog"] = GetLoginLog()
	allInfo["listen"] = GetListening()
	allInfo["service"] = GetService()
	allInfo["process"] = GetProcess()
	return allInfo
}
