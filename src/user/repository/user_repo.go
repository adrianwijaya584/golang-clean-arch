package repository

import (
	"clean_arch/domain"
	userdto "clean_arch/dto/userDto"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) GetAll() ([]domain.UserGetAllAPI, error) {
	users := []domain.User{}

	err := repo.db.Model(domain.User{}).Order("created_at DESC").Scan(&users).Error

	userData := []domain.UserGetAllAPI{}

	copier.Copy(&userData, &users)

	return userData, err
}

func (repo *userRepo) GetById(id string) (domain.User, error) {
	userData := domain.User{}

	err := repo.db.Model(domain.User{}).Where("id=?", id).Find(&userData).Error

	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (_userRepo *userRepo) Create(id string, body userdto.CreateUserDto) error {
	err := _userRepo.db.Debug().Create(&domain.User{
		ID:   id,
		Name: body.Name,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
