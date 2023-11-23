package db

import (
	"fmt"

	"github.com/shakezidin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Username, Password, Host, Port, DBname,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection to DB %s Failed, Error: %s", DBname, err)
	}

	if err := database.AutoMigrate(models.User{}); err != nil {
		fmt.Printf("Automigrate failed for %s failed", DBname)
	}

	return database
}
