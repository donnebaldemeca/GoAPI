package middleware

import (
	"errors"
	"net/http"

	"github.com/donnebaldemeca/GoAPI/api"
	"github.com/donnebaldemeca/GoAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// writer is used to write the response, i.e., http status code, headers, body
		// request is the incoming HTTP request containing header, body, etc.

		var username string = request.URL.Query().Get("username") // Get username from query parameters
		var token = request.Header.Get("Authorization")           // Get token from Authorization header
		var err error

		if username == "" || token == "" { // If username or token is missing
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(writer, UnAuthorizedError) // Return 400 Bad Request with error message
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(writer) // Return 500 Internal Server Error if DB connection fails
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username) // Fetch login details from DB

		if (loginDetails == nil) || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(writer, UnAuthorizedError) // Return 400 Bad Request if auth fails
			return
		}

		next.ServeHTTP(writer, request) // Call the next handler if authorization is successful
		// See api/api.go for next handler (GetCoinBalance)
	})
}
