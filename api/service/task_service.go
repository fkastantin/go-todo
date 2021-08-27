package service

import(
	"github.com/fkastantin/go-todo/repository"
	"github.com/fkastantin/go-todo/model"
)

type TaskService struct{
	TaskRepository *repository.TaskRepository
}

func (taskService TaskService) Save(task *model.Task) error {
	return taskService.TaskRepository.Save(task)
}

func (taskService TaskService) FindAll() (*[]model.Task, error) {
	return taskService.TaskRepository.FindAll()
}

func (taskService TaskService) Update(task *model.Task) error {
	return taskService.TaskRepository.Update(task)
}

func (taskService TaskService) Delete(id uint) error {
	task := model.Task{ID: id}
	return taskService.TaskRepository.Delete(&task)
}

func (taskService TaskService) FindById(id uint) (*model.Task, error) {
	task := model.Task{ID: id}
	return taskService.TaskRepository.Find(&task)
}