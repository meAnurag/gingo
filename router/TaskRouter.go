package router

import (
	"go-gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(taskController *controller.TaskController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	taskRouter := router.Group("/task")
	taskRouter.GET("", taskController.FindAll)
	taskRouter.GET("/:taskId", taskController.FindById)
	taskRouter.POST("", taskController.Create)
	taskRouter.PUT("/:taskId", taskController.Update)
	taskRouter.DELETE("/:taskId", taskController.Delete)

	return service
}
