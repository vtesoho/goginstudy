package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.POST("/test", func(c *gin.Context) {
		bodyByts,err:=ioutil.ReadAll(c.Request.Body)
		if err != nil{
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		// 通过readally读取之后，必须要重新在写入才能在拿取到body里面的内容
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName:=c.PostForm("first_name")
		lastName:=c.DefaultPostForm("last_name","default_last_name")


		c.String(http.StatusOK,"%s %s %s",firstName,lastName,string(bodyByts))
	})
	r.Run()
}
