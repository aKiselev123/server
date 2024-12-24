package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllSchedules(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "Расписание"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()
	schedules := []models.Schedule{}
	for rows.Next() {
		t := models.Schedule{}
		err := rows.Scan(&t.Id, &t.Lesson_id, &t.Group_id, &t.Class_id, &t.Day_id, &t.Classroom_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		schedules = append(schedules, t)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, schedules)
}

func UpdateSchedule(c *gin.Context) {
	var updateSchedules models.Schedule
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updateSchedules); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "Расписание" SET Номер_дня_недели = $1, Номер_аудитории = $2, Номер_группы = $3, Номер_предмета = $4, Порядковый_номер_занятия = $5 WHERE id = %s`, id)
	_, err := db.Exec(query, updateSchedules.Day_id, updateSchedules.Classroom_id, updateSchedules.Group_id, updateSchedules.Lesson_id, updateSchedules.Class_id)
	if err != nil {
		fmt.Println("Ошибка изменения расписания:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("Расписание изменено:", updateSchedules)
	c.JSON(http.StatusCreated, updateSchedules)
}

func CreateSchedule(c *gin.Context) {
	// Структура для привязки данных
	var newSchedule models.Schedule

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "Расписание" (Номер_дня_недели, Номер_аудитории, Номер_группы, Номер_предмета, Порядковый_номер_занятия) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, newSchedule.Lesson_id, newSchedule.Group_id, newSchedule.Class_id, newSchedule.Day_id, newSchedule.Classroom_id)
	if err != nil {
		fmt.Println("Ошибка добавления пары:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления пары"})
		return
	}

	// Успешный ответ
	fmt.Println("Пара добавлена:", newSchedule)
	c.JSON(http.StatusCreated, newSchedule)
}
