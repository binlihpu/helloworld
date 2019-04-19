package main

import (
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/binlihpu/helloworld/gin_swagger/docs"

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api/v1
func main() {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/hello", HelloWorldHandler)
	r.Run(":8088")
}

// HelloWorldHandler  HelloWorld
// @Summary HelloWorld
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "Greet"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/hello [get]
func HelloWorldHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"errMsg": "success",
		"data":   "hello world",
	})
}
