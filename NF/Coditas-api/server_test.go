package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Replace Gin's default validator with the custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pan", validatePAN)
		v.RegisterValidation("mobile", validateMobile)
	}

	// Initialize the handler and router
	handler := NewHandler()
	router := gin.Default()
	router.POST("/create-user", handler.CreateUser)

	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   string
	}{
		{
			name:       "Valid Input",
			payload:    `{"name":"John Doe","pan":"ABCDE1234F","mobile":"9876543210","email":"john.doe@example.com"}`,
			statusCode: http.StatusOK,
			response:   `{"message":"request payload completed succesfully!"}`,
		},
		{
			name:       "Missing Name",
			payload:    `{"pan":"ABCDE1234F","mobile":"9876543210","email":"john.doe@example.com"}`,
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Key: 'RequestPayload.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			name:       "Invalid PAN",
			payload:    `{"name":"John Doe","pan":"INVALIDPAN","mobile":"9876543210","email":"john.doe@example.com"}`,
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Key: 'RequestPayload.PAN' Error:Field validation for 'PAN' failed on the 'pan' tag"}`,
		},
		{
			name:       "Invalid Mobile",
			payload:    `{"name":"John Doe","pan":"ABCDE1234F","mobile":"12345","email":"john.doe@example.com"}`,
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Key: 'RequestPayload.Mobile' Error:Field validation for 'Mobile' failed on the 'mobile' tag"}`,
		},
		{
			name:       "Invalid Email",
			payload:    `{"name":"John Doe","pan":"ABCDE1234F","mobile":"9876543210","email":"invalid-email"}`,
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Key: 'RequestPayload.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/create-user", bytes.NewBufferString(tt.payload))
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			// Assert the status code
			assert.Equal(t, tt.statusCode, resp.Code)

			// Assert the response body
			assert.JSONEq(t, tt.response, resp.Body.String(), "Response body mismatch")
		})
	}
}
