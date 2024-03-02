package main

import (
	"os"
	"server/models"
	"server/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Инициализация нового роутера Gin
	r := gin.Default()

	// Загрузка переменных среды из файла .env
	err := godotenv.Load()
	if err != nil {
		// Паника в случае ошибки загрузки .env файла
		panic(err.Error())
	}

	// Конфигурация базы данных
	Config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Инициализация подключения к базе данных
	models.InitDB(Config)
	models.InitMinio()
	// Настройка маршрутов
	routes.Route(r)

	// Запуск сервера на порту 8080
	r.Run(":8081")
	fmt.Println("server listening on 8081")
}
