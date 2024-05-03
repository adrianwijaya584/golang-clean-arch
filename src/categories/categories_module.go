package categories

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"clean_arch/src/categories/delivery"
	"clean_arch/src/categories/repository"
	"clean_arch/src/categories/usecase"
)

func CategoriesModuleInit(dbConn *gorm.DB, app *echo.Echo) {
	categoryRepo := repository.CategoryRepository(dbConn)
	categoryUseCase := usecase.CategoryUseCase(categoryRepo)
	delivery.CategoryDelivery(app, categoryUseCase)
}
