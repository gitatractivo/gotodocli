package models

type Tag struct {
	ID uint `json:"id" gorm:"primary_key;unique;autoIncrement"`
	Name string `json:"name" gorm:"not null;unique;validate:required,min=3,max=100"`
	Tasks []Task `json:"tasks" gorm:"many2many:task_tags;"`
}