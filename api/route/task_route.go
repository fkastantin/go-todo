package route

import (
	"github.com/fkastantin/go-todo/controller"

    "github.com/gin-gonic/gin"
)

type TaskRouter struct{
	TaskController *controller.TaskController
	Gin *gin.Engine
}

func (taskRouter TaskRouter) Setup() {
	task := taskRouter.Gin.Group("/tasks") //Router group
    
	task.GET("/", taskRouter.TaskController.GetTasks)
	task.POST("/", taskRouter.TaskController.AddTask)
	task.GET("/:id", taskRouter.TaskController.GetTaskById)
	task.DELETE("/:id", taskRouter.TaskController.DeleteTask)
	task.PUT("/:id", taskRouter.TaskController.UpdateTask)
}
