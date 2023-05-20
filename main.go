package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kilianp07/MuscleApp/api"
	"github.com/kilianp07/MuscleApp/utils/env"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !env.ValidateEnv() {
		return
	}

	api := api.NewApi()
	api.StartApi()
}
