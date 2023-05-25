package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kilianp07/MuscleApp/api"
	"github.com/kilianp07/MuscleApp/utils/env"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, trying to load from system environment variables")
	}

	if !env.ValidateEnv() {
		return
	}

	api := api.NewApi()
	api.StartApi()
}
