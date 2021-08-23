package collect

import (
	"hids/agent/comman"
	"regexp"
	"strings"
	"time"
)

type Listen struct {
	Proto   string    `json:"proto"`
	Address string    `json:"address"`
	Name    string    `json:"name"`
	Pid     string    `json:"pid"`
	Time    time.Time `json:"time"`
}

func GetListening() (resData []Listen) {
	data := comman.Cmdexec("ss -ntlp")
	listenList := strings.Split(data, "\n")
	if len(listenList) <= 2 {
		return
	}
	regLine := regexp.MustCompile("\\s+")
	regCol := regexp.MustCompile(`users:\(\("(.*?)",pid=(.*?),fd=(.*?)\)\)`)
	for _, info := range listenList[1 : len(listenList)-1] {
		if strings.Contains(info, "127.0.0.1") {
			continue
		}
		info = regLine.ReplaceAllString(strings.TrimSpace(info), " ")
		infoList := strings.Split(info, " ")
		if len(infoList) < 6 {
			continue
		}
		m := Listen{}
		m.Time = time.Now()
		m.Proto = "TCP"
		if strings.Contains(infoList[3], "::") {
			m.Address = strings.Replace(infoList[3], "::", "0.0.0.0", 1)
		} else if strings.Contains(infoList[3], "*") {
			m.Address = strings.Replace(infoList[3], "*", "0.0.0.0", 1)
		}
		flag := false
		for _, v := range resData {
			if v.Address == m.Address {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		users := regCol.FindAllStringSubmatch(infoList[5], -1)
		if len(users) < 1 || len(users[0]) < 4 {
			continue
		}
		m.Name = users[0][1]
		m.Pid = users[0][2]
		resData = append(resData, m)
	}
	return resData
}
