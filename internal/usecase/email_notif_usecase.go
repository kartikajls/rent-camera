package usecase

type EmailNotificationUsecase interface {
	SendRegistrationEmail(userID int64, email string, username string) error
	SendBookingConfirmationEmail(userID int64, email string, username string, orderID int64) error
	SendPaymentConfirmationEmail(userID int64, email string, username string, orderID int64) error
}
