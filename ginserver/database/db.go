package database

import (
	"fmt"
	"ginserver/config"
	"ginserver/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Config("HOST"), config.Config("PORT"), config.Config("USER"), config.Config("PASS"), config.Config("DB"))
	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("tidak dapat terkoneksi ke database")
	}

	DB.AutoMigrate(&model.Todo{})
	fmt.Println("telah terkoneksi")
}
