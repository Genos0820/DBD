package main

import (
	"coditas/api/middlewares"
	"coditas/api/services"
	"coditas/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.LatencyLoggerMiddleware())

	// Replacing Gin's default validator with the custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pan", utils.ValidatePAN)
		v.RegisterValidation("mobile", utils.ValidateMobile)
	}
	handler := services.NewHandler()

	router.POST("/create-user", handler.CreateUser)

	router.Run(":8080")
}
