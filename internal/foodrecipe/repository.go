package foodrecipe

import (
	"wongnok-api/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	// ส่วนของการเชื่อมต่อฐานข้อมูลและการดำเนินการกับข้อมูลจะถูกเพิ่มเข้ามาที่นี่ในอนาคต
	DB *gorm.DB
}

// NewRepository สร้างอินสแตนซ์ใหม่ของ Repository โดยรับพารามิเตอร์เป็นการเชื่อมต่อฐานข้อมูล
func NewRepository(db *gorm.DB) Repository {
	return Repository{
		DB: db,
	}
}

// รับ pointer ของ model.FoodRecipe เพื่อให้สามารถแก้ไขค่า ID และ timestamps ได้
func (repo Repository) Create(recipe *model.FoodRecipe) error {
	// ทำการบันทึกข้อมูล recipe ลงในฐานข้อมูล
	return repo.DB.Create(recipe).Error
}
