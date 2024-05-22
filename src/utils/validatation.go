package utils

import (
	"regexp"
	"unicode"
)

// create a function which validate a email string and return a boolean
func IsEmailValid(email string) bool {

	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Match the email string against the pattern
	match, _ := regexp.MatchString(pattern, email)

	// Return true if the email is valid, false otherwise
	return match
}

func IsPasswordComplex(password string, minLength uint8, hasUpperCase bool, hasLowerCase bool, hasNumber bool, hasSpecialChar bool) bool {
	// Check if the password meets the minimum length requirement
	if len(password) < int(minLength) {
		return false
	}

	// Check if the password has at least one uppercase letter
	if hasUpperCase && !containsUpperCase(password) {
		return false
	}

	// Check if the password has at least one lowercase letter
	if hasLowerCase && !containsLowerCase(password) {
		return false
	}

	// Check if the password has at least one number
	if hasNumber && !containsNumber(password) {
		return false
	}

	// Check if the password has at least one special character
	if hasSpecialChar && !containsSpecialChar(password) {
		return false
	}

	// All checks passed, the password is complex
	return true
}

// Helper function to check if the password contains at least one uppercase letter
func containsUpperCase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// Helper function to check if the password contains at least one lowercase letter
func containsLowerCase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

// Helper function to check if the password contains at least one number
func containsNumber(password string) bool {
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

// Helper function to check if the password contains at least one special character
func containsSpecialChar(password string) bool {
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
