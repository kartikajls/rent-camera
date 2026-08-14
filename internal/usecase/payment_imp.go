package usecase

import (
	"fmt"
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/service"
)

type paymentUsecase struct {
	repository            repository.PaymentRepository
	userRepository        repository.UserRepository
	rentalOrderRepository repository.RentalOrderRepository
	whatsappService       service.WhatsAppService
}

func NewPaymentUsecase(
	repository repository.PaymentRepository,
	userRepository repository.UserRepository,
	rentalOrderRepository repository.RentalOrderRepository,
	whatsappService service.WhatsAppService,
) PaymentUsecase {
	return &paymentUsecase{
		repository:            repository,
		userRepository:        userRepository,
		rentalOrderRepository: rentalOrderRepository,
		whatsappService:       whatsappService,
	}
}

func (u *paymentUsecase) CreatePayment(userID int64, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {

	// Create payment
	payment, err := u.repository.CreatePayment(
		userID,
		req,
	)
	if err != nil {
		return nil, err
	}

	// Ambil data user untuk nomor WhatsApp
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		fmt.Println(
			"Failed to get user for WhatsApp notification:",
			err,
		)

		return payment, nil
	}

	// Pesan WhatsApp
	message := fmt.Sprintf(
		"Halo %s,\n\n"+
			"Pembayaran rental kamera berhasil dibuat.\n\n"+
			"Payment ID: #%d\n"+
			"Order ID: #%d\n"+
			"Nominal: Rp%.2f\n"+
			"Status: %s\n\n"+
			"Silakan menunggu konfirmasi pembayaran dari admin.\n\n"+
			"Rent Camera",
		user.Username,
		payment.PaymentID,
		payment.RentalOrderID,
		payment.Amount,
		payment.PaymentStatus,
	)

	// Kirim WhatsApp
	if err := u.whatsappService.Send(
		user.Phone,
		message,
	); err != nil {
		fmt.Println(
			"WhatsApp notification failed:",
			err,
		)
	}

	return payment, nil
}

func (u *paymentUsecase) GetPaymentByID(paymentID int64, userID int64) (*dto.PaymentResponse, error) {

	return u.repository.GetPaymentByID(
		paymentID,
		userID,
	)
}

func (u *paymentUsecase) GetPaymentsByUserID(userID int64) ([]dto.PaymentResponse, error) {

	return u.repository.GetPaymentsByUserID(
		userID,
	)
}

func (u *paymentUsecase) GetAllPayments() ([]dto.PaymentResponse, error) {

	return u.repository.GetAllPayments()
}

func (u *paymentUsecase) UpdatePaymentStatus(paymentID int64, req dto.UpdatePaymentStatusRequest) (*dto.PaymentResponse, error) {

	// Update payment status
	payment, err := u.repository.UpdatePaymentStatus(
		paymentID,
		req,
	)
	if err != nil {
		return nil, err
	}

	// Ambil rental order
	order, err := u.rentalOrderRepository.GetOrderByID(
		payment.RentalOrderID,
	)
	if err != nil {
		fmt.Println(
			"Failed to get rental order for WhatsApp notification:",
			err,
		)

		return payment, nil
	}

	// Ambil user
	user, err := u.userRepository.GetByID(
		order.UserID,
	)
	if err != nil {
		fmt.Println(
			"Failed to get user for WhatsApp notification:",
			err,
		)

		return payment, nil
	}

	// Pesan WhatsApp
	message := fmt.Sprintf(
		"Halo %s,\n\n"+
			"Status pembayaran rental kamera kamu telah diperbarui.\n\n"+
			"Payment ID: #%d\n"+
			"Order ID: #%d\n"+
			"Nominal: Rp%.2f\n"+
			"Metode Pembayaran: %s\n"+
			"Status: %s\n\n"+
			"Terima kasih telah menggunakan Rent Camera.",
		user.Username,
		payment.PaymentID,
		payment.RentalOrderID,
		payment.Amount,
		payment.PaymentMethod,
		payment.PaymentStatus,
	)

	// Kirim WhatsApp
	if err := u.whatsappService.Send(
		user.Phone,
		message,
	); err != nil {
		fmt.Println(
			"WhatsApp notification failed:",
			err,
		)
	}

	return payment, nil
}
