package config

import (
	"os"
	"pkg/coba/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function to create connection to our env file
func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local" 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Koneksi ke database gagal")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal menutup koneksi ke database")
	}
	dbSQL.Close()
}