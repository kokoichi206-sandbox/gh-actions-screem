package handler

type ClientMsg struct {
	Type   int    `json:"type"`
	WorkId string `json:"work_id"`
}

type broadcastMsg struct {
	Type          int    `json:"type"`
	WorkId        string `json:"work_id"`
	StepNumber    int    `json:"step_number"`
	MaxSteps      int    `json:"max_steps"`
	StepCompleted bool   `json:"step_completed"`
	Done          bool   `json:"done"`
}
