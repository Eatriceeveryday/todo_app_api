package entities

type Todo struct {
	TodoId string `json:"todo_id"`
	Title  string `json:"title"`
	UserId string `json:"user_id,omitempty"`
}
