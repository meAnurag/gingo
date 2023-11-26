package controller

import (
	"go-gin/data/request"
	"go-gin/data/response"
	"go-gin/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskSerivce service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{taskSerivce: taskService}
}

func (controller *TaskController) Create(c *gin.Context) {
	createTaskRequest := request.CreateTaskRequest{}

	if err := c.ShouldBindJSON(&createTaskRequest); err != nil {
		res := response.Response{
			Status: "Error",
			Error:  strings.Split(err.Error(), "Error")[1],
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	controller.taskSerivce.Create(createTaskRequest)

	res := response.Response{
		Status: "Success",
		Data:   "Task created.",
	}

	c.JSON(http.StatusOK, res)

}

func (controller *TaskController) Update(c *gin.Context) {
	updateTaskRequest := request.UpdateTaskRequest{}
	err := c.ShouldBindJSON(&updateTaskRequest)
	if err != nil {
		panic(err)
	}

	taskId := c.Param("taskId")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}

	updateTaskRequest.Id = id

	controller.taskSerivce.Update(updateTaskRequest)

	webResponse := response.Response{
		Status: "ok",
		Data:   nil,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TaskController) Delete(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}
	controller.taskSerivce.Delete(id)

	webResponse := response.Response{
		Status: "ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TaskController) FindById(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		panic(err)
	}

	taskResponse := controller.taskSerivce.FindById(id)

	webResponse := response.Response{
		Status: "ok",
		Data:   taskResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TaskController) FindAll(c *gin.Context) {
	taskResponse := controller.taskSerivce.FindAll()

	res := response.Response{
		Status: "ok",
		Data:   taskResponse,
	}

	c.JSON(http.StatusOK, res)
}
