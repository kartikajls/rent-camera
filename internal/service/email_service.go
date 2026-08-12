package service

type EmailService interface {
	SendRegistrationEmail(
		toEmail string,
		username string,
	) (string, error)

	SendBookingConfirmationEmail(
		toEmail string,
		username string,
		orderID int64,
	) (string, error)

	SendPaymentConfirmationEmail(
		toEmail string,
		username string,
		orderID int64,
	) (string, error)
}
