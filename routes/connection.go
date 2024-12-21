package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// var DB *gorm.DB
var DB *sql.DB

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
	// strconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// // Подключаемся к PostgreSQL через GORM
	// db, err := gorm.Open(postgres.Open(strcon), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	// }

	// // Выводим успешное сообщение
	// fmt.Println("Успешное подключение к PostgreSQL!")
	strconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err := sql.Open("pgx", strconn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверка соединения
	if err = db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
	// Сохраняем подключение в глобальной переменной
	DB = db
}

// GetDB возвращает экземпляр подключения к базе данных
//
//	func GetDB() *gorm.DB {
//		return DB
//	}
func GetDB() *sql.DB {
	return DB
}
