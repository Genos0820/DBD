package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// ValidatePAN function to validate the pan input
func ValidatePAN(fl validator.FieldLevel) bool {
	log.Printf("Validating Pan")
	pan := fl.Field().String()
	if len(pan) != 10 {
		return false
	}
	for i := 0; i < 5; i++ {
		if pan[i] < 'A' || pan[i] > 'Z' {
			return false
		}
	}
	for i := 5; i < 9; i++ {
		if pan[i] < '0' || pan[i] > '9' {
			return false
		}
	}
	return pan[9] >= 'A' && pan[9] <= 'Z'
}

// ValidateMobile function validate the mobile input
func ValidateMobile(fl validator.FieldLevel) bool {
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
