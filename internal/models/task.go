package models

import "time"

type Task struct{
	ID uint `json:"id" gorm:"primary_key;unique;autoIncrement"`
	Title string `json:"title" gorm:"not null"`
	Completed bool `json:"completed" gorm:"default:false"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

