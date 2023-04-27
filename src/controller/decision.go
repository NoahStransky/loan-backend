package controller

import (
	"bernie.top/demyst/loan-backend/src/thirdparty"
)

const (
	ASSESSMENTGRANT    = "100"
	ASSESSMENTPOSSIBLE = "60"
	ASSESSMENTDEFAULT  = "20"
)

type BalanceSheetInfo struct {
	Year         string
	Month        string
	ProfitOrLoss float64
	AssetsValue  float64
}
type MakeDecisionRequestBody struct {
	LoanAmount      float64
	Name            string
	YearEstablished string
	BalanceSheet    []BalanceSheetInfo
}
type Decision struct {
	service *thirdparty.DecisionEngine
}

func NewDecision() *Decision {
	desicionEngineService := thirdparty.NewDecisionEngine()
	return &Decision{
		service: desicionEngineService,
	}
}

// getDecision from Decision engine
// mock from thirdparty module
func (c *Decision) MakeDecision(request *MakeDecisionRequestBody) *thirdparty.DecisionEngineResponse {
	profit := c.calculateProfitOrLossByYear(request.BalanceSheet)
	preAssessment := c.checkAssessment(profit, request.LoanAmount)
	serviceRequest := &thirdparty.DecisionEngineRequest{
		Name:            request.Name,
		YearEstablished: request.YearEstablished,
		ProfitOrLoss:    profit,
		PreAssessment:   preAssessment,
	}
	return c.service.MakeDecision(serviceRequest)
}

func (c *Decision) calculateProfitOrLossByYear(balances []BalanceSheetInfo) float64 {
	var sum float64
	sum = 0
	for _, value := range balances {
		sum += value.ProfitOrLoss
	}
	return sum / float64(len(balances))
}
func (c *Decision) checkAssessment(profit float64, loanAmount float64) string {
	if profit > loanAmount {
		return ASSESSMENTGRANT
	}
	if profit > 0 {
		return ASSESSMENTPOSSIBLE
	}
	return ASSESSMENTDEFAULT
}
