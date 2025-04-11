package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234AB",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456CDE",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789FGH",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 190,
		Username: "alex",
	},
	"json": {
		Coins: 150,
		Username: "jason",
	},
	"marie": {
		Coins: 180,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second*1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second*1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}