package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

/* test
curl -X GET 'http://127.0.0.1:8080/test'

gin.Logger()  中间件用于打印请求的相关信息
gin.Recovery() 用于捕获程序panic，保护程序不至于崩溃


 */

func main() {
	f,_:=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r:=gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		name:=c.DefaultQuery("name","defaule_name")
		panic("test panic")
		c.String(200,"%s",name)
	})

	r.Run()

}
