package usecase

import (
	"fmt"
	"time"

	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/service"
)

type emailNotificationUsecase struct {
	emailNotificationRepository repository.EmailNotificationRepository
	emailService                service.EmailService
}

func NewEmailNotificationUsecase(
	emailNotificationRepository repository.EmailNotificationRepository,
	emailService service.EmailService,
) EmailNotificationUsecase {
	return &emailNotificationUsecase{
		emailNotificationRepository: emailNotificationRepository,
		emailService:                emailService,
	}
}

// =====================================================
// REGISTRATION
// =====================================================

func (u *emailNotificationUsecase) SendRegistrationEmail(userID int64, email string, username string) error {

	subject := "Registrasi Berhasil"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>

			<h2>Welcome to Jati Rent Camera</h2>

			<p>Halo <strong>%s</strong>,</p>

			<p>
				Selamat! Akun kamu berhasil dibuat.
			</p>

			<p>
				Terima kasih telah bergabung dengan layanan kami.
			</p>

			<p>
				Regards,<br>
				Jati Rent Camera
			</p>

		</body>
		</html>
	`, username)

	return u.sendNotification(
		userID,
		nil,
		email,
		"registration",
		subject,
		htmlContent,
	)
}

// =====================================================
// BOOKING
// =====================================================

func (u *emailNotificationUsecase) SendBookingConfirmationEmail(userID int64, email string, username string, orderID int64) error {

	subject := "Booking Berhasil Dikonfirmasi"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>

			<h2>Booking Berhasil</h2>

			<p>Halo <strong>%s</strong>,</p>

			<p>
				Booking kamu telah berhasil dikonfirmasi.
			</p>

			<p>
				<strong>Order ID:</strong> %d
			</p>

			<p>
				Terima kasih telah menggunakan layanan kami.
			</p>

			<p>
				Regards,<br>
				Jati Rent Camera
			</p>

		</body>
		</html>
	`, username, orderID)

	rentalOrderID := orderID

	return u.sendNotification(
		userID,
		&rentalOrderID,
		email,
		"booking_confirmation",
		subject,
		htmlContent,
	)
}

// =====================================================
// PAYMENT
// =====================================================

func (u *emailNotificationUsecase) SendPaymentConfirmationEmail(userID int64, email string, username string, orderID int64) error {

	subject := "Pembayaran Berhasil"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>

			<h2>Pembayaran Berhasil</h2>

			<p>Halo <strong>%s</strong>,</p>

			<p>
				Pembayaran untuk order kamu telah berhasil.
			</p>

			<p>
				<strong>Order ID:</strong> %d
			</p>

			<p>
				Terima kasih telah melakukan pembayaran.
			</p>

			<p>
				Regards,<br>
				Jati Rent Camera
			</p>

		</body>
		</html>
	`, username, orderID)

	rentalOrderID := orderID

	return u.sendNotification(
		userID,
		&rentalOrderID,
		email,
		"payment_confirmation",
		subject,
		htmlContent,
	)
}

// =====================================================
// TOP UP SUCCESS
// =====================================================

func (u *emailNotificationUsecase) SendTopUpSuccess(
	userID int64,
	email string,
	amount float64,
) error {

	subject := "Top Up Berhasil"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>

			<h2>Top Up Berhasil</h2>

			<p>Halo,</p>

			<p>
				Top up kamu telah berhasil diproses.
			</p>

			<p>
				<strong>Jumlah Top Up:</strong>
				Rp %.2f
			</p>

			<p>
				Saldo kamu telah berhasil ditambahkan.
			</p>

			<p>
				Terima kasih telah menggunakan layanan kami.
			</p>

			<p>
				Regards,<br>
				Jati Rent Camera
			</p>

		</body>
		</html>
	`, amount)

	return u.sendNotification(
		userID,
		nil,
		email,
		"top_up_success",
		subject,
		htmlContent,
	)
}

// =====================================================
// TOP UP FAILED
// =====================================================

func (u *emailNotificationUsecase) SendTopUpFailed(userID int64, email string, amount float64) error {

	subject := "Top Up Ditolak"

	htmlContent := fmt.Sprintf(`
		<html>
		<body>

			<h2>Top Up Ditolak</h2>

			<p>Halo,</p>

			<p>
				Top up kamu tidak dapat diproses.
			</p>

			<p>
				<strong>Jumlah Top Up:</strong>
				Rp %.2f
			</p>

			<p>
				Silakan hubungi admin jika membutuhkan
				informasi lebih lanjut.
			</p>

			<p>
				Regards,<br>
				Jati Rent Camera
			</p>

		</body>
		</html>
	`, amount)

	return u.sendNotification(
		userID,
		nil,
		email,
		"top_up_failed",
		subject,
		htmlContent,
	)
}

// =====================================================
// SEND NOTIFICATION
// =====================================================

func (u *emailNotificationUsecase) sendNotification(userID int64, rentalOrderID *int64, email string, notificationType string, subject string, htmlContent string) error {

	notification := &entity.EmailNotification{
		UserID:           userID,
		RentalOrderID:    rentalOrderID,
		Email:            email,
		NotificationType: notificationType,
		Subject:          subject,
		Status:           "pending",
	}

	// Simpan notification sebagai pending
	if err := u.emailNotificationRepository.Create(notification); err != nil {
		return fmt.Errorf(
			"failed to create email notification: %w",
			err,
		)
	}

	// Kirim email melalui Brevo
	messageID, err := u.emailService.SendEmail(
		email,
		subject,
		htmlContent,
	)

	if err != nil {

		errorMessage := err.Error()

		notification.Status = "failed"
		notification.ErrorMessage = &errorMessage

		// Jangan mengganti error utama jika update gagal
		_ = u.emailNotificationRepository.Update(notification)

		return fmt.Errorf(
			"failed to send email: %w",
			err,
		)
	}

	// Email berhasil dikirim
	now := time.Now()
	provider := "brevo"

	notification.Status = "sent"
	notification.Provider = &provider
	notification.ProviderMessageID = &messageID
	notification.SentAt = &now

	if err := u.emailNotificationRepository.Update(notification); err != nil {
		return fmt.Errorf(
			"email sent but failed to update notification: %w",
			err,
		)
	}

	return nil
}
