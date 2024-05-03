package usecase

import (
	"clean_arch/domain"
)

type categoryUseCase struct {
	categoryRepository domain.CategoryRepository
}

func CategoryUseCase(categoryRepository domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepository: categoryRepository,
	}
}

// GetAll implements domain.CategoryUseCase.
func (usecase *categoryUseCase) GetAll() ([]domain.GetCategory, error) {
	categories, err := usecase.categoryRepository.GetAll()

	return categories, err
}

// GetById implements domain.CategoryUseCase.
func (usecase *categoryUseCase) GetById(id string) (domain.Category, error) {
	category, err := usecase.categoryRepository.GetById(id)

	if err != nil {
		return domain.Category{}, domain.ErrNotFound
	}

	return category, nil
}
