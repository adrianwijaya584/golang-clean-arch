package repository

import (
	"clean_arch/domain"
	tablenames "clean_arch/utils/table_names"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func CategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepo{
		db: db.Model(domain.Category{}).Table(tablenames.Categories),
	}
}

func (util *categoryRepo) GetAll() ([]domain.Category, error) {
	categories := []domain.Category{}

	err := util.db.Scan(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (util *categoryRepo) GetById(id string) (domain.Category, error) {
	var category domain.Category

	err := util.db.First(&category, id).Error

	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}
