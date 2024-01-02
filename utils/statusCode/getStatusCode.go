package statuscode

import (
	"clean_arch/domain"
	"log"
	"net/http"
)

type StatusCodeUtils struct {
	DeliveryMsg string
}

func (ut StatusCodeUtils) GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Println(ut.DeliveryMsg, err.Error())

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
