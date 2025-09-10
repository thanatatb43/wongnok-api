package foodrecipe

import (
	"net/http"
	"strconv"
	"wongnok-api/internal/data"
	"wongnok-api/internal/model/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	// handler เรียกใช้ service ในการจัดการตรรกะทางธุรกิจ จากนั้น service จะเรียกใช้ repository ในการจัดการข้อมูล
	// สร้าง struct สำหรับจัดการ food recipes
	Service Service
}

// NewHandler สร้างอินสแตนซ์ใหม่ของ Handler โดยรับพารามิเตอร์เป็นการเชื่อมต่อฐานข้อมูล
func NewHandler(db *gorm.DB) Handler {
	return Handler{
		Service: NewService(db), // สร้าง service ใหม่โดยส่งการเชื่อมต่อฐานข้อมูลไปให้
	}
}

// /food-recipes
func (handler Handler) Get(ctx *gin.Context) {
	// ประกาศตัวแปรมารับ arrays
	var recipes []gin.H

	// ทำการ loop ดึงข้อมูลจาก data.Recipes (internal/data/recipe.go) มาเก็บในตัวแปร recipes
	for _, recipe := range data.Recipes {
		recipes = append(recipes, recipe)
	}

	// ส่งข้อมูลกลับไปที่ client
	ctx.JSON(http.StatusOK, recipes)
}

// /food-recipes/:id
func (handler Handler) GetByID(ctx *gin.Context) {
	// ดึง id จาก parameter
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)

	// ตรวจสอบ error กรณีที่ id ไม่ใช่ตัวเลข
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid path parameter (ID)"})
		return
	}

	// ดึงข้อมูล recipe จาก data.Recipes (internal/data/recipe.go) ตาม id ที่ได้รับมา
	result, ok := data.Recipes[id]

	// ตรวจสอบกรณีที่ไม่พบ recipe ที่มี id ดังกล่าว
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// POST /food-recipes
func (handler Handler) Create(ctx *gin.Context) {
	var request dto.FoodRecipeCreateRequest // สร้าง struct สำหรับรับข้อมูลจาก client

	// ดึงข้อมูลจาก body ที่ส่งมาเก็บในตัวแปร request (&request คือการส่ง address ของตัวแปร request ไป request ข้างในเปลี่ยน request ข้างนอกจะเปลี่ยนตาม)
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := handler.Service.Create(request) // เรียกใช้ service เพื่อสร้าง recipe ใหม่
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูล recipe ที่ถูกสร้างขึ้นกลับไปยัง client
	ctx.JSON(http.StatusCreated, recipe.ToResponse())
}
