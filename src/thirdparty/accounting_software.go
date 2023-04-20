package thirdparty

type BalanceSheet struct {
	Year         string  `json:"year,omitempty"`
	Month        string  `json:"month,omitempty"`
	ProfitOrLoss float64 `json:"profitOrLoss,omitempty"`
	AssetsValue  float64 `json:"assetsValue,omitempty"`
}

type AccountingSoftware struct{}

func (a *AccountingSoftware) makeDecision() []*BalanceSheet {
	response := make([]*BalanceSheet, 0)
	response = append(response, &BalanceSheet{
		Year:         "2020",
		Month:        "12",
		ProfitOrLoss: 100,
		AssetsValue:  1000,
	})
	return response
}
