package repository

import (
	"gorm.io/gorm"

	"clean_arch/domain"
	tablenames "clean_arch/utils/table_names"
)

type categoryRepo struct {
	db *gorm.DB
}

func CategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepo{
		db: db.Model(domain.Category{}).Table(tablenames.Categories),
	}
}

func (util *categoryRepo) GetAll() ([]domain.GetCategory, error) {
	categories := []domain.GetCategory{}

	err := util.db.Select("id", "name").Scan(&categories).Error

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
