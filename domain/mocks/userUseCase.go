package mocks

import (
	"clean_arch/domain"
	userdto "clean_arch/dto/userDto"

	"github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

func (_m *UserUseCase) GetAll() ([]domain.UserGetAllAPI, error) {
	ret := _m.Called()

	return ret.Get(0).([]domain.UserGetAllAPI), ret.Error(1)
}

func (_m *UserUseCase) GetById(id string) (domain.User, error) {
	ret := _m.Called()

	var r0 domain.User
	var r1 error

	if rf, ok := ret.Get(0).(func() domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.User)
		} else {
			r1 = ret.Error(0)
		}
	}

	return r0, r1
}

func (_m *UserUseCase) Create(id string, body userdto.CreateUserDto) error {
	ret := _m.Called()

	var r0 error

	if rf, ok := ret.Get(0).(func(userdto.CreateUserDto) error); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
