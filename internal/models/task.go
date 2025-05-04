package models

import "time"

// todos: add validation and make fields optional like due date, category, tags, project, user, subtask, reminder, completed by, assigned to, assigned by, attachments, priority, description,
type Task struct {
	ID              uint       `json:"id" gorm:"primary_key;unique;autoIncrement"`
	Title           string     `json:"title" gorm:"not null;validate:required,min=3,max=100" `
	Completed       bool       `json:"completed" gorm:"default:false"`
	Description     *string    `json:"description"`
	CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	Priority        *int       `json:"priority"`
	DueDate         *time.Time `json:"due_date"`
	Category        *string    `json:"category"`
	Tags            []Tag      `json:"tags" gorm:"many2many:task_tags;"`
	ProjectID       *uint      `json:"project_id"`
	Project         *Project   `json:"project" gorm:"foreignKey:ProjectID"`
	Subtasks        []Task     `json:"subtasks" gorm:"foreignKey:ParentID"`
	ParentID        *uint      `json:"parent_id"`
	Parent          *Task      `json:"parent" gorm:"foreignKey:ParentID"`
	Reminder        *time.Time `json:"reminder"`
	ReminderSet     bool       `json:"reminder_set" gorm:"default:false"`
	CompletedAt     *time.Time `json:"completed_at"`
	CompletedBy     *uint      `json:"completed_by"`
	CompletedByUser *User      `json:"completed_by_user" gorm:"foreignKey:CompletedBy"`
	AssignedTo      *uint      `json:"assigned_to"`
	AssignedToUser  *User      `json:"assigned_to_user" gorm:"foreignKey:AssignedTo"`
	AssignedBy      *uint      `json:"assigned_by"`
	AssignedByUser  *User      `json:"assigned_by_user" gorm:"foreignKey:AssignedBy"`
}

//Here tags have to be unique and case insensitive and i need a new table for tags and i need to add a new field for tags in the task table and maybe a set that will only contain unique tags

type TaskTag struct {
	TaskID uint `json:"task_id"`
	TagID  uint `json:"tag_id"`
}
