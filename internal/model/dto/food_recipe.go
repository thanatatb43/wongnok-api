package dto

import "time"

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

// ตัวแปลงค่าที่จะส่งกลับไปยัง client เป็นรูปแบบ JSON
type FoodRecipeResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Ingredients  string    `json:"ingredients"`
	Instructions string    `json:"instructions"`
	ImageURL     *string   `json:"image_url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

/*
	{
		"Name": "ข้าวผัดกุ้ง",
		"Description": "ข้าวผัดกุ้งรสชาติกลมกล่อม",
		"Ingredients": "ข้าว, กุ้ง, ไข่, ต้นหอม",
		"Instructions": "ผัดข้าวกับไข่ ใส่กุ้งและต้นหอม",
		"ImageURL": "https://example.com/image1.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 1
	},
	{
		"Name": "ต้มยำกุ้ง",
		"Description": "ต้มยำกุ้งรสจัดจ้าน",
		"Ingredients": "กุ้ง, ข่า, ตะไคร้, ใบมะกรูด, พริก",
		"Instructions": "ต้มเครื่องต้มยำ ใส่กุ้ง ปรุงรส",
		"ImageURL": "https://example.com/image2.jpg",
		"DifficultyID": 2,
		"CookingDurationID": 2
	},
	{
		"Name": "แกงเขียวหวานไก่",
		"Description": "แกงเขียวหวานไก่เข้มข้น",
		"Ingredients": "ไก่, มะเขือ, กะทิ, พริกแกงเขียวหวาน",
		"Instructions": "ผัดพริกแกง ใส่ไก่ เติมกะทิและมะเขือ",
		"ImageURL": "https://example.com/image3.jpg",
		"DifficultyID": 2,
		"CookingDurationID": 3
	},
	{
		"Name": "ผัดไทย",
		"Description": "ผัดไทยเส้นเหนียวนุ่ม",
		"Ingredients": "เส้นจันทน์, กุ้ง, ถั่วงอก, ไข่",
		"Instructions": "ผัดเส้นกับไข่ ใส่กุ้งและถั่วงอก",
		"ImageURL": "https://example.com/image4.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 2
	},
	{
		"Name": "ข้าวมันไก่",
		"Description": "ข้าวมันไก่เนื้อนุ่ม",
		"Ingredients": "ไก่, ข้าว, น้ำจิ้ม, แตงกวา",
		"Instructions": "ต้มไก่ หุงข้าวมัน เสิร์ฟพร้อมน้ำจิ้ม",
		"ImageURL": "https://example.com/image5.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 2
	},
	{
		"Name": "ส้มตำไทย",
		"Description": "ส้มตำไทยรสแซ่บ",
		"Ingredients": "มะละกอ, ถั่วฝักยาว, มะเขือเทศ, พริก",
		"Instructions": "ตำส่วนผสมทั้งหมด ปรุงรส",
		"ImageURL": "https://example.com/image6.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 1
	},
	{
		"Name": "หมูทอดกระเทียม",
		"Description": "หมูทอดกระเทียมหอมกรอบ",
		"Ingredients": "หมู, กระเทียม, ซอสปรุงรส",
		"Instructions": "หมักหมู ทอดกับกระเทียม",
		"ImageURL": "https://example.com/image7.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 1
	},
	{
		"Name": "แกงส้มผักรวม",
		"Description": "แกงส้มผักรวมรสเปรี้ยว",
		"Ingredients": "ผักรวม, น้ำแกงส้ม, กุ้ง",
		"Instructions": "ต้มผักกับน้ำแกงส้ม ใส่กุ้ง",
		"ImageURL": "https://example.com/image8.jpg",
		"DifficultyID": 2,
		"CookingDurationID": 2
	},
	{
		"Name": "ไข่เจียวหมูสับ",
		"Description": "ไข่เจียวหมูสับฟูกรอบ",
		"Ingredients": "ไข่, หมูสับ, น้ำปลา",
		"Instructions": "ตีไข่กับหมูสับ ทอดในน้ำมันร้อน",
		"ImageURL": "https://example.com/image9.jpg",
		"DifficultyID": 1,
		"CookingDurationID": 1
	},
	{
		"Name": "ข้าวผัดปู",
		"Description": "ข้าวผัดปูเนื้อแน่น",
		"Ingredients": "ข้าว, ปู, ไข่, ต้นหอม",
		"Instructions": "ผัดข้าวกับไข่ ใส่เนื้อปูและต้นหอม",
		"ImageURL": "https://example.com/image10.jpg",
		"DifficultyID": 2,
		"CookingDurationID": 2
	}
*/
