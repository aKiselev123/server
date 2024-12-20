package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB инициализирует подключение к PostgreSQL
func ConnectDB() {
	// Загружаем переменные окружения из .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}

	// Получаем параметры подключения из переменных окружения
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Подключаемся к PostgreSQL через GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	// Выводим успешное сообщение
	fmt.Println("Успешное подключение к PostgreSQL!")

	// Сохраняем подключение в глобальной переменной
	DB = db
}

// GetDB возвращает экземпляр подключения к базе данных
func GetDB() *gorm.DB {
	return DB
}
