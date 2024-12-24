package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllDays(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "День_недели"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()
	days := []models.Day{}
	for rows.Next() {
		t := models.Day{}
		err := rows.Scan(&t.Id, &t.Name, &t.Type_week)
		if err != nil {
			fmt.Println(err)
			continue
		}
		days = append(days, t)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, days)
}

func UpdateDay(c *gin.Context) {
	var updatedDays models.Day
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updatedDays); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "День_недели" SET Название_дня_недели = $1, Тип_недели = $2 WHERE id = %s`, id)
	_, err := db.Exec(query, updatedDays.Name, updatedDays.Type_week)
	if err != nil {
		fmt.Println("Ошибка изменения дня недели:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update day week to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("День недели изменен:", updatedDays)
	c.JSON(http.StatusCreated, updatedDays)
}

func CreateDay(c *gin.Context) {
	// Структура для привязки данных
	var newDay models.Day

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newDay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Валидация данных
	if newDay.Name == "" || newDay.Type_week == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Проблема с полями"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "День_недели" (Название_дня_недели, Тип_недели) VALUES ($1, $2)`
	_, err := db.Exec(query, newDay.Name, newDay.Type_week)
	if err != nil {
		fmt.Println("Ошибка добавления преподавателя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления дня недели"})
		return
	}

	// Успешный ответ
	fmt.Println("День добавлен:", newDay)
	c.JSON(http.StatusCreated, newDay)
}
