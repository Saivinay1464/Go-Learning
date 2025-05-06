package handlers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Saivinay1464/Go-Learning/Nutrition/models"
	"github.com/Saivinay1464/Go-Learning/Nutrition/storage"
	"github.com/gin-gonic/gin"
)

func UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open file"})
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)

	_, err = reader.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CSV header"})
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading CSV"})
			return
		}

		serving, _ := strconv.Atoi(record[1])
		calories, _ := strconv.Atoi(record[2])
		fat, _ := strconv.ParseFloat(record[3], 64)
		protein, _ := strconv.ParseFloat(record[4], 64)
		carbs, _ := strconv.ParseFloat(record[5], 64)
		fiber, _ := strconv.ParseFloat(record[6], 64)

		food := models.Food{
			Name:     record[0],
			Serving:  serving,
			Calories: calories,
			Carbs:    carbs,
			Protein:  protein,
			Fat:      fat,
			Fiber:    fiber,
		}

		data, _ := json.Marshal(food)
		key := strings.ToLower(record[0])

		storage.RedisClient.Set(storage.Ctx, key, data, 0)
	}

	c.JSON(http.StatusOK, gin.H{"message": "CSV uploaded and data stored successfully"})
}
