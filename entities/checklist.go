package entities

type Checklist struct {
	CheckId     string `json:"check_id"`
	Detail      string `json:"detail"`
	IsCompleted bool   `json:"is_completed"`
}
