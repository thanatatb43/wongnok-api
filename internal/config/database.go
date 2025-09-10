package config

// Database โครงสร้างสำหรับเก็บค่าคอนฟิกของฐานข้อมูล
type Database struct {
	URL string `env:"DATABASE_URL" envDefault:"postgres://postgres:pass2word@localhost:5432/wongnok?sslmode=disable"`
}
