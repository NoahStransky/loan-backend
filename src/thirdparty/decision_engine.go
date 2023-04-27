package thirdparty

import (
	"math/rand"
)

const (
	APPROVED      = "approved"
	REJECTED      = "rejected"
	GRANTVALUE    = "100"
	POSSIBLEVALUE = "60"
	DEFAULTVALUE  = "20"
)

type DecisionEngineRequest struct {
	Name            string
	YearEstablished string
	ProfitOrLoss    float64
	PreAssessment   string
}

type DecisionEngineResponse struct {
	Outcome string `json:"outcome"`
}

type DecisionEngine struct{}

func NewDecisionEngine() *DecisionEngine {
	return &DecisionEngine{}
}

func (e *DecisionEngine) MakeDecision(req *DecisionEngineRequest) *DecisionEngineResponse {
	result := &DecisionEngineResponse{
		Outcome: REJECTED,
	}
	// mock decision
	randomValue := rand.Intn(100)
	switch req.PreAssessment {
	case GRANTVALUE:
		{
			result.Outcome = APPROVED
			return result
		}
	case POSSIBLEVALUE:
		{
			if randomValue < 60 {
				result.Outcome = APPROVED
			}
			return result
		}
	case DEFAULTVALUE:
		{
			if randomValue < 20 {
				result.Outcome = APPROVED

			}
			return result
		}
	}
	return result
}
