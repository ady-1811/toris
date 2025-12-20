package ai

type CommandResponse struct {
	Command     []string `json:"commands"`
	Confidence  float64  `json:"confidence"`
	Instruction []string `json:"instructions"`
	RiskScore   int      `json:"risk_score"`
}
