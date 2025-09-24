package api

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
