package routes

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllTeachers(c *gin.Context) {
	// Получаем подключение к базе данных через GORM
	db := GetDB()

	// // Создаём срез для хранения
	// var teachers []models.Teacher

	// // Выполняем запрос к базе данных
	// if err := db.Find(&teachers).Error; err != nil {
	// 	log.Fatalf("Ошибка при извлечении пользователей: %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при извлечении пользователей"})
	// 	return
	// }
	rows, err := db.Query(`select * from "Преподаватель"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	teachers := []models.Teacher{}
	for rows.Next() {
		t := models.Teacher{}
		err := rows.Scan(&t.Id, &t.Last_name, &t.First_name, &t.Patronymic)
		if err != nil {
			fmt.Println(err)
			continue
		}
		teachers = append(teachers, t)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, teachers)
}

func CreateTeachers(c *gin.Context) {

}
