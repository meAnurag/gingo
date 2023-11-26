package service

import (
	"go-gin/data/request"
	"go-gin/data/response"
)

type TaskService interface {
	Create(task request.CreateTaskRequest)
	Update(task request.UpdateTaskRequest)
	Delete(taskId int)
	FindById(tagsId int) response.TaskResponse
	FindAll() []response.TaskResponse
}

// type TagsService interface {
// 	Create(tags request.CreateTagsRequest)
// 	Update(tags request.UpdateTagsRequest)
// 	Delete(tagsId int)
// 	FindById(tagsId int) response.TagsResponse
// 	FindAll() []response.TagsResponse
