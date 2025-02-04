package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResponse represents the validation error response
type ValidationResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Details []ValidationError `json:"details"`
	Code    int               `json:"code"`
}

// getErrorMsg returns a user-friendly error message for validation tags
func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", fe.Param())
	case "username":
		return "Username must be 3-30 characters and contain only letters, numbers, and underscores"
	}
	return "Invalid value"
}

// ValidateRequest validates the request body against a struct type
func ValidateRequest(structType interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new instance of the struct
		req := structType

		// Bind JSON to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ValidationResponse{
				Error:   "validation_error",
				Message: "Invalid request body",
				Details: []ValidationError{{
					Field:   "body",
					Message: "Invalid JSON format",
				}},
				Code: http.StatusBadRequest,
			})
			c.Abort()
			return
		}

		// Validate struct
		if err := validate.Struct(req); err != nil {
			var errors []ValidationError

			for _, err := range err.(validator.ValidationErrors) {
				field := strings.ToLower(err.Field())
				errors = append(errors, ValidationError{
					Field:   field,
					Message: getErrorMsg(err),
				})
			}

			c.JSON(http.StatusBadRequest, ValidationResponse{
				Error:   "validation_error",
				Message: "Validation failed",
				Details: errors,
				Code:    http.StatusBadRequest,
			})
			c.Abort()
			return
		}

		// Set validated struct in context
		c.Set("validatedRequest", req)
		c.Next()
	}
}
