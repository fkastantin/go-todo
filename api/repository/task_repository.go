package repository

import(
	"gorm.io/gorm"
	"github.com/fkastantin/go-todo/model"
)

type TaskRepository struct{
	DB *gorm.DB
}


func (repo TaskRepository) Save(task *model.Task) error {
	return repo.DB.Create(task).Error
}

func (repo TaskRepository) FindAll() (*[]model.Task, error) {
	var tasks []model.Task
	err := repo.DB.Model(&model.Task{}).Order("created_at desc").Find(&tasks).Error

	return &tasks, err
}

func (repo TaskRepository) Update(task *model.Task) error{
	return repo.DB.Updates(task).Error
}

func (repo TaskRepository) Find(task *model.Task) (*model.Task, error) {
	var foundTask model.Task
	err := repo.DB.Where(task).Take(&foundTask).Error

	return &foundTask, err
}

func (repo TaskRepository) Delete(task *model.Task) error{
	return repo.DB.Delete(task).Error
} 