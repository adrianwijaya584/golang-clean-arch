package domain

import (
	userdto "clean_arch/dto/userDto"
	tablenames "clean_arch/utils/table_names"
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type User struct {
	ID   string `gorm:"primaryKey;size:40"`
	Name string `json:"name" gorm:"column:name;size:40"`
	gorm.Model
}

func (user User) TableName() string {
	return tablenames.Users
}

type UserGetAllAPI struct {
	ID        string
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	GetAll() ([]UserGetAllAPI, error)
	GetById(id string) (User, error)
	Create(id string, body userdto.CreateUserDto) error
}

type UserUseCase interface {
	GetAll() ([]UserGetAllAPI, error)
	GetById(id string) (User, error)
	Create(id string, body userdto.CreateUserDto) error
}
