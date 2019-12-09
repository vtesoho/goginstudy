package main

import (
	"net/http"

	yaagGin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	yaag.Init(&yaag.Config{On:true,DocTitle:"ginTest",DocPath:"apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})
	r.Use(yaagGin.Document())

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.GET("/plain", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"result": "Hello World!"})
	})
	r.GET("/complex", func(c *gin.Context) {
		value := c.Query("key")
		c.JSON(http.StatusOK, gin.H{"value": value})
	})
	r.Run()
}
