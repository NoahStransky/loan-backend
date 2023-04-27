package thirdparty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const BalanceFile = "file/balance_sheet.json"

type BalanceSheet struct {
	Year         string  `json:"year,omitempty"`
	Month        string  `json:"month,omitempty"`
	ProfitOrLoss float64 `json:"profitOrLoss,omitempty"`
	AssetsValue  float64 `json:"assetsValue,omitempty"`
}

type AccountingSoftware struct {
	Sheets map[string][]BalanceSheet
}

func NewAccountingSoftware(path string) *AccountingSoftware {
	if path == "" {
		path = BalanceFile
	}
	balanceSheet := readBalanceSheet(path)
	return &AccountingSoftware{
		Sheets: balanceSheet,
	}
}

func readBalanceSheet(path string) map[string][]BalanceSheet {
	var result map[string][]BalanceSheet
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println("Error opening Balance sheet", err)
		return result
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	return result
}

func (a *AccountingSoftware) GetBalanceSheet(providerName string) []BalanceSheet {
	response := a.Sheets[providerName]
	return response
}
