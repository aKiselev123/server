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
	router.GET("/teacher", routes.GetAllTeachers)
	router.POST("/teacher", routes.CreateTeacher)
	router.PUT("/teacher/:id", routes.UpdateTeacher)
	// router.DELETE("/teacher/:id", routes.DeleteTeachers)

	// Группы
	router.GET("/groups", routes.GetAllGroups)
	router.POST("/group", routes.CreateGroup)
	router.PUT("/group/:id", routes.UpdateGroup)
	// router.DELETE("/group/:id", routes.DeleteGroups)

	// Аудитории
	router.GET("/classroom", routes.GetAllClassrooms)
	router.POST("/classroom", routes.CreateClassroom)
	router.PUT("/classroom/:id", routes.UpdateClassroom)
	// router.DELETE("/classroom/:id", routes.DeleteClassrooms)

	// Дни недели
	router.GET("/day", routes.GetAllDays)
	router.POST("/day", routes.CreateDay)
	router.PUT("/day/:id", routes.UpdateDay)
	// router.DELETE("/day/:id", routes.DeleteDays)

	// Пары
	router.GET("/schedule", routes.GetAllSchedules)
	router.POST("/schedule", routes.CreateSchedule)
	router.PUT("/schedule/:id", routes.UpdateSchedule)
	// router.DELETE("/schedule/:id", routes.DeleteSchedules)

	// Предметы
	router.GET("/lesson", routes.GetAllLessons)
	router.POST("/lesson", routes.CreateLesson)
	router.PUT("/lesson/:id", routes.UpdateLesson)
	// router.DELETE("/lesson/:id", routes.DeleteLessons)

	router.Run(":" + port)
}
