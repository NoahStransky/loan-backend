package thirdparty

const (
	APPROVED = "approved"
	REJECTED = "rejected"
)

type DecisionEngineResponse struct {
	Outcome string `json:"outcome"`
}

type DecisionEngine struct{}

func (e *DecisionEngine) makeDecision() *DecisionEngineResponse {
	return &DecisionEngineResponse{
		Outcome: "approved",
	}
}
