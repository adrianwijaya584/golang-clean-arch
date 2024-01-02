package usecase

import (
	"clean_arch/domain"
	userdto "clean_arch/dto/userDto"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (userU *userUsecase) GetAll() ([]domain.UserGetAllAPI, error) {
	userData, err := userU.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (userU *userUsecase) GetById(id string) (domain.User, error) {
	userData, err := userU.userRepo.GetById(id)

	if err != nil {
		return domain.User{}, err
	}

	if userData.Name == "" {
		return domain.User{}, domain.ErrNotFound
	}

	return userData, err
}

func (_userU *userUsecase) Create(id string, body userdto.CreateUserDto) error {
	err := _userU.userRepo.Create(id, body)

	if err != nil {
		return err
	}

	return nil
}
