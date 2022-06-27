package apiErrors

import "net/http"

const (
	ServerError     = "SERVER_ERROR"
	ApiUnauthorized = "API_UNAUTHORIZED"
	ApiBadRequest   = "API_BAD_REQUEST"
)

var (
	apiErrors = []apiError{
		{
			Id:      ServerError,
			Message: "The server encountered an unexpected condition that prevented it from fulfilling the request.",
			Code:    1250,
			Status:  http.StatusInternalServerError,
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
