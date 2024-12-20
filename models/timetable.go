package models

type Teacher struct {
	id         uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	last_name  *string `json:"last_name"`                          // Фамилия
	first_name *string `json:"first_name"`                         // Имя
	patronymic *string `json:"patronymic"`                         // Отчество
}

type Lesson struct {
	id        uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	name      *string `json:"name"`                               // Название занятия
	professor uint    `json:"professor"`                          // ID профессора
}

type Schedule struct {
	id           uint `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	lesson_id    uint `json:"lesson_number"`                      // id предмета
	group_id     uint `json:"group_number"`                       // id группы
	class_id     uint `json:"class_number"`                       // id предмета
	day_id       uint `json:"day_number"`                         // День недели (id)
	classroom_id uint `json:"classroom_number"`                   // id аудитории
}

type Group struct {
	id      uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	name    *string `json:"name"`                               // Название группы
	course  uint    `json:"course"`                             // Курс
	number  uint    `json:"number"`                             // Номер группы
	program *string `json:"program"`                            // Учебная программа (БО/МО/CO)
}

type Classroom struct {
	id     uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	number *string `json:"number"`                             // Номер аудитории
}

type Day struct {
	id        uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	name      *string `json:"name"`                               // Название дня
	type_week *string `json:"type_week"`                          // Тип недели (числитель/знаменатель)
}
