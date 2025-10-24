package database

import (
	"fmt"
	"log"
	"student_management/internal/config"
	"student_management/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	DB = db
	log.Println("connected to the database successfully!")
}
