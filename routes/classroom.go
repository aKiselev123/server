package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllClassrooms(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "Аудитория"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении аудиторий: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()
	classrooms := []models.Classroom{}
	for rows.Next() {
		c := models.Classroom{}
		err := rows.Scan(&c.Id, &c.Number)
		if err != nil {
			fmt.Println(err)
			continue
		}
		classrooms = append(classrooms, c)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, classrooms)
}

func UpdateClassroom(c *gin.Context) {
	var updatedClassroms models.Classroom
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updatedClassroms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "Аудитория" SET Номер_аудитории = $1 WHERE id = %s`, id)
	_, err := db.Exec(query, updatedClassroms.Number)
	if err != nil {
		fmt.Println("Ошибка изменения аудитории:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update classroom to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("Аудитория изменена:", updatedClassroms)
	c.JSON(http.StatusCreated, updatedClassroms)
}

func CreateClassroom(c *gin.Context) {
	// Структура для привязки данных
	var newClassroom models.Classroom

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newClassroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Валидация данных
	if newClassroom.Number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Проблема с полями"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "Аудитория" (Номер_аудитории) VALUES ($1)`
	_, err := db.Exec(query, newClassroom.Number)
	if err != nil {
		fmt.Println("Ошибка добавления аудитории:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления аудитории"})
		return
	}

	// Успешный ответ
	fmt.Println("Аудитория добавлена:", newClassroom)
	c.JSON(http.StatusCreated, newClassroom)
}
