package usecase_test

import (
	"clean_arch/domain"
	"clean_arch/domain/mocks"
	userdto "clean_arch/dto/userDto"
	"clean_arch/src/user/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockUserRepository(ctrl)

	// mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.UserGetAllAPI{
		ID:   "1",
		Name: "Adrian",
	}

	mockListUser := make([]domain.UserGetAllAPI, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		m.
			EXPECT().
			GetAll().Return(mockListUser, nil).AnyTimes()
		// mockUserRepo.On("GetAll").Return(mockListUser, nil).Once()

		u := usecase.NewUserUseCase(m)

		list, err := u.GetAll()

		assert.Len(t, list, len(mockListUser))
		assert.NoError(t, err)

		// mockUserRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockUserRepository(ctrl)
	// userRepoMock := new(mocks.UserRepository)
	userData := domain.User{
		ID:   "1",
		Name: "Adrian",
	}

	t.Run("success", func(t *testing.T) {
		// userRepoMock.On("GetById", mock.Anything).Return(userData, nil).Once()
		m.EXPECT().GetById(userData.ID).Return(userData, nil).AnyTimes()

		u := usecase.NewUserUseCase(m)

		data, err := u.GetById(userData.ID)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		// userRepoMock.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		// userRepoMock.On("GetById", mock.Anything).Return(domain.User{}, errors.New("")).Once()
		m.EXPECT().GetById(userData.ID).Return(domain.User{}, errors.New("")).AnyTimes()

		u := usecase.NewUserUseCase(m)

		data, err := u.GetById("1")

		assert.Error(t, err)
		assert.Equal(t, domain.User{}, data)

		// userRepoMock.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mocks.NewMockUserRepository(ctrl)

	// userRepoMock := new(mocks.UserRepository)
	userData := domain.User{
		ID:   "1",
		Name: "Adrian",
	}

	t.Run("success", func(t *testing.T) {
		tempUserData := userdto.CreateUserDto{
			Name: userData.Name,
		}
		m.EXPECT().Create(userData.ID, tempUserData).Return(nil).AnyTimes()
		// userRepoMock.On("Create", mock.Anything).Return(nil).Once()

		u := usecase.NewUserUseCase(m)
		err := u.Create(userData.ID, tempUserData)

		assert.NoError(t, err)
		assert.Equal(t, userData.Name, tempUserData.Name)
		// userRepoMock.AssertExpectations(t)
	})

}
