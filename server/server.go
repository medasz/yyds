package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hids/server/es/es"
	"hids/server/handle"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: server ip:port")
		fmt.Println("Example: server 0.0.0.0:8080")
		return
	}
	es.ConnectES()
	r := gin.Default()
	r.POST("/saveInfo", handle.SaveInfo)
	r.GET("/getInfo", handle.GetInfo)
	err := r.Run(os.Args[1])
	if err != nil {
		panic(err)
	}
}
