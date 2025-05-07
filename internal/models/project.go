package models

import "time"

type Project struct {
	ID uint `json:"id" gorm:"primary_key;unique;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Tasks []Task `json:"tasks" gorm:"foreignKey:ProjectID"`
	TasksCount int `json:"tasks_count" gorm:"-"`
	CompletedTasksCount int `json:"completed_tasks_count" gorm:"-"`
	Users []User `json:"users" gorm:"many2many:project_users;"`
	UserIDs []uint `json:"user_ids" gorm:"-"`
}


