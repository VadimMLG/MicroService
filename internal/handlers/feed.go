package handlers

import (
	"net/http"
	"strconv"

	"github.com/VadimMLG/MicroService/internal/storage"
	"github.com/gin-gonic/gin"
)

func GetFeed(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Получение ленты из PostgreSQL
	posts, err := storage.GetUserFeed(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feed"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
