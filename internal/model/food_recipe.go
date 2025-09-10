package model

import (
	"wongnok-api/internal/model/dto"

	"gorm.io/gorm"
)

type FoodRecipe struct {
	gorm.Model
	Name              string
	Description       string
	Ingredients       string
	Instructions      string
	ImageURL          *string // ใช้ pointer เพื่อให้สามารถเก็บค่า null ได้
	DifficultyID      uint
	CookingDurationID uint
	Difficulty        Difficulty      `gorm:"foreignKey:DifficultyID"`
	CookingDuration   CookingDuration `gorm:"foreignKey:CookingDurationID"`
}

// ตัวแปลงค่าจาก model.FoodRecipe ส่งค่าจาก client มาเป็น struct FoodRecipe
func (recipe FoodRecipe) FromRequest(request dto.FoodRecipeCreateRequest) FoodRecipe {
	// mapping ค่าจาก request ไปยัง struct FoodRecipe
	return FoodRecipe{
		Name:              request.Name,
		Description:       request.Description,
		Ingredients:       request.Ingredients,
		Instructions:      request.Instructions,
		ImageURL:          &request.ImageURL,
		DifficultyID:      request.DifficultyID,
		CookingDurationID: request.CookingDurationID,
	}
}

// ตัวแปลงค่าจาก model.FoodRecipe เป็นรูปแบบที่ต้องการส่งกลับไปยัง client
func (recipe FoodRecipe) ToResponse() dto.FoodRecipeResponse {
	return dto.FoodRecipeResponse{
		ID:           recipe.ID,
		Name:         recipe.Name,
		Description:  recipe.Description,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		ImageURL:     recipe.ImageURL,
		CreatedAt:    recipe.CreatedAt,
		UpdatedAt:    recipe.UpdatedAt,
	}
}
