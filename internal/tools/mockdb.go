package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"donne": {
		AuthToken: "A1B2C3",
		Username:  "donne",
	},
	"jen": {
		AuthToken: "D4E5F6",
		Username:  "jen",
	},
	"adam": {
		AuthToken: "G7H8I9",
		Username:  "adam",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"donne": {
		Coins:    100,
		Username: "donne",
	},
	"jen": {
		Coins:    250,
		Username: "jen",
	},
	"adam": {
		Coins:    500,
		Username: "adam",
	},
}

func (database *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1) // Simulate DB latency of 1 second

	var clientData = LoginDetails{}
	clientData, exists := mockLoginDetails[username]
	if !exists {
		return nil
	}
	return &clientData
}

func (database *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1) // Simulate DB latency of 1 second

	var coinData = CoinDetails{}
	coinData, exists := mockCoinDetails[username]
	if !exists {
		return nil
	}
	return &coinData
}

func (database *mockDB) SetupDatabase() error {
	return nil
}
