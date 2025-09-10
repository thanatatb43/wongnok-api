package model

import "gorm.io/gorm"

// gorm.Model เป็นโครงสร้างพื้นฐานที่มีฟิลด์ ID, CreatedAt, UpdatedAt, DeletedAt อยู่แล้ว
type CookingDuration struct {
	gorm.Model
	Name string
}
