package ai 

type CommandResponse struct {
	Command string `json:"command"`
	Confidence float64 `json:"confidence"`
}