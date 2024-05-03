package delivery

import (
	"clean_arch/domain"
	statuscode "clean_arch/utils/statusCode"
	"net/http"

	"github.com/labstack/echo/v4"
)

type categoryDelivery struct {
	categoryUseCase domain.CategoryUseCase
	statusCodeUtil  statuscode.StatusCodeUtils
}

func CategoryDelivery(app *echo.Echo, usecase domain.CategoryUseCase) {
	handler := &categoryDelivery{
		categoryUseCase: usecase,
		statusCodeUtil: statuscode.StatusCodeUtils{
			DeliveryMsg: "Error in category delivery",
		},
	}

	categoriesRoute := app.Group("/categories")
	{
		categoriesRoute.GET("", handler.GetAll)
		categoriesRoute.GET("/:id", handler.GetById)
	}
}

func (delivery *categoryDelivery) GetAll(c echo.Context) error {
	categories, err := delivery.categoryUseCase.GetAll()

	if err != nil {
		return c.JSON(delivery.statusCodeUtil.GetStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, categories)
}

func (delivery *categoryDelivery) GetById(c echo.Context) error {
	category, err := delivery.categoryUseCase.GetById(c.Param("id"))

	if err != nil {
		return c.JSON(delivery.statusCodeUtil.GetStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, category)
}
