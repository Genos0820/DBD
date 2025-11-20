package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type RequestPayload struct {
	Name   string `json:"name" binding:"required"`
	PAN    string `json:"pan" binding:"required,pan"`
	Mobile string `json:"mobile" binding:"required,mobile"`
	Email  string `json:"email" binding:"required,email"`
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var payload RequestPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "request payload completed succesfully!"})
}

func validatePAN(fl validator.FieldLevel) bool {
	log.Printf("Validating Pan")
	pan := fl.Field().String()
	if len(pan) != 10 {
		return false
	}
	for i := 0; i < 5; i++ {
		if pan[i] < 'A' || pan[i] > 'Z' {
			log.Printf("Validating Pan1")
			return false
		}
	}
	for i := 5; i < 9; i++ {
		if pan[i] < '0' || pan[i] > '9' {
			log.Printf("Validating Pan2")
			return false
		}
	}
	log.Printf("Validating Pan3")
	return pan[9] >= 'A' && pan[9] <= 'Z'
}

func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	if len(mobile) != 10 {
		return false
	}
	for _, ch := range mobile {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func LatencyLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		log.Printf("Request to %s took %v", ctx.Request.URL.Path, latency)
	}
}

func main() {
	router := gin.Default()

	router.Use(LatencyLoggerMiddleware())

	// Replace Gin's default validator with the custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pan", validatePAN)
		v.RegisterValidation("mobile", validateMobile)
	}
	handler := NewHandler()

	router.POST("/create-user", handler.CreateUser)

	router.Run(":8080")
}
