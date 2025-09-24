package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/donnebaldemeca/GoAPI/api"
	"github.com/donnebaldemeca/GoAPI/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(writer http.ResponseWriter, request *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, request.URL.Query()) // Decode values from URL query string into 'params' struct

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(writer)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username) // Fetch coin details from DB
	if tokenDetails == nil {
		log.Error(err)
		api.RequestErrorHandler(writer, err)
		return
	}

	var httpResponse = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json") // Set content type to JSON in packet header
	err = json.NewEncoder(writer).Encode(httpResponse)      // Encode response struct to JSON and write to response body to packet
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}
}
