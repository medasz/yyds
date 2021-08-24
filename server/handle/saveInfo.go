package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hids/server/es/es"
	"log"
	"net/http"
)

func SaveInfo(c *gin.Context) {
	tmp := make(map[string][]map[string]string)
	err := c.BindJSON(&tmp)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	//fmt.Println(tmp["process"])
	//fmt.Println(tmp["loginlog"])
	//fmt.Println(tmp["listen"])
	//fmt.Println(tmp["service"])
	if _, ok := tmp["process"]; ok {
		fmt.Println("开始保存process数据")
		es.SaveInfo("process", tmp["process"])
		fmt.Println("保存process数据结束")
	} else if _, ok = tmp["loginlog"]; ok {
		fmt.Println("开始保存loginlog数据")
		es.SaveInfo("loginlog", tmp["loginlog"])
		fmt.Println("保存loginlog数据结束")
	} else if _, ok = tmp["listen"]; ok {
		fmt.Println("开始保存listen数据")
		es.SaveInfo("listen", tmp["listen"])
		fmt.Println("保存listen数据结束")
	} else if _, ok = tmp["service"]; ok {
		fmt.Println("开始保存service数据")
		es.SaveInfo("service", tmp["service"])
		fmt.Println("保存service数据结束")
	}

	c.JSON(200, gin.H{"message": "success"})
}
