package db

import (
	"backend/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf(".envファイルの読み込みに失敗: %v\n", err)
		return
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	portStr := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSLMODE")
	TimeZone := os.Getenv("TIMEZONE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("データベースポート番号の取得に失敗: %v\n", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベース接続に失敗: %v\n", err)
		return
	}
	log.Println("データベース接続に成功")

	db.AutoMigrate(&models.Todo{})
	log.Println("マイグレーションに成功")

	DB = db

}
