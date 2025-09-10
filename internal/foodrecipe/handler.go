package foodrecipe

import (
	"net/http"
	"strconv"
	"wongnok-api/internal/data"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// สร้าง struct สำหรับจัดการ food recipes
}

func NewHandler() Handler {
	return Handler{}
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
	var request gin.H

	// ดึงข้อมูลจาก body ที่ส่งมาเก็บในตัวแปร request (&request คือการส่ง address ของตัวแปร request ไป request ข้างในเปลี่ยน request ข้างนอกจะเปลี่ยนตาม)
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var lastID int64
	// หา id ล่าสุดใน data.Recipes (internal/data/recipe.go)
	for id := range data.Recipes {
		lastID = id
	}

	// สร้าง id ใหม่
	lastID++
	request["id"] = lastID
	// เพิ่ม recipe ใหม่เข้าไปใน data.Recipes (internal/data/recipe.go)
	data.Recipes[lastID] = request

	// ส่ง response กลับไปที่ client
	ctx.JSON(http.StatusCreated, request)
}
