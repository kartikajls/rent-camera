package dto

type TopUpResponse struct {
	TopUpID int64   `json:"top_up_id"`
	UserID  int64   `json:"user_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}
