package model

import (
	"time"
)

type Task struct{
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:100" json:"name"`
	Description string `gorm:"size:2000" json:"desciption"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (task Task) TableName() string{
	return "task"
}

