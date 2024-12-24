package routes

import (
	"fmt"
	"log"
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetAllGroups(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query(`select * from "Группа"`)
	if err != nil {
		log.Fatalf("Ошибка при извлечении пользователей: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
	}
	defer rows.Close()
	groups := []models.Group{}
	for rows.Next() {
		g := models.Group{}
		err := rows.Scan(&g.Id, &g.Name, &g.Course, &g.Number, &g.Program)
		if err != nil {
			fmt.Println(err)
			continue
		}
		groups = append(groups, g)
	}

	// Отправляем ответ в формате JSON
	c.JSON(http.StatusOK, groups)
}

func UpdateGroup(c *gin.Context) {
	var updatedGroup models.Group
	id := c.Param("id")

	// Привязка входящих данных JSON к структуре
	if err := c.ShouldBindJSON(&updatedGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDB()

	query := fmt.Sprintf(`UPDATE "Группа" SET Наименование = $1, Курс = $2, Номер_группы = $3, Программа_обучения = $4 WHERE id = %s`, id)
	_, err := db.Exec(query, updatedGroup.Name, updatedGroup.Course, updatedGroup.Number, updatedGroup.Program)
	if err != nil {
		fmt.Println("Ошибка изменения группы:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update group to the database"})
		return
	}

	// Успешный ответ
	fmt.Println("Группа изменена:", updatedGroup)
	c.JSON(http.StatusCreated, updatedGroup)
}

func CreateGroup(c *gin.Context) {
	// Структура для привязки данных
	var newGroup models.Group

	// Привязка JSON из запроса к структуре
	if err := c.BindJSON(&newGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// SQL-запрос на добавление данных
	query := `INSERT INTO "Группа" (Наименование, Курс, Номер_группы, Программа_обучения) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, newGroup.Name, newGroup.Course, newGroup.Number, newGroup.Program)
	if err != nil {
		fmt.Println("Ошибка добавления предмета:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления группы"})
		return
	}

	// Успешный ответ
	fmt.Println("Группа добавлена:", newGroup)
	c.JSON(http.StatusCreated, newGroup)
}
