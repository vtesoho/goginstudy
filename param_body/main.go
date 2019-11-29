package main

import (
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

		c.String(http.StatusOK,string(bodyByts))
	})
	r.Run()
}
