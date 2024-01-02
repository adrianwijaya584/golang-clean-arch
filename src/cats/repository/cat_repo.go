package repository

import (
	"clean_arch/domain"
	"clean_arch/utils/db"

	"gorm.io/gorm"
)

type catRepoUtils struct {
	db *gorm.DB
}

func NewCatRepo() domain.CatRepository {
	return &catRepoUtils{
		db: db.DbConn(),
	}
}

func (util *catRepoUtils) GetAll() ([]domain.Cat, error) {
	catData := []domain.Cat{}

	err := util.db.Model(domain.Cat{}).Scan(&catData).Error

	if err != nil {
		return catData, err
	}

	return catData, nil
}
