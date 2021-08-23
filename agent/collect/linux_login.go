package collect

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"hids/agent/comman"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type utmp struct {
	UtType int32     //Type of login
	UtPid  int32     //Process ID of login process
	UtLine [32]byte  //Devicename
	UtId   [4]byte   //Inittab ID
	UtUser [32]byte  //Username
	UtHost [256]byte //Hostname for remote login
	UtExit struct {
		ETermination int16 //Process termination status
		EExit        int16 // Process exit status
	} //The structure describing the status of a terminated process.  This type is used in `struct utmp' below.Exit status of a process marked as DEAD_PROCESS.
	UtSession int32 //Session ID, used for windowing
	UtTv      struct {
		TvSec  int32 //Seconds
		TvUsec int32 //Microseconds
	} //Time entry was made
	UtAddrV6 [4]int32 //Internet address of remote host
	Unused   [20]byte //Reserved for future use
}
type LoginLog struct {
	Status   string    `json:"status"`
	Remote   string    `json:"remote"`
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

func GetLast() (resData []LoginLog) {
	file, err := os.Open("/var/log/wtmp")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	for {
		wtmp := new(utmp)
		err = binary.Read(file, binary.LittleEndian, wtmp)
		if err != nil {
			log.Println(err.Error())
			break
		}
		//fmt.Println(wtmp.UtType)
		if wtmp.UtType == 7 && int64(wtmp.UtTv.TvSec) > comman.LastTime {
			m := LoginLog{}
			m.Status = "true"
			m.Remote = string(bytes.Trim(wtmp.UtHost[:], "\x00"))
			if m.Remote == "" {
				continue
			}
			m.Username = string(bytes.Trim(wtmp.UtUser[:], "\x00"))
			m.Time = time.Unix(int64(wtmp.UtTv.TvSec), 0) //.Format(time.RFC1123)
			resData = append(resData, m)
		}
	}
	//fmt.Println("成功记录数:", len(resData))
	return resData
}

func GetLastbCmd() (resData []LoginLog) {
	cmd := exec.Command("lastb", "-F")
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err.Error())
		return
	}
	reg, err := regexp.Compile("\\s+")
	if err != nil {
		log.Println(err.Error())
		return
	}
	for true {
		r := bufio.NewReader(outPipe)
		line, _, err := r.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Println(err.Error())
			continue
		}
		line = reg.ReplaceAllLiteral(line, []byte(" "))

		col := bytes.SplitN(line, []byte("-"), 2)
		if len(col) < 2 {
			continue
		}

		if len(col) == 2 {
			coll := bytes.TrimSpace(col[0])
			col3 := bytes.SplitN(coll, []byte(" "), 4)
			localTime, err := time.LoadLocation("Asia/Shanghai")
			tt, err := time.ParseInLocation("Mon Jan 2 15:04:05 2006", string(col3[3]), localTime)
			if err != nil {
				log.Println(err)
				continue
			}
			if tt.Unix() > comman.LastTime {
				m := LoginLog{}
				m.Status = "false"
				m.Remote = string(col3[2])
				if m.Remote == "" {
					continue
				}
				m.Username = string(col3[0])
				m.Time = tt
				resData = append(resData, m)
			}
		}
	}
	return resData
}

func GetLastb() (resData []LoginLog) {
	file, err := os.Open("/var/log/btmp")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	for {
		wtmp := new(utmp)
		err = binary.Read(file, binary.LittleEndian, wtmp)
		if err != nil {
			log.Println(err.Error())
			break
		}
		//fmt.Println(wtmp.UtType)
		if int64(wtmp.UtTv.TvSec) > comman.LastTime {
			m := LoginLog{}
			m.Status = "false"
			m.Remote = string(bytes.Trim(wtmp.UtHost[:], "\x00"))
			if m.Remote == "" {
				continue
			}
			m.Username = string(bytes.Trim(wtmp.UtUser[:], "\x00"))
			m.Time = time.Unix(int64(wtmp.UtTv.TvSec), 0) //.Format(time.RFC1123)
			resData = append(resData, m)
		}
	}
	//fmt.Println("成功记录数:", len(resData))
	return resData
}
func GetLoginLog() (resultData []LoginLog) {
	resultData = append(resultData, GetLast()...)
	//resultData = append(resultData, GetLastb()...)
	resultData = append(resultData, GetLastbCmd()...)
	return resultData
}
