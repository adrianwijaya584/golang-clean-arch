package delivery_test

import (
	"clean_arch/domain"
	"clean_arch/domain/mocks"
	"clean_arch/src/user/delivery"
	statuscode "clean_arch/utils/statusCode"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	mockUCase *mocks.UserUseCase
	handler   delivery.UserHandler
)

func init() {
	mockUCase = new(mocks.UserUseCase)

	handler = delivery.UserHandler{
		UserUCase: mockUCase,
		StatusCodeUtils: statuscode.StatusCodeUtils{
			DeliveryMsg: "error in user delivery test :",
		},
	}
}

func TestGetAll(t *testing.T) {

	mockUser := domain.UserGetAllAPI{
		ID:   "1",
		Name: "Adrian",
	}

	mockListUser := make([]domain.UserGetAllAPI, 0)
	mockListUser = append(mockListUser, mockUser)

	mockUCase.On("GetAll", mock.Anything).Return(mockListUser, nil).Once()

	e := echo.New()

	req, err := http.NewRequestWithContext(context.TODO(), echo.GET, "/users", strings.NewReader(""))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert.NoError(t, err)

	err = handler.GetAll(ctx)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	mockUCase.AssertExpectations(t)
}

func TestGetById(t *testing.T) {
	mockUser := domain.User{
		ID:   "1",
		Name: "Adrian",
	}

	mockUCase.On("GetById", mock.Anything).Return(mockUser, nil).Once()

	e := echo.New()

	req, err := http.NewRequestWithContext(context.TODO(), echo.GET, "/users/"+mockUser.ID, strings.NewReader(""))
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	assert.NoError(t, err)

	err = handler.GetById(ctx)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	mockUCase.AssertExpectations(t)
}
