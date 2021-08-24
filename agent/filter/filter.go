package filter

import (
	"bytes"
	"hids/agent/collect"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var baseProcess = []map[string]string{}
var baseSystems = []map[string]string{}

func Init() {
	loadBaseProcess()
	loadBaseSystems()
}
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
func loadBaseProcess() {
	file, err := os.Open("/root/hids/agent/filter/base_process.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	dataLines := bytes.Split(data, []byte("\n"))
	for _, dataLine := range dataLines {
		processData := make(map[string]string)
		tmp := bytes.Split(dataLine, []byte("!"))
		processData["pid"] = string(tmp[0])
		processData["ppid"] = string(tmp[1])
		processData["cmdline"] = string(tmp[2])
		processData["name"] = string(tmp[3])
		baseProcess = append(baseProcess, processData)
	}
}
func loadBaseSystems() {
	file, err := os.Open("/root/hids/agent/filter/base_systems.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	systemRaw, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	systemList := strings.Split(string(systemRaw), "\n")
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
		tmp := make(map[string]string)
		if systemColList[0] != qiguai {
			tmp["unit"] = systemColList[0]
			tmp["load"] = systemColList[1]
			tmp["active"] = systemColList[2]
			tmp["sub"] = systemColList[3]
			tmp["description"] = systemColList[4]
		} else {
			systemColListQiGuai := strings.SplitN(systemLineFmt, " ", 6)
			tmp["unit"] = systemColListQiGuai[1]
			tmp["load"] = systemColListQiGuai[2]
			tmp["active"] = systemColListQiGuai[3]
			tmp["sub"] = systemColListQiGuai[4]
			tmp["description"] = systemColListQiGuai[5]
		}
		baseSystems = append(baseSystems, tmp)
	}
}

func FilterProcess(data []collect.Process) (resData []collect.Process) {
	for _, processVal := range data[:] {
		if isSame(processVal.Name, baseProcess, "name") && isSame(processVal.Cmdline, baseProcess, "cmdline") {

		} else {
			resData = append(resData, processVal)
		}
	}
	return resData
}

func FilterSystem(data []collect.Service) (resData []collect.Service) {
	for _, systemVal := range data[:] {
		if isSame(systemVal.Unit, baseSystems, "unit") {

		} else {
			resData = append(resData, systemVal)
		}
	}
	return resData
}

func isSame(src string, baseData []map[string]string, key string) bool {
	for _, val := range baseData {
		if src == val[key] {
			return true
		}
	}
	return false
}
