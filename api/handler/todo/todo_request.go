package todo

type TodoRequest struct {
	Title string `json:"title" validate:"required"`
}
