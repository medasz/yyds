package collect

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"log"
	"time"
)

type Process struct {
	Pid     string    `json:"pid"`
	Ppid    string    `json:"ppid"`
	Cmdline string    `json:"cmdline"`
	Name    string    `json:"name"`
	Time    time.Time `json:"time"`
}

func GetProcess() (resData []Process) {
	p, _ := process.Pids()
	for _, pid := range p {
		p, err := process.NewProcess(pid)
		if err != nil {
			log.Println(err)
			continue
		}
		tmp := Process{}
		tmp.Time=time.Now()
		tmp.Pid = fmt.Sprintf("%d", pid)
		ppid, err := p.Ppid()
		if err != nil {
			log.Println(err)
			ppid = -1
		}
		tmp.Ppid = fmt.Sprintf("%d", ppid)
		cmdline, err := p.Cmdline()
		if err != nil {
			log.Println(err)
			cmdline = ""
		}
		tmp.Cmdline = cmdline
		name, err := p.Name()
		if err != nil {
			log.Println(err)
			name = ""
		}
		tmp.Name = name
		resData = append(resData, tmp)
	}
	return resData
}
