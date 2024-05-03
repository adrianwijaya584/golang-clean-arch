package domain

import (
	"gorm.io/gorm"

	tablenames "clean_arch/utils/table_names"
)

type Category struct {
	ID   string `gorm:"primaryKey;size:30" `
	Name string `gorm:"size:35"`
	gorm.Model
}

type GetCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (category Category) TableName() string {
	return tablenames.Categories
}

type CategoryRepository interface {
	GetAll() ([]GetCategory, error)
	GetById(id string) (Category, error)
}

type CategoryUseCase interface {
	GetAll() ([]GetCategory, error)
	GetById(id string) (Category, error)
}
