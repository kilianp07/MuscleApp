package database

import (
	"fmt"
	"os"

	exerciseModel "github.com/kilianp07/MuscleApp/models/Exercise"
	objectiveModel "github.com/kilianp07/MuscleApp/models/objective"
	userModel "github.com/kilianp07/MuscleApp/models/user"
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var models = []any{
	&userModel.User{},
	&weightModel.Weight{},
	&objectiveModel.Objective{},
	&exerciseModel.Exercise{},
}

func ConnectDatabase() (*gorm.DB, error) {

	mysqlDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(mysqlDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(models...); err != nil {
		return nil, err
	}

	fmt.Println("Database connected and migrated successfully!")

	return database, nil
}
