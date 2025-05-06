package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Saivinay1464/Go-Learning/Nutrition/models"
	"github.com/Saivinay1464/Go-Learning/Nutrition/storage"
	"github.com/gin-gonic/gin"
)

func GetFood(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))

	val, err := storage.RedisClient.Get(storage.Ctx, name).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Food not found"})
		return
	}

	var food models.Food
	json.Unmarshal([]byte(val), &food)

	c.JSON(http.StatusOK, food)
}
