package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"hids/web/handle"
)

func main() {
	web := gin.Default()
	//加载html模版
	web.LoadHTMLGlob("/root/hids/web/templates/*")
	//加载静态文件
	web.Static("/static", "/root/hids/web/static")
	//加载基于cookie的存储引擎,设置用于加密的密钥
	cookieStore := cookie.NewStore([]byte("Ljkdjr^djd"))
	//设置session中间件,
	web.Use(sessions.Sessions("user_session", cookieStore))
	web.Use(AuthRequired())
	web.GET("/", handle.Index)
	web.GET("/login", handle.Login)
	web.POST("/login", handle.Login)
	err := web.Run("0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
