package main

import (
	"log"
	"wongnok-api/internal/config"
	"wongnok-api/internal/foodrecipe"

	"github.com/caarlos0/env/v11"
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
	db, err := gorm.Open(postgres.Open(conf.Database.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error when connect to database:", err)
	}

	// ตรวจสอบการเชื่อมต่อฐานข้อมูล ว่าสามารถปิดการเชื่อมต่อได้หรือไม่
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	foodrecipeHandler := foodrecipe.NewHandler(db)

	router := gin.Default()

	// ประกาศ group สำหรับ version 1
	groupV1 := router.Group("/api/v1")

	groupV1.GET("/food-recipes", foodrecipeHandler.Get)
	groupV1.GET("/food-recipes/:id", foodrecipeHandler.GetByID)
	groupV1.POST("/food-recipes", foodrecipeHandler.Create)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
