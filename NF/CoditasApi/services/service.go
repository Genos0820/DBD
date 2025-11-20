package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Name   string `json:"name" binding:"required"`
	PAN    string `json:"pan" binding:"required,pan"`
	Mobile string `json:"mobile" binding:"required,mobile"`
	Email  string `json:"email" binding:"required"`
}

type Handler struct {
}

// NewHandler create new handler type and return it's address
func NewHandler() *Handler {
	return &Handler{}
}

// CreateUser handling the request and binding the data in the payload
func (h *Handler) CreateUser(ctx *gin.Context) {
	var payload RequestPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		// For other errors, return the default error message
		ctx.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "request payload completed successfully!"})
}
