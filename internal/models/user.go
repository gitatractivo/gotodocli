package models

type User struct {
	ID         uint      `json:"id" gorm:"primary_key;unique;autoIncrement"`
	Username   string    `json:"username" gorm:"not null;unique"`
	Email      string    `json:"email" gorm:"not null;unique"`
	Password   string    `json:"password" gorm:"not null"`
	Projects   []Project `json:"projects" gorm:"many2many:project_users;"`
	ProjectIDs []uint    `json:"project_ids" gorm:"-"`
}
