package delivery

import (
	"clean_arch/domain"
	statuscode "clean_arch/utils/statusCode"
	"net/http"

	"github.com/labstack/echo/v4"
)

type catDelivery struct {
	catUseCase domain.CatUseCase
	log        statuscode.StatusCodeUtils
}

func NewCatDelivery(app *echo.Echo, catUseCase domain.CatUseCase) {
	handler := &catDelivery{
		catUseCase: catUseCase,
		log: statuscode.StatusCodeUtils{
			DeliveryMsg: "error cat delivery :",
		},
	}

	catRoutes := app.Group("/cats")
	{
		catRoutes.GET("", handler.GetAll)
	}
}

func (util *catDelivery) GetAll(c echo.Context) error {
	catData, err := util.catUseCase.GetAll()

	if err != nil {
		return c.JSON(util.log.GetStatusCode(err), err.Error())
	}

	return c.JSON(http.StatusOK, catData)
}
