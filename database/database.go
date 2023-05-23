package database

import (
	"fmt"
	"os"

	userModel "github.com/kilianp07/MuscleApp/models/user"
	weightModel "github.com/kilianp07/MuscleApp/models/weight"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var models = []any{
	&userModel.Model{},
	&weightModel.Model{},
}

func ConnectDatabase() (db *gorm.DB, err error) {

	mysqlDB := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_host") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(mysqlDB), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected and migrated successfully!")

	return database, nil
}
