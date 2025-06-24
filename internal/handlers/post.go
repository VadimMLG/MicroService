package handlers

import (
	"net/http"
	"strconv"

	"github.com/VadimMLG/MicroService/internal/models"
	"github.com/VadimMLG/MicroService/internal/storage"
	"github.com/VadimMLG/MicroService/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Генерация уникальной ссылки
	link, err := utils.GenerateUniqueLink()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Link generation failed"})
		return
	}
	post.Link = link

	// Сохранение в PostgreSQL
	if err := storage.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
		return
	}

	// Кэширование в Redis
	if err := storage.CachePostLink(link, post.ID); err != nil {
		// Логируем ошибку, но не прерываем выполнение
		log.Println("Redis caching failed:", err)
	}

	c.JSON(http.StatusCreated, post)
}
