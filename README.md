# wongnok-api

# Router
router link -> cmd/server/main.go -> internal -> (router_folder)/handler.go -> (router_folder)/service.go -> (router_folder)/repository.go

# Data
internal/data/recipe.go

create migrations folder

set .env ให้ goose เชื่อมต่อกับ database (postgres)

create new migrations
(สร้างตาราง difficulties)
goose create add_difficulties_table sql 
goose up

chceck goose
goose status

สร้าง folder internal/config
config/database.go
config/config.go

สร้าง Models
ใน folder data create data/model
data/model/difficulty.go

handler.go ใช้เรียก services ต่างในไฟล์ service.go
service.go เป็น logic ของ router นั้นๆ
repository.go จะเป็นตัวติดต่อ database

dto แปลงข้อมูลจากหน้าบ้านไปหลังบ้าน หรือกลับกัน
ใน folder model create folder dto

ความต่างระหว่าง model กับ dto
model จะเอาไว้ติดต่อกับฐานข้อมูลโดยตรง จึงมีตัวแปรที่ตรงกับ fields ใน database table
dto จะเอาไว้เป็นตัวกลางระหว่าง client กับ database จึงสามารถใส่ตัวแปรตาม logic ได้ ไม่จำเป็นต้องใส่ตาม database

errors wrapping 
go get github.com/pkg/errors