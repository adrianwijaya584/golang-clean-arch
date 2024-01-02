package usecase

import (
	"clean_arch/domain"
)

type catUseCase struct {
	catRepo domain.CatRepository
}

func NewCatUseCase(catRepo domain.CatRepository) domain.CatUseCase {
	return &catUseCase{
		catRepo: catRepo,
	}
}

func (usecase *catUseCase) GetAll() ([]domain.Cat, error) {
	catData, err := usecase.catRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return catData, nil
}
