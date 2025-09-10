package model

import "gorm.io/gorm"

type Difficulty struct {
	// ชื่อตัวแปรต้องขึ้นต้นด้วยตัวใหญ่ เพื่อให้ package อื่นสามารถเข้าถึงได้ และ orm จะแปลงชื่อตัวแปรเป็นชื่อคอลัมน์ในฐานข้อมูลให้อัตโนมัติ
	gorm.Model
	Name string
}
