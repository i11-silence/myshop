package model

type Order struct {
	Username string  `json:"username"`
	Address  string  `json:"address"`
	Order    []Good  `json:"order"`
	Total    float64 `json:"total"`
}
