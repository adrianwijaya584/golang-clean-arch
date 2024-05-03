package statuscode

import (
	"clean_arch/domain"
	"net/http"
)

type StatusCodeUtils struct {
	DeliveryMsg string
}

func (ut StatusCodeUtils) GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

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
