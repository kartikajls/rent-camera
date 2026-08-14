package service

type WasenderService interface {
	SendMessage(to string, text string) error
}
