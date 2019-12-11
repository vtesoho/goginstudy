package user

import "net/http"
import "github.com/gin-gonic/gin"

// @你好世界
// @Description 你好
// @Summary 测试bbbb
// @Accept  json
// @Produce json
// @Param   name     path    string     true        "name"
// @Success 200 {string} string	"name,helloWorld"
// @Router /api/v1/sayHello/{name} [get]
func SayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, name+",helloWorld")
}
