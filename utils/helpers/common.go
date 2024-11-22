package helpers

import (
	"crypto/rand"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/**
 * initialize
 */
func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("trimData", trimData)
	}
}

/**
 * trim string common
 */
func TrimString(s string) string {
	return strings.TrimSpace(s)
}

/**
 *  custom validator
 */
var trimData validator.Func = func(fl validator.FieldLevel) bool {
	data := strings.TrimSpace(fl.Field().Interface().(string))
	if data == "" {
		return false
	}
	return true
}

/**
 * random string generator
 */
func GenerateRandomString(length int) (string, error) {
	// Define the character set you want to use
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Calculate the number of characters in the charset
	charsetLength := len(charset)

	// Generate random indexes to select characters from the charset
	randomIndexes := make([]byte, length)
	_, err := rand.Read(randomIndexes)
	if err != nil {
		return "", err
	}

	// Create the random string using the characters from the charset
	randomString := make([]byte, length)
	for i, b := range randomIndexes {
		randomString[i] = charset[b%byte(charsetLength)]
	}

	return string(randomString), nil
}
