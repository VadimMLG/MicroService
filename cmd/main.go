package main

import (
	"log"
	"github.com/VadimMLG/MicroService/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация подключений к БД и Redis
	// (реализуем позже в storage/postgres.go и storage/redis.go)
	
	r := gin.Default()
	
	// Регистрация маршрутов
	r.POST("/posts", handlers.CreatePost)
	r.GET("/feed/:userID", handlers.GetFeed)
	
	log.Println("Server started on :8080")
	r.Run(":8080")
}
