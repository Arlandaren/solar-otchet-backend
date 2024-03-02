package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Stats{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Template{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&ScheduleTemplate{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&GeneratedStats{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Histories{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}
