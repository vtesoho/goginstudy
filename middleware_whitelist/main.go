package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func IPAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		ipList:=[]string{
			"127.0.0.1",
		}
		flag:=false
		clientIP:= c.ClientIP()
		for _, host:=range ipList{
			if clientIP == host{
				flag = true
				break
			}
		}
		if !flag{
			c.String(401,"%s not in iplist",clientIP)
			c.Abort()
		}
	}
}


func main() {
	r:=gin.Default()
	container := dig.New()
	container.Provide("")
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200,"hello")
	})
	r.Run()
}
