package repository

import model "go-gin/models"

type TaskRepository interface {
	Save(task model.Task)
	Update(task model.Task)
	Delete(taskId int)
	FindById(taskId int) (task model.Task, err error)
	FindAll() []model.Task
}
