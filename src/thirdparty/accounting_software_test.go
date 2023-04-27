package thirdparty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestNewAccountingSoftware(t *testing.T) {
	fakeBalanceSheet := map[string][]BalanceSheet{
		"Provider1": {{
			Year:         "2023",
			Month:        "12",
			ProfitOrLoss: 100,
			AssetsValue:  1000,
		},
		},
		"Provider2": {{
			Year:         "2023",
			Month:        "11",
			ProfitOrLoss: 500,
			AssetsValue:  2000},
		},
	}
	fakeJson, _ := json.Marshal(fakeBalanceSheet)
	ioutil.WriteFile("fake_balance_sheet.json", fakeJson, 0644)
	defer os.Remove("fake_balance_sheet.json")

	accountingSoftware := NewAccountingSoftware("fake_balance_sheet.json")
	fmt.Println("accountingSoftware ", accountingSoftware)

	if len(accountingSoftware.Sheets) != 2 {
		t.Errorf("Expected 2 balance sheets, got %d", len(accountingSoftware.Sheets))
	}

	provider1Sheet := accountingSoftware.Sheets["Provider1"][0]
	if provider1Sheet.Year != "2023" || provider1Sheet.Month != "12" || provider1Sheet.ProfitOrLoss != 100 || provider1Sheet.AssetsValue != 1000 {
		t.Errorf("Unexpected balance sheet data for Provider1: %+v", provider1Sheet)
	}

	provider2Sheet := accountingSoftware.Sheets["Provider2"][0]
	if provider2Sheet.Year != "2023" || provider2Sheet.Month != "11" || provider2Sheet.ProfitOrLoss != 500 || provider2Sheet.AssetsValue != 2000 {
		t.Errorf("Unexpected balance sheet data for Provider2: %+v", provider2Sheet)
	}
}
func TestMakeDecision(t *testing.T) {
	fakeBalanceSheet := map[string][]BalanceSheet{
		"Provider1": {{
			Year:         "2023",
			Month:        "12",
			ProfitOrLoss: 100,
			AssetsValue:  1000},
		},
		"Provider2": {{
			Year:         "2023",
			Month:        "11",
			ProfitOrLoss: 500,
			AssetsValue:  2000,
		},
		},
	}

	accountingSoftware := &AccountingSoftware{
		Sheets: fakeBalanceSheet,
	}

	response := accountingSoftware.GetBalanceSheet("Provider1")

	expectedResponse := []BalanceSheet{fakeBalanceSheet["Provider1"][0]}
	if len(response) != 1 || response[0] != expectedResponse[0] {
		t.Errorf("Unexpected response for GetBalanceSheet(\"Provider1\"): %+v", response)
	}

	response = accountingSoftware.GetBalanceSheet("Provider2")

	expectedResponse = []BalanceSheet{fakeBalanceSheet["Provider2"][0]}
	if len(response) != 1 || response[0] != expectedResponse[0] {
		t.Errorf("Unexpected response for MakeDecision(\"Provider2\"): %+v", response)
	}
}
