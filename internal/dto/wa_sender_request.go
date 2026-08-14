package dto

type SendWARequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
