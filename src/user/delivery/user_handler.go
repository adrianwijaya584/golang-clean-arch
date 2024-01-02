package delivery

import (
	"clean_arch/domain"
	userdto "clean_arch/dto/userDto"
	"net/http"

	statuscode "clean_arch/utils/statusCode"

	"github.com/labstack/echo/v4"
	"github.com/lucsky/cuid"
)

type UserHandler struct {
	statuscode.StatusCodeUtils
	UserUCase domain.UserUseCase
}

func NewUsersHandler(app *echo.Echo, userUCase domain.UserUseCase) {
	handler := &UserHandler{
		UserUCase: userUCase,
		StatusCodeUtils: statuscode.StatusCodeUtils{
			DeliveryMsg: "error in user delivery :",
		},
	}

	userRoute := app.Group("/users")
	{
		userRoute.GET("", handler.GetAll)
		userRoute.GET("/:id", handler.GetById)
		userRoute.POST("", handler.Create)
	}
}

func (uHandler *UserHandler) GetAll(c echo.Context) error {
	userData, err := uHandler.UserUCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Error in server",
		})
	}

	return c.JSON(http.StatusOK, userData)
}

func (uHandler *UserHandler) GetById(c echo.Context) error {
	id := c.Param("id")

	userData, err := uHandler.UserUCase.GetById(id)

	if err != nil {
		return c.JSON(uHandler.GetStatusCode(err), echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userData)
}

func (_UserHandler *UserHandler) Create(c echo.Context) error {
	var userBody userdto.CreateUserDto

	if err := c.Bind(&userBody); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := c.Validate(userBody); err != nil {
		return err
	}

	id := cuid.New()

	err := _UserHandler.UserUCase.Create(id, userBody)

	if err != nil {
		return c.JSON(echo.ErrInternalServerError.Code, echo.Map{
			"error": "internal Server Error",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user created.",
	})
}
