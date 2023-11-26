package repository

import (
	"errors"
	"go-gin/data/request"
	models "go-gin/models"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	Db *gorm.DB
}

func NewTaskRepositoryImpl(Db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{Db: Db}
}

func (t TaskRepositoryImpl) Save(task models.Task) {
	result := t.Db.Create(&task)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t TaskRepositoryImpl) Update(task models.Task) {
	var updateTag = request.UpdateTaskRequest{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
	}
	result := t.Db.Model(&task).Updates(updateTag)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t TaskRepositoryImpl) Delete(taskId int) {
	var task models.Task
	result := t.Db.Where("id = ?", taskId).Delete(&task)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (t TaskRepositoryImpl) FindById(taskId int) (models.Task, error) {
	var task models.Task
	result := t.Db.Find(&task, taskId)
	if result != nil {
		return task, nil
	} else {
		return task, errors.New("not found")
	}
}

func (t TaskRepositoryImpl) FindAll() []models.Task {
	var tasks []models.Task
	result := t.Db.Find(&tasks)
	if result.Error != nil {
		panic(result.Error)
	}
	return tasks
}
