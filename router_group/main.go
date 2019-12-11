package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	v1 := r.Group("/api/v1")
	{
		v2 := v1.Group("/user")
		{
			v2.GET("/add", func(c *gin.Context) {
				c.String(http.StatusOK,"userAddOk")
			})
		}
		v1.GET("/sayHello/:name", func(c *gin.Context) {
			c.String(http.StatusOK,"sayHello")
		})
	}
}
