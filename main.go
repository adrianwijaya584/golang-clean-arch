package main

import (
	categoryDelivery "clean_arch/src/categories/delivery"
	categoryRepo "clean_arch/src/categories/repository"
	categoryUseCase "clean_arch/src/categories/usecase"
	catDelivery "clean_arch/src/cats/delivery"
	catRepository "clean_arch/src/cats/repository"
	catUseCase "clean_arch/src/cats/usecase"
	userDelivery "clean_arch/src/user/delivery"
	userRepo "clean_arch/src/user/repository"
	userUseCase "clean_arch/src/user/usecase"
	customvalidator "clean_arch/utils/custom_validator"
	"clean_arch/utils/db"
	"clean_arch/utils/env"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	app := echo.New()

	db.DbInit()

	app.Use(middleware.CORS())
	app.Use(middleware.Secure())
	app.Use(middleware.Gzip())
	app.Use(middleware.Recover())
	app.Validator = &customvalidator.CustomValidator{
		Validator: validator.New(
			validator.WithRequiredStructEnabled(),
		),
	}

	app.RouteNotFound("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": fmt.Sprintf("[%s] %s is not found", c.Request().Method, c.Request().RequestURI),
		})
	})

	dbConn := db.DbConn()

	userRepo := userRepo.NewUserRepo(dbConn)
	userUCase := userUseCase.NewUserUseCase(userRepo)
	userDelivery.NewUsersHandler(app, userUCase)

	catRepo := catRepository.NewCatRepo()
	catUCase := catUseCase.NewCatUseCase(catRepo)
	catDelivery.NewCatDelivery(app, catUCase)

	categoryRepo := categoryRepo.CategoryRepository(dbConn)
	categoryUseCase := categoryUseCase.CategoryUseCase(categoryRepo)
	categoryDelivery.CategoryDelivery(app, categoryUseCase)

	app.Start(":" + os.Getenv(env.PORT))
}
