package controller

import (
	"github.com/fkastantin/go-todo/service"
	"github.com/fkastantin/go-todo/model"
	"github.com/fkastantin/go-todo/infra"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	TaskService *service.TaskService
}


func (taskController TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := taskController.TaskService.FindAll()

	if err != nil {
		infra.ErrorJSON(ctx, http.StatusBadRequest, "Bad Request")
	} else {
		infra.SuccessJSON(ctx, http.StatusOK, tasks)
	}
}

func (taskController TaskController) AddTask(ctx *gin.Context) {
	var task model.Task
	ctx.ShouldBindJSON(&task)

	err := taskController.TaskService.Save(&task)

	if err != nil {
		infra.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create task")
	} else {
		infra.SuccessJSON(ctx, http.StatusCreated, nil)
	}
}

func (taskController TaskController) GetTaskById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := 	strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		infra.ErrorJSON(ctx, http.StatusBadRequest, "invalid id")
	} else {
		task, err := taskController.TaskService.FindById(uint(id))

		if err != nil {
			infra.ErrorJSON(ctx, http.StatusInternalServerError, "can not retrieve the task")
		} else {
			infra.SuccessJSON(ctx, http.StatusFound, task)
		}
	}
}

func (taskController TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := 	strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		infra.ErrorJSON(ctx, http.StatusBadRequest, "invalid id")
	} else {
		err := taskController.TaskService.Delete(uint(id))

		if err != nil {
			infra.ErrorJSON(ctx, http.StatusInternalServerError, "can not delete")
		} else {
			infra.SuccessJSON(ctx, http.StatusFound, nil)
		}
	}
}


func (taskController TaskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := 	strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		infra.ErrorJSON(ctx, http.StatusBadRequest, "invalid id")
	} else {
		var task model.Task
		ctx.ShouldBindJSON(&task)
		task.ID = uint(id)
		
		err := taskController.TaskService.Update(&task)

		if err != nil {
			infra.ErrorJSON(ctx, http.StatusInternalServerError, "can not update")
		} else {
			infra.SuccessJSON(ctx, http.StatusFound, nil)
		}
	}
}



