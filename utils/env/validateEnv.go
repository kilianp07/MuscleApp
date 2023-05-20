package env

import (
	"fmt"
	"os"
)

var requiredEnv = []string{
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_HOST",
	"DB_PORT",
	"SECRET_KEY",
	"API_PORT",
}

// Ensure all the required environment variables are set
func ValidateEnv() bool {

	for _, key := range requiredEnv {
		_, found := os.LookupEnv(key)
		if !found {
			fmt.Printf("Missing %s key", key)
			return false
		}
	}
	return true
}
