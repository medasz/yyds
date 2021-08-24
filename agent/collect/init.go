//获取进程列表，服务列表，端口列表，登录日志
package collect

<<<<<<< HEAD
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
=======
var allInfo = make(map[string]interface{})

func GetAllInfo() map[string]interface{} {
	allInfo["loginlog"] = GetLoginLog()
	allInfo["listen"] = GetListening()
	allInfo["service"] = GetService()
>>>>>>> 28a89f9c51ff1890d2c2fdbabcc30bcb5244c2ac
	allInfo["process"] = GetProcess()
	return allInfo
}
