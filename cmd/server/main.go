package main

import (
	"wongnok-api/internal/foodrecipe"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	foodrecipeHandler := foodrecipe.NewHandler()

	// ประกาศ group สำหรับ version 1
	groupV1 := router.Group("/api/v1")

	groupV1.GET("/food-recipes", foodrecipeHandler.Get)
	groupV1.GET("/food-recipes/:id", foodrecipeHandler.GetByID)
	groupV1.POST("/food-recipes", foodrecipeHandler.Create)

	router.Run(":8080")
}
