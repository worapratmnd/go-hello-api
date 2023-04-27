package main

import (
	_ "hello/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/
//	@BasePath		/
func main() {
	router := gin.Default()

	router.GET("/", hello)
	router.POST("/greeting", greeting)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":3001")
}

// ShowAccount godoc
//
//	@Summary		Show an hello
//	@Description	get hello by name
//	@Tags			hello
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	GreetingResponse
//	@Router			/ [get]
func hello(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "Hello World"})
}

// ShowAccount godoc
//
//	@Summary		Show an Greeting
//	@Description	get greeting by name
//	@Tags			greeting
//	@Accept			json
//	@Produce		json
//	@Param			request	body		GreetingRequest	true	"Account ID"
//	@Success		200		{object}	GreetingResponse
//	@Router			/greeting [post]
func greeting(c *gin.Context) {
	var (
		greet    GreetingRequest
		response GreetingResponse
	)
	c.BindJSON(&greet)
	response.Message = "Hello " + greet.Name
	c.JSON(200, response)
}

type GreetingRequest struct {
	Name string `json:"name" example:"World"`
}

type GreetingResponse struct {
	Message string `json:"message"`
}
