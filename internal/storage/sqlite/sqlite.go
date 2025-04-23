package sqlite

import (
	"github.com/gitatractivo/gotodocli/internal/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type SQLiteStorage struct {
	db *gorm.DB
}


func NewSQLiteStorage(dbPath string)(*SQLiteStorage, error){
	db,err :=gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Task{})
	
	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) CreateTask(task *models.Task) error {
	return s.db.Create(task).Error
}

func (s *SQLiteStorage) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Find(&tasks).Error
	return tasks, err
}

func (s *SQLiteStorage) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	 err := s.db.First(&task, id).Error
	return &task, err
}

func (s *SQLiteStorage) UpdateTask(task *models.Task) error {
	return s.db.Save(task).Error
}

func (s *SQLiteStorage) DeleteTask(id uint) error {
	return s.db.Delete(&models.Task{}, id).Error
}

// func (s *SQLiteStorage) Close() error {
// 	return 
// }


