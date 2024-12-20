package routes

import (
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllTeachers(c *gin.Context) {
	// Получаем подключение к базе данных через GORM
	db := GetDB()

	// Создаём срез для хранения пользователей
	var teachers []models.Teacher

	// Выполняем запрос к базе данных
	if err := db.Find(&teachers).Error; err != nil {
		log.Fatalf("Ошибка при извлечении пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при извлечении пользователей"})
		return
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, teachers)
}
