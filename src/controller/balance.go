package controller

import (
	"bernie.top/demyst/loan-backend/src/thirdparty"
)

type Balance struct {
	service *thirdparty.AccountingSoftware
}

func NewBalance() *Balance {
	accountingSoftwareService := thirdparty.NewAccountingSoftware("")
	return &Balance{
		service: accountingSoftwareService,
	}
}

// getBalance from balance engine
// mock from thirdparty module
func (c *Balance) GetBalanceByProvider(provider string) []thirdparty.BalanceSheet {
	return c.service.GetBalanceSheet(provider)
}
