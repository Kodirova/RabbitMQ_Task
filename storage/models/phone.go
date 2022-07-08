package models

type Phone struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
}

type Response struct {
	Phone map[int64]string `json:"phone"`
}
