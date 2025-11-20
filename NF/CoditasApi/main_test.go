package main

import (
	"bytes"
	"coditas/api/middlewares"
	"coditas/api/services"
	"coditas/api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestMainIntegration(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router from main.go
	router := setupRouter()

	// Test cases
	tests := []struct {
		name           string
		payload        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid input",
			payload: `{
                "name": "John Doe",
                "pan": "ABCDE1234F",
                "mobile": "9876543210",
                "email": "john.doe@example.com"
            }`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"request payload completed successfully!"}`,
		},
		{
			name: "Invalid PAN format",
			payload: `{
                "name": "John Doe",
                "pan": "INVALIDPAN",
                "mobile": "9876543210",
                "email": "john.doe@example.com"
            }`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error_message":"Key: 'RequestPayload.PAN' Error:Field validation for 'PAN' failed on the 'pan' tag"}`,
		},
		{
			name: "Missing required fields",
			payload: `{
                "pan": "ABCDE1234F",
                "mobile": "9876543210"
            }`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error_message":"Key: 'RequestPayload.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'RequestPayload.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name:           "Empty request body",
			payload:        `{}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error_message":"Key: 'RequestPayload.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'RequestPayload.PAN' Error:Field validation for 'PAN' failed on the 'required' tag\nKey: 'RequestPayload.Mobile' Error:Field validation for 'Mobile' failed on the 'required' tag\nKey: 'RequestPayload.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			name: "Extra fields in the request payload",
			payload: `{
                "name": "John Doe",
                "pan": "ABCDE1234F",
                "mobile": "9876543210",
                "email": "john.doe@example.com",
                "extraField": "extraValue"
            }`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"request payload completed successfully!"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request
			req := httptest.NewRequest(http.MethodPost, "/create-user", bytes.NewBufferString(tt.payload))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Assert the response body
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

// setupRouter initializes the router as defined in main.go
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.LatencyLoggerMiddleware())

	// Replacing Gin's default validator with the custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pan", utils.ValidatePAN)
		v.RegisterValidation("mobile", utils.ValidateMobile)
	}
	handler := services.NewHandler()

	router.POST("/create-user", handler.CreateUser)

	return router
}
