package main

import (
	"github.com/alonelegion/musicstore_rest_api/controllers"
	"github.com/alonelegion/musicstore_rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	models.ConnectDB()

	// Маршруты
	route.GET("/tracks", controllers.GetAllTracks)

	// Запуск сервера
	route.Run()
}
