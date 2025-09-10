package foodrecipe

import (
	"wongnok-api/internal/model"
	"wongnok-api/internal/model/dto"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Service เป็นโครงสร้างที่ใช้ในการจัดการตรรกะทางธุรกิจที่เกี่ยวข้องกับ FoodRecipe
type Service struct {
	Repository Repository
}

// NewService สร้างอินสแตนซ์ใหม่ของ Service โดยรับพารามิเตอร์เป็นการเชื่อมต่อฐานข้อมูล
func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(request dto.FoodRecipeCreateRequest) (model.FoodRecipe, error) {

	var recipe model.FoodRecipe
	recipe = recipe.FromRequest(request) // แปลงค่าจาก request เป็น model.FoodRecipe

	if err := service.Repository.Create(&recipe); err != nil {
		return model.FoodRecipe{}, errors.Wrap(err, "failed to create food recipe")
	}

	return recipe, nil // ส่งกลับ recipe ที่ถูกสร้างขึ้นพร้อมกับค่า nil สำหรับ error
}
