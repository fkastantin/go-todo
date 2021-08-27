package main

import (
	"github.com/fkastantin/go-todo/model"
	"github.com/fkastantin/go-todo/controller"
	"github.com/fkastantin/go-todo/infra"
	"github.com/fkastantin/go-todo/repository"
	"github.com/fkastantin/go-todo/route"
	"github.com/fkastantin/go-todo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Env struct{
	db *gorm.DB
}

func main(){
	// Load environment variables
	infra.LoadEnvVariables()
	// Initlize database connection
	db := infra.GetDBConnection()

	// Repositories
	taskRepository := repository.TaskRepository{DB : db}

	// Services
	taskService := service.TaskService{TaskRepository: &taskRepository}

	// Controller
	taskController := controller.TaskController{TaskService: &taskService}
	
	// Setup GIN
	router := gin.Default()

	// Routes
	taskRouter := route.TaskRouter{TaskController: &taskController, Gin: router}

	// Setup Routes
	taskRouter.Setup()


	// Auto migration
	db.AutoMigrate(&model.Task{})

	// Expose port
	router.Run(":8000")
}