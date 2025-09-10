package dto

// ตัวรับค่าจาก client เมื่อสร้างสูตรอาหารใหม่
type FoodRecipeCreateRequest struct {
	Name              string
	Description       string
	Ingredients       string
	Instructions      string
	ImageURL          string
	DifficultyID      uint
	CookingDurationID uint
}
