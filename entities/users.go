package entities

type User struct {
	Uid      string `json:"uid,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
