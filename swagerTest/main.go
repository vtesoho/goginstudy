package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vteso/test/swagerTest/user"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "vteso/test/swagerTest/docs"

)


// @title 测试
// @version 0.0.1
// @description  测试
// @BasePath /
func main() {
	r := gin.Default()

	// 创建路由组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/record/:userId", record)
		v1.GET("/sayHello/:name", user.SayHello)
	}


	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @获取指定ID记录
// @Description get record by ID
// @Summary 测试aaaaaa
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /api/v1/record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}


