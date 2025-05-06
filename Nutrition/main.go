package main

import (
	"github.com/Saivinay1464/Go-Learning/Nutrition/handlers"
	"github.com/Saivinay1464/Go-Learning/Nutrition/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitSampleData()

	r := gin.Default()

	r.GET("/food/:name", handlers.GetFood)
	r.POST("/upload", handlers.UploadCSV)

	r.Run(":8080")
}
