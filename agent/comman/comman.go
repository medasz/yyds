package comman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var LastTime int64

func Cmdexec(cmd string) (data string) {
	argv := strings.Split(cmd, " ")
	c := exec.Command(argv[0], argv[1:]...)
	output, err := c.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		return data
	}
	return string(output)

}

func PostData(url string, dataJson []byte) {
	jsonByteIo := bytes.NewBuffer(dataJson)
	//创建请求
	req, err := http.NewRequest("POST", url, jsonByteIo)
	if err != nil {
		log.Println(err)
		return
	}
	//添加请求头
	req.Header.Set("Content-Type", "application/json")
	//创建客户端
	client := http.Client{}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	fmt.Println("数据提交成功")
}
func GetData(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	dataJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	data := make(map[string]string)
	err = json.Unmarshal(dataJson, &data)
	if err != nil {
		log.Println(err)
		return ""
	}
	return data["time"]
}
