//获取进程列表，服务列表，端口列表，登录日志
package collect

import "fmt"

var allInfo = make(map[string]interface{})

func GetAllInfo() map[string]interface{} {
	fmt.Println("获取ssh登录历史")
	allInfo["loginlog"] = GetLoginLog()
	fmt.Println("获取监听端口")
	allInfo["listen"] = GetListening()
	fmt.Println("获取服务")
	allInfo["service"] = GetService()
	fmt.Println("获取进程")
	allInfo["process"] = GetProcess()
	return allInfo
}
