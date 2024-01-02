package main

import (
	"clean_arch/src/cats/delivery"
	"clean_arch/src/cats/repository"
	"clean_arch/src/cats/usecase"
	userDelivery "clean_arch/src/user/delivery"
	userRepo "clean_arch/src/user/repository"
	userUseCase "clean_arch/src/user/usecase"
	customvalidator "clean_arch/utils/custom_validator"
	"clean_arch/utils/db"
	"clean_arch/utils/env"
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

	dbConn := db.DbConn()

	if os.Getenv(env.GO_ENV) == "dev" {
		dbConn = dbConn.Debug()
	}

	userRepo := userRepo.NewUserRepo(dbConn)
	userUCase := userUseCase.NewUserUseCase(userRepo)
	userDelivery.NewUsersHandler(app, userUCase)

	catRepo := repository.NewCatRepo()
	catUCase := usecase.NewCatUseCase(catRepo)
	delivery.NewCatDelivery(app, catUCase)

	app.Start(":" + os.Getenv(env.PORT))
}
