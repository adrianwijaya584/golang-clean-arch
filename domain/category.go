package domain

import (
	tablenames "clean_arch/utils/table_names"

	"gorm.io/gorm"
)

type Category struct {
	ID   string `gorm:"primaryKey;size:30" `
	Name string `gorm:"size:35"`
	gorm.Model
}

func (category Category) TableName() string {
	return tablenames.Categories
}

type CategoryRepository interface {
	GetAll() ([]Category, error)
	GetById(id string) (Category, error)
}

type CategoryUseCase interface {
	GetAll() ([]Category, error)
	GetById(id string) (Category, error)
}
