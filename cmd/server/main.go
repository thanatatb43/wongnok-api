package main

import (
	"wongnok-api/internal/foodrecipe"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	foodrecipeHandler := foodrecipe.NewHandler()

	router.GET("/food-recipes", foodrecipeHandler.Get)
}
