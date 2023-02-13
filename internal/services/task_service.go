package services

import (
	"github.com/adailsonm/desafio-sword/internal/models"
	"github.com/adailsonm/desafio-sword/internal/repository"
	"gorm.io/gorm"
)

type TaskService struct {
	repository repository.UserRepository
}

func NewTaskService(repository repository.UserRepository) TaskService {
	return TaskService{
		repository: repository,
	}
}

func (s TaskService) WithTrx(trxHandle *gorm.DB) TaskService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s TaskService) GetOneTask(id uint) (task models.Task, err error) {
	return task, s.repository.Find(&task, id).Error
}

func (s TaskService) GetAllTask() (tasks []models.Task, err error) {
	return tasks, s.repository.Find(&tasks).Error
}

func (s TaskService) CreateTask(task models.Task) error {
	return s.repository.Create(&task).Error
}

func (s TaskService) UpdateTask(task models.Task) error {
	return s.repository.Save(&task).Error
}

func (s TaskService) DeleteTask(id uint) error {
	return s.repository.Delete(&models.Task{}, id).Error
}
