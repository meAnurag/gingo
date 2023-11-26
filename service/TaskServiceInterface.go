package service

import (
	"go-gin/data/request"
	"go-gin/data/response"
	model "go-gin/models"
	repository "go-gin/repository/Task"

	"github.com/go-playground/validator/v10"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	Validate       *validator.Validate
}

func NewTaskServiceImpl(taskRepository repository.TaskRepository, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		Validate:       validate,
	}
}

func (t TaskServiceImpl) Create(task request.CreateTaskRequest) {

	if err := t.Validate.Struct(task); err != nil {
		panic(err)
	}
	newTask := model.Task{Id: task.Id, Name: task.Name, Description: task.Description}

	t.TaskRepository.Save(newTask)
}

// func (t TaskServiceImpl) FindById()

func (t TaskServiceImpl) Update(task request.UpdateTaskRequest) {
	taskData, err := t.TaskRepository.FindById(task.Id)
	if err != nil {
		panic(err)
	}
	taskData.Name = task.Name
	taskData.Description = task.Description
	t.TaskRepository.Update(taskData)
}

func (t TaskServiceImpl) Delete(taskId int) {
	t.TaskRepository.Delete(taskId)
}

func (t TaskServiceImpl) FindById(taskId int) response.TaskResponse {
	taskData, err := t.TaskRepository.FindById(taskId)
	if err != nil {
		panic(err)
	}
	taskResponse := response.TaskResponse{
		Id:          taskData.Id,
		Name:        taskData.Name,
		Description: taskData.Description,
	}
	return taskResponse
}

func (t TaskServiceImpl) FindAll() []response.TaskResponse {
	result := t.TaskRepository.FindAll()

	var tags []response.TaskResponse

	for _, value := range result {
		tag := response.TaskResponse{
			Id:          value.Id,
			Name:        value.Name,
			Description: value.Description,
		}

		tags = append(tags, tag)
	}
	return tags
}
