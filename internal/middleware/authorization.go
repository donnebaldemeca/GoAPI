package middleware

import (
	"errors"
	"net/http"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {

	var username string.r
}
