package collect

import (
	"hids/agent/comman"
	"regexp"
	"strings"
	"time"
)

type Service struct {
	Unit        string    `json:"unit"`
	Load        string    `json:"load"`
	Active      string    `json:"active"`
	Sub         string    `json:"sub"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

//获取新系统服务
func GetService() (resData []Service) {
	systemRaw := comman.Cmdexec("systemctl list-units --all")
	systemList := strings.Split(systemRaw, "\n")
	qiguai := "●"
	if len(systemList) < 10 {
		return
	}
	reg := regexp.MustCompile("\\s+")
	for _, systemLine := range systemList[1 : len(systemList)-8] {
		systemLine = strings.TrimSpace(systemLine)
		systemLineFmt := reg.ReplaceAllString(systemLine, " ")
		systemColList := strings.SplitN(systemLineFmt, " ", 5)
		if len(systemColList) < 5 {
			continue
		}
		tmp := Service{}
		tmp.Time = time.Now()
		if systemColList[0] != qiguai {
			tmp.Unit = systemColList[0]
			tmp.Load = systemColList[1]
			tmp.Active = systemColList[2]
			tmp.Sub = systemColList[3]
			tmp.Description = systemColList[4]
		} else {
			systemColListQiGuai := strings.SplitN(systemLineFmt, " ", 6)
			tmp.Unit = systemColListQiGuai[1]
			tmp.Load = systemColListQiGuai[2]
			tmp.Active = systemColListQiGuai[3]
			tmp.Sub = systemColListQiGuai[4]
			tmp.Description = systemColListQiGuai[5]
		}
		resData = append(resData, tmp)
	}
	return resData
}
