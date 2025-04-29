package responses

import "net/http"

type MessageDetail struct {
	HTTPCode int
	Message  string
}

var messages = map[int]MessageDetail{
	Success:      {http.StatusOK, "Success"},
	EmailInvalid: {http.StatusBadRequest, "Email is invalid"},

	Unauthorized: {http.StatusUnauthorized, "Unauthorized"},
	Forbidden:    {http.StatusForbidden, "Access denied"},

	UserNotFound: {http.StatusNotFound, "User not found"},
	SystemError:  {http.StatusInternalServerError, "Internal server error"},
}
