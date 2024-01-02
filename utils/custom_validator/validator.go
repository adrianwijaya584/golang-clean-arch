package customvalidator

import (
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type ValidationErrorRes struct {
	Field   string
	Tag     string
	Param   string
	Message string
}

func (cv *CustomValidator) Validate(i any) error {
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(cv.Validator, trans)

	if err := cv.Validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)

		errRes := []ValidationErrorRes{}

		for _, e := range errs {
			errRes = append(errRes, ValidationErrorRes{
				Field:   e.Field(),
				Tag:     e.ActualTag(),
				Param:   e.Param(),
				Message: e.Translate(trans),
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
			"errors": errRes,
		})
	}

	return nil
}
