package domain

import "gorm.io/gorm"

type Cat struct {
	ID   string `gorm:"primaryKey;size:40"`
	Name string `gorm:"size:30"`
	Race string `gorm:"size:30"`
	gorm.Model
}

type CatRepository interface {
	GetAll() ([]Cat, error)
}

type CatUseCase interface {
	GetAll() ([]Cat, error)
}
