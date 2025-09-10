package main

import (
	"log"
	"wongnok-api/internal/config"
	"wongnok-api/internal/foodrecipe"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// load configuration
	var conf config.Config

	if err := env.Parse(&conf); err != nil {
		log.Fatal("Failed to parse env var: ", err)
	}

	// Database connection
	gorm.Open(postgres.Open(conf.Database.URL), &gorm.Config{
		// แจ้งให้ GORM แสดง log การทำงานในระดับ Info
		Logger: logger.Default.LogMode(logger.Info),
	})

	router := gin.Default()

	foodrecipeHandler := foodrecipe.NewHandler()

	// ประกาศ group สำหรับ version 1
	groupV1 := router.Group("/api/v1")

	groupV1.GET("/food-recipes", foodrecipeHandler.Get)
	groupV1.GET("/food-recipes/:id", foodrecipeHandler.GetByID)
	groupV1.POST("/food-recipes", foodrecipeHandler.Create)

	router.Run(":8080")
}
