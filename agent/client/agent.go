package client

import (
	"encoding/json"
	"fmt"
	"hids/agent/collect"
	"hids/agent/comman"
	"hids/agent/filter"
	"log"
	"os"
	"time"
)

type Agent struct {
	//定时器
	timer time.Duration
	//登录日志
	//进程列表
	//端口列表
	//服务列表
	//基础进程
	basePorcess []map[string]string
	//基础服务
	baseSystems []map[string]string
	//数据接收接口
	url string
}

func (a *Agent) Init() {
	a.url = os.Args[1]
	a.timer = 2
	filter.Init()
	go func() {
		url := "http://" + a.url + "/getInfo"
		for {
			timeString := comman.GetData(url)
			if timeString == "" {
				comman.LastTime = time.Now().Unix()
			} else if timeString == "all" {
				comman.LastTime = 869241600
			} else {
				lastTime, err := time.Parse("2006-01-02T15:04:05Z07:00", timeString)
				if err != nil {
					log.Println(err)
					continue
				}
				comman.LastTime = lastTime.Unix()
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (a *Agent) Run() {
	//定时上传登录日志，进程列表，端口列表，服务列表
	a.GetInfo()
}

func (a *Agent) GetInfo() {
	//搜集历史用户登录信息
	timer := time.NewTicker(a.timer * time.Minute)
	for _ = range timer.C {
		//fmt.Println(timerNow)
		//搜集历史登录信息，进程列表，端口列表，服务列表
		data := collect.GetAllInfo()
		//fmt.Println(data["loginlog"])
		//fmt.Println(data["listen"])
		//fmt.Println(data["service"])
		//fmt.Println(data["process"])
		//fmt.Println(len(data["service"]))
		data = a.FilterInfo(data)
		a.SendInfo(data)
		//fmt.Println(len(data["service"]))
	}
}

func (a *Agent) SendInfo(data map[string]interface{}) {
	url := "http://" + a.url + "/saveInfo"
	for k, v := range data {
		fmt.Printf("发送数据%s\n",k)
		tmp := make(map[string]interface{})
		tmp[k] = v
		dataJson, err := json.Marshal(tmp)
		if err != nil {
			log.Println(err)
			continue
		}
		comman.PostData(url, dataJson)
		fmt.Printf("发送数据%s结束\n",k)
	}

}

//过滤基础进程和服务
func (a *Agent) FilterInfo(data map[string]interface{}) (resData map[string]interface{}) {
	resData = make(map[string]interface{})
	resData["loginlog"] = data["loginlog"]
	resData["listen"] = data["listen"]
	resData["process"] = filter.FilterProcess(data["process"].([]collect.Process))
	resData["service"] = filter.FilterSystem(data["service"].([]collect.Service))
	return resData
}
