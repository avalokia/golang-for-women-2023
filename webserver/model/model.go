package model

type User struct {
	Email   string `json:"email"`
	Address string `json:"address"`
	Job     string `json:"job"`
	Reason  string `json:"reason"`
}
