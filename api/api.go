package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct { // API will take parameter Username for CoinBalance endpoint
	Username string
}

type CoinBalanceResponse struct { // API will return Response code and Balance for CoinBalance endpoint
	Code    int
	Balance int64
}

type Error struct { // API will return Response code and Error message for Error response
	Code    int
	Message string
}

func writerError(writer http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}

	writer.Header().Set("Content-Type", "application/json") // Set content type to JSON in packet header
	writer.WriteHeader(code)                                // Write HTTP status code to response header of packet

	json.NewEncoder(writer).Encode(response) // Encode response struct to JSON and write to response body to packet
}

var (
	RequestErrorHandler = func(writer http.ResponseWriter, err error) {
		writerError(writer, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(writer http.ResponseWriter) {
		writerError(writer, "Internal Error", http.StatusInternalServerError)
	}
)
