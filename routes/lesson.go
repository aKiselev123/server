package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllLessons(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "Предмет"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()
	lessons := []models.Lesson{}
	for rows.Next() {
		t := models.Lesson{}
		err := rows.Scan(&t.Id, &t.Name, &t.Professor)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lessons = append(lessons, t)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, lessons)
}

func UpdateLesson(c *gin.Context) {
	var updatedLessons models.Lesson
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updatedLessons); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "Предмет" SET Именование_дисциплины = $1, Преподаватель = $2 WHERE id = %s`, id)
	_, err := db.Exec(query, updatedLessons.Name, updatedLessons.Professor)
	if err != nil {
		fmt.Println("Ошибка изменения предмета:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("Предмет изменен:", updatedLessons)
	c.JSON(http.StatusCreated, updatedLessons)
}

func CreateLesson(c *gin.Context) {
	// Структура для привязки данных
	var newLesson models.Lesson

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newLesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "Предмет" (Именование_дисциплины, Преподаватель) VALUES ($1, $2)`
	_, err := db.Exec(query, newLesson.Name, newLesson.Professor)
	if err != nil {
		fmt.Println("Ошибка добавления предмета:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления предмета"})
		return
	}

	// Успешный ответ
	fmt.Println("Предмет добавлен:", newLesson)
	c.JSON(http.StatusCreated, newLesson)
}
