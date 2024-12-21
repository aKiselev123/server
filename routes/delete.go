package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func DeleteById(c *gin.Context) {
	// Получение названия таблицы из параметров URL
	tableName := c.Param("tableName")

	// Получение ID из параметров URL
	id := c.Param("id")

	// Проверка наличия ID
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID"})
		return
	}

	// Проверка названия таблицы
	if tableName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table name"})
		return
	}

	// Проверка на валидность названия таблицы (опционально)
	validTables := map[string]bool{
		"Преподаватель": true,
		"Предмет":       true,
		"Расписание":    true,
		"Аудитория":     true,
		"Группа":        true,
		"День_недели":   true,
	}
	if !validTables[tableName] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное имя таблицы"})
		return
	}

	// Получение подключения к базе данных
	db := GetDB()

	// Формирование SQL-запроса
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`, tableName)

	// Выполнение запроса
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Ошибка удаления записи:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления записи из БД"})
		return
	}

	// Успешный ответ
	fmt.Printf("Запись с ID %s удалена из таблицы %s\n", id, tableName)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Запись с ID %s удалена из таблица %s", id, tableName)})
}
