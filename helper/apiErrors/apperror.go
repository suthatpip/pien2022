package apiErrors

import "net/http"

const (
	ServiceUnavailable = "SERVICE_UNAVAILABLE"
	ApiUnauthorized    = "API_UNAUTHORIZED"
	ApiBadRequest      = "API_BAD_REQUEST"
)

var (
	apiErrors = []apiError{
		{
			Id:      ServiceUnavailable,
			Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			Code:    1250,
			Status:  http.StatusServiceUnavailable,
		},
		{
			Id:      ApiUnauthorized,
			Message: "Unauthorized",
			Code:    1300,
			Status:  http.StatusUnauthorized,
		},
		{
			Id:      ApiBadRequest,
			Message: "Bad Request",
			Code:    1301,
			Status:  http.StatusBadRequest,
		},
	}
)
