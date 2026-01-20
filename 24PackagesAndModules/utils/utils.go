// Utility functions package
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Config represents package configuration
type Config struct {
	Debug   bool
	Version string
}

// Package-level configuration (private)
var config = Config{
	Debug:   false,
	Version: "1.0.0",
}

// init() function initializes the package
func init() {
	fmt.Println("Utils package initialized!")
	config.Debug = true
}

// Public functions

// Greet returns a greeting message
func Greet(name string) string {
	if name == "" {
		name = "Guest"
	}
	return fmt.Sprintf("Hello, %s! Welcome to Go packages.", name)
}

// ReverseString reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FormatNumber formats a number with commas
func FormatNumber(num float64) string {
	str := strconv.FormatFloat(num, 'f', 2, 64)
	parts := strings.Split(str, ".")

	// Add commas to the integer part
	intPart := parts[0]
	var result string
	for i := len(intPart) - 1; i >= 0; i-- {
		if (len(intPart)-i-1)%3 == 0 && i != len(intPart)-1 {
			result = "," + result
		}
		result = string(intPart[i]) + result
	}

	// Add decimal part if exists
	if len(parts) > 1 {
		result += "." + parts[1]
	}

	return result
}

// IsValidEmail checks if email format is valid
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// TruncateString truncates a string to specified length
func TruncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// GetConfig returns the current package configuration
func GetConfig() Config {
	return config
}

// SetDebug sets the debug mode
func SetDebug(debug bool) {
	config.Debug = debug
}

// Private helper functions

// cleanString removes extra whitespace (private)
func cleanString(s string) string {
	return strings.TrimSpace(s)
}

// validateInput performs basic validation (private)
func validateInput(input string) bool {
	return len(input) > 0 && len(input) <= 1000
}

// Public constants

const (
	MaxStringLength = 1000
	DefaultGreeting = "Hello"
)

// Private constants

const (
	maxRetries     = 3
	defaultTimeout = 30
)

// Public struct

// User represents a user with validation
type User struct {
	ID    int
	Name  string
	Email string
}

// NewUser creates a new user with validation
func NewUser(id int, name, email string) *User {
	if !IsValidEmail(email) {
		email = "unknown@example.com"
	}
	if name == "" {
		name = "Unknown User"
	}

	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// String returns a string representation of User
func (u *User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}

// Public interface

// Validator defines validation interface
type Validator interface {
	IsValid() bool
}

// EmailValidator implements Validator for email validation
type EmailValidator struct {
	Email string
}

// IsValid implements Validator interface
func (ev *EmailValidator) IsValid() bool {
	return IsValidEmail(ev.Email)
}

// NewEmailValidator creates a new email validator
func NewEmailValidator(email string) *EmailValidator {
	return &EmailValidator{Email: email}
}
