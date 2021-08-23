package handle

import "github.com/gin-gonic/gin"

func Login(c *gin.Context)  {
	c.HTML(200,"login-1.html",nil)
}