package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllTeachers(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "Преподаватель"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
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

func CreateTeacher(c *gin.Context) {
	// Структура для привязки данных
	var newTeacher models.Teacher

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newTeacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Валидация данных
	if newTeacher.Last_name == "" || newTeacher.First_name == "" || newTeacher.Patronymic == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields Last_name and First_name are required"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "Преподаватель" (Фамилия, Имя, Отчество) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, newTeacher.Last_name, newTeacher.First_name, newTeacher.Patronymic)
	if err != nil {
		fmt.Println("Ошибка добавления преподавателя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления преподавателя"})
		return
	}

	// Успешный ответ
	fmt.Println("Преподаватель добавлен:", newTeacher)
	c.JSON(http.StatusCreated, newTeacher)
}

func UpdateTeacher(c *gin.Context) {
	var updatedTeacher models.Teacher
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updatedTeacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "Преподаватель" SET Фамилия = $1, Имя = $2, Отчество = $3 WHERE id = %s`, id)
	_, err := db.Exec(query, updatedTeacher.Last_name, updatedTeacher.First_name, updatedTeacher.Patronymic)
	if err != nil {
		fmt.Println("Ошибка изменения преподавателя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("Преподаватель изменен:", updatedTeacher)
	c.JSON(http.StatusCreated, updatedTeacher)
}
