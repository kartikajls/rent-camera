package usecase

import (
	"time"

	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/service"
)

type emailNotificationUsecase struct {
	emailRepository repository.EmailNotificationRepository
	emailService    service.EmailService
}

func NewEmailNotificationUsecase(
	emailRepository repository.EmailNotificationRepository,
	emailService service.EmailService,
) EmailNotificationUsecase {
	return &emailNotificationUsecase{
		emailRepository: emailRepository,
		emailService:    emailService,
	}
}

func (u *emailNotificationUsecase) SendRegistrationEmail(
	userID int64,
	email string,
	username string,
) error {

	notification := &entity.EmailNotification{
		UserID:           userID,
		RentalOrderID:    nil,
		Email:            email,
		NotificationType: "registration_confirmation",
		Subject:          "Registration Confirmation - Rent Camera",
		Status:           "pending",
	}

	if err := u.emailRepository.Create(notification); err != nil {
		return err
	}

	messageID, err := u.emailService.SendRegistrationEmail(
		email,
		username,
	)

	if err != nil {
		errorMessage := err.Error()

		notification.Status = "failed"
		notification.ErrorMessage = &errorMessage

		_ = u.emailRepository.Update(notification)

		return err
	}

	provider := "brevo"

	notification.Status = "sent"
	notification.Provider = &provider
	notification.ProviderMessageID = &messageID

	now := time.Now()
	notification.SentAt = &now

	return u.emailRepository.Update(notification)
}

func (u *emailNotificationUsecase) SendBookingConfirmationEmail(
	userID int64,
	email string,
	username string,
	orderID int64,
) error {

	notification := &entity.EmailNotification{
		UserID:           userID,
		RentalOrderID:    &orderID,
		Email:            email,
		NotificationType: "booking_confirmation",
		Subject:          "Booking Confirmation - Rent Camera",
		Status:           "pending",
	}

	if err := u.emailRepository.Create(notification); err != nil {
		return err
	}

	messageID, err := u.emailService.SendBookingConfirmationEmail(
		email,
		username,
		orderID,
	)

	if err != nil {
		errorMessage := err.Error()

		notification.Status = "failed"
		notification.ErrorMessage = &errorMessage

		_ = u.emailRepository.Update(notification)

		return err
	}

	provider := "brevo"

	notification.Status = "sent"
	notification.Provider = &provider
	notification.ProviderMessageID = &messageID

	now := time.Now()
	notification.SentAt = &now

	return u.emailRepository.Update(notification)
}

func (u *emailNotificationUsecase) SendPaymentConfirmationEmail(
	userID int64,
	email string,
	username string,
	orderID int64,
) error {

	notification := &entity.EmailNotification{
		UserID:           userID,
		RentalOrderID:    &orderID,
		Email:            email,
		NotificationType: "payment_confirmation",
		Subject:          "Payment Confirmation - Rent Camera",
		Status:           "pending",
	}

	if err := u.emailRepository.Create(notification); err != nil {
		return err
	}

	messageID, err := u.emailService.SendPaymentConfirmationEmail(
		email,
		username,
		orderID,
	)

	if err != nil {
		errorMessage := err.Error()

		notification.Status = "failed"
		notification.ErrorMessage = &errorMessage

		_ = u.emailRepository.Update(notification)

		return err
	}

	provider := "brevo"

	notification.Status = "sent"
	notification.Provider = &provider
	notification.ProviderMessageID = &messageID

	now := time.Now()
	notification.SentAt = &now

	return u.emailRepository.Update(notification)
}
