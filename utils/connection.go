package util

import (
	"fmt"
	model "iris-todos/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		viper.Get("DB_HOST"),
		viper.Get("DB_USER"),
		viper.Get("DB_PASS"),
		viper.Get("DB_NAME"),
		viper.Get("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Todo{}, &model.User{})
}
