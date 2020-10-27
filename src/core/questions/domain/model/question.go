package model

type Question struct {
	Text      string    `json:"text"`
	CreatedAt string    `json:"createdAt"`
	Choices   []Choices `json:"choices"`
}
