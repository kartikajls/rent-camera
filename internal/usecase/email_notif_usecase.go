package usecase

type EmailNotificationUsecase interface {
	SendRegistrationEmail(userID int64, email string, username string) error
	SendBookingConfirmationEmail(userID int64, email string, username string, orderID int64) error
	SendPaymentConfirmationEmail(userID int64, email string, username string, orderID int64) error
	SendTopUpSuccess(userID int64, email string, amount float64) error
	SendTopUpFailed(userID int64, email string, amount float64) error
}
