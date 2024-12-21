package models

type Teacher struct {
	Id         int    `json:"id"`         // Primary key
	Last_name  string `json:"last_name"`  // Фамилия
	First_name string `json:"first_name"` // Имя
	Patronymic string `json:"patronymic"` // Отчество
}

type Lesson struct {
	Id        int    `json:"id"`        // Primary key
	Name      string `json:"name"`      // Название занятия
	Professor int    `json:"professor"` // ID профессора
}

type Schedule struct {
	Id           int `json:"id"`               // Primary key
	Lesson_id    int `json:"lesson_number"`    // id предмета
	Group_id     int `json:"group_number"`     // id группы
	Class_id     int `json:"class_number"`     // id предмета
	Day_id       int `json:"day_number"`       // День недели (id)
	Classroom_id int `json:"classroom_number"` // id аудитории
}

type Group struct {
	Id      int    `json:"id"`      // Primary key
	Name    string `json:"name"`    // Название группы
	Course  int    `json:"course"`  // Курс
	Number  int    `json:"number"`  // Номер группы
	Program string `json:"program"` // Учебная программа (БО/МО/CO)
}

type Classroom struct {
	Id     int    `json:"id"`     // Primary key
	Number string `json:"number"` // Номер аудитории
}

type Day struct {
	Id        int    `json:"id"`        // Primary key
	Name      string `json:"name"`      // Название дня
	Type_week string `json:"type_week"` // Тип недели (числитель/знаменатель)
}
