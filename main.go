package main

import (
	"os"

	"server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	routes.ConnectDB()
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	//Удаление из таблицы по Id
	router.DELETE("/delete/:tableName/:id", routes.DeleteById)

	// Преподаватели
	router.GET("/teachers", routes.GetAllTeachers)
	router.POST("/teachers", routes.CreateTeachers)
	router.PUT("/teachers/:id", routes.UpdateTeachers)
	// router.DELETE("/teachers/:id", routes.DeleteTeachers)

	// // Группы
	// router.GET("/groups", routes.GetAllGroups)
	// router.POST("/groups", routes.CreateGroups)
	// router.PUT("/groups/:id", routes.UpdateGroups)
	// router.DELETE("/groups/:id", routes.DeleteGroups)

	// // Аудитории
	// router.GET("/classrooms", routes.GetAllClassrooms)
	// router.POST("/classrooms", routes.CreateClassrooms)
	// router.PUT("/classrooms/:id", routes.UpdateClassrooms)
	// router.DELETE("/classrooms/:id", routes.DeleteClassrooms)

	// // Дни недели
	// router.GET("/days", routes.GetAllDays)
	// router.POST("/days", routes.CreateDays)
	// router.PUT("/days/:id", routes.UpdateDays)
	// router.DELETE("/days/:id", routes.DeleteDays)

	// // Пары
	// router.GET("/schedules", routes.GetAllSchedules)
	// router.POST("/schedules", routes.CreateSchedules)
	// router.PUT("/schedules/:id", routes.UpdateSchedules)
	// router.DELETE("/schedules/:id", routes.DeleteSchedules)

	// // Предметы
	// router.GET("/lessons", routes.GetAllLessons)
	// router.POST("/lessons", routes.CreateLessons)
	// router.PUT("/lessons/:id", routes.UpdateLessons)
	// router.DELETE("/lessons/:id", routes.DeleteLessons)

	router.Run(":" + port)
}
