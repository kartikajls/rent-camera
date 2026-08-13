package service

type EmailService interface {
	SendEmail(to string, subject string, htmlContent string) (string, error)
}
