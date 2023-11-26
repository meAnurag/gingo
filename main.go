package main

import (
	"go-gin/config"
	"go-gin/controller"
	model "go-gin/models"
	repository "go-gin/repository/Task"
	"go-gin/router"
	"go-gin/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()

	validate := validator.New()

	db.Table("tasks").AutoMigrate(&model.Task{})

	taskrespository := repository.NewTaskRepositoryImpl(db)

	taskService := service.NewTaskServiceImpl(taskrespository, validate)

	taskController := controller.NewTaskController(taskService)

	routes := router.NewRouter(taskController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
