package middleware

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// Custom validation rules
var (
	usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")
)

// RegisterCustomValidators adds custom validation rules to the validator
func RegisterCustomValidators(v *validator.Validate) {
	// Username validation
	v.RegisterValidation("username", validateUsername)
	
	// Password strength validation
	v.RegisterValidation("strongpassword", validateStrongPassword)
	
	// Content validation (no excessive whitespace, etc.)
	v.RegisterValidation("content", validateContent)
}

// validateUsername checks if username meets our criteria
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	
	// Check length (3-30 characters)
	if len(username) < 3 || len(username) > 30 {
		return false
	}
	
	// Check if matches allowed pattern
	return usernameRegex.MatchString(username)
}

// validateStrongPassword checks password strength
func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	return hasUpper && hasLower && hasNumber && hasSpecial
}

// validateContent checks message content
func validateContent(fl validator.FieldLevel) bool {
	content := fl.Field().String()
	
	// Check if content is just whitespace
	trimmed := strings.TrimSpace(content)
	if len(trimmed) == 0 {
		return false
	}
	
	// Check for excessive whitespace
	spaces := regexp.MustCompile(`\s{3,}`).FindString(content)
	if len(spaces) > 0 {
		return false
	}
	
	return true
} 