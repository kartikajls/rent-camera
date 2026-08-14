package usecase

import (
	"errors"
	"fmt"
	"strings"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/service"
)

type topUpUsecase struct {
	topUpRepository repository.TopUpRepository
	userRepository  repository.UserRepository
	whatsappService service.WhatsAppService
}

func NewTopUpUsecase(
	topUpRepository repository.TopUpRepository,
	userRepository repository.UserRepository,
	whatsappService service.WhatsAppService,
) TopUpUsecase {
	return &topUpUsecase{
		topUpRepository: topUpRepository,
		userRepository:  userRepository,
		whatsappService: whatsappService,
	}
}

// USER - CREATE TOP UP
func (u *topUpUsecase) Create(userID int64, request dto.CreateTopUpRequest) (*dto.TopUpResponse, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if request.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	topUp := &entity.TopUp{
		UserID: userID,
		Amount: request.Amount,
		Status: "pending",
	}

	err := u.topUpRepository.Create(topUp)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.GetByID(topUp.UserID)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf(
		"Halo %s,\n\n"+
			"Permintaan top up kamu berhasil dibuat.\n\n"+
			"Top Up ID: #%d\n"+
			"Nominal: Rp%.2f\n"+
			"Status: Pending\n\n"+
			"Silakan menunggu konfirmasi admin.\n\n"+
			"Rent Camera",
		user.Username,
		topUp.TopUpID,
		topUp.Amount,
	)

	if err := u.whatsappService.Send(
		user.Phone,
		message,
	); err != nil {
		fmt.Println("WhatsApp notification failed:", err)
	}

	return &dto.TopUpResponse{
		TopUpID: topUp.TopUpID,
		UserID:  topUp.UserID,
		Amount:  topUp.Amount,
		Status:  topUp.Status,
	}, nil
}

// GET BY ID
func (u *topUpUsecase) GetByID(id int64) (*dto.TopUpResponse, error) {

	if id <= 0 {
		return nil, errors.New("invalid top up id")
	}

	topUp, err := u.topUpRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	return &dto.TopUpResponse{
		TopUpID: topUp.TopUpID,
		UserID:  topUp.UserID,
		Amount:  topUp.Amount,
		Status:  topUp.Status,
	}, nil
}

// USER - GET BY USER ID
func (u *topUpUsecase) GetByUserID(userID int64) ([]dto.TopUpResponse, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	topUps, err := u.topUpRepository.GetByUserID(userID)

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.TopUpResponse,
		0,
		len(topUps),
	)

	for _, topUp := range topUps {

		responses = append(
			responses,
			dto.TopUpResponse{
				TopUpID: topUp.TopUpID,
				UserID:  topUp.UserID,
				Amount:  topUp.Amount,
				Status:  topUp.Status,
			},
		)
	}

	return responses, nil
}

// ADMIN - GET ALL
func (u *topUpUsecase) GetAll() ([]dto.TopUpResponse, error) {

	topUps, err := u.topUpRepository.GetAll()

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.TopUpResponse,
		0,
		len(topUps),
	)

	for _, topUp := range topUps {

		responses = append(
			responses,
			dto.TopUpResponse{
				TopUpID: topUp.TopUpID,
				UserID:  topUp.UserID,
				Amount:  topUp.Amount,
				Status:  topUp.Status,
			},
		)
	}

	return responses, nil
}

// ADMIN - UPDATE STATUS
func (u *topUpUsecase) UpdateStatus(id int64, status string) error {

	if id <= 0 {
		return errors.New("invalid top up id")
	}

	status = strings.ToLower(
		strings.TrimSpace(status),
	)

	switch status {
	case "pending", "success", "failed", "expired":
		// valid
	default:
		return errors.New("invalid top up status")
	}

	return u.topUpRepository.UpdateStatus(
		id,
		status,
	)
}

// ADMIN - APPROVE
func (u *topUpUsecase) Approve(id int64) error {

	if id <= 0 {
		return errors.New("invalid top up id")
	}

	// Cari top up
	topUp, err := u.topUpRepository.GetByID(id)
	if err != nil {
		return errors.New("top up not found")
	}

	// Hanya pending yang boleh di-approve
	if topUp.Status != "pending" {
		return errors.New(
			"top up has already been processed",
		)
	}

	// Cari user
	user, err := u.userRepository.GetByID(topUp.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	// Tambahkan saldo
	user.DepositAmount += topUp.Amount

	// Update saldo user
	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}

	// Ubah status top up menjadi success
	err = u.topUpRepository.UpdateStatus(
		id,
		"success",
	)
	if err != nil {
		return err
	}

	// WHATSAPP NOTIFICATION
	message := fmt.Sprintf(
		"Halo %s,\n\n"+
			"Top up kamu telah berhasil disetujui.\n\n"+
			"Top Up ID: #%d\n"+
			"Nominal: Rp%.2f\n"+
			"Status: Success\n"+
			"Saldo sekarang: Rp%.2f\n\n"+
			"Terima kasih telah menggunakan Rent Camera.",
		user.Username,
		topUp.TopUpID,
		topUp.Amount,
		user.DepositAmount,
	)

	if err := u.whatsappService.Send(
		user.Phone,
		message,
	); err != nil {
		fmt.Println(
			"WhatsApp notification failed:",
			err,
		)
	}

	return nil
}

// ADMIN - REJECT
func (u *topUpUsecase) Reject(id int64) error {

	if id <= 0 {
		return errors.New("invalid top up id")
	}

	// Cari top up
	topUp, err := u.topUpRepository.GetByID(id)
	if err != nil {
		return errors.New("top up not found")
	}

	// Hanya pending yang boleh di-reject
	if topUp.Status != "pending" {
		return errors.New(
			"top up has already been processed",
		)
	}

	// Cari user
	user, err := u.userRepository.GetByID(topUp.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	// Ubah status top up menjadi rejected
	err = u.topUpRepository.UpdateStatus(
		id,
		"rejected",
	)
	if err != nil {
		return err
	}

	// =========================================
	// WHATSAPP NOTIFICATION
	// =========================================

	message := fmt.Sprintf(
		"Halo %s,\n\n"+
			"Top up kamu ditolak.\n\n"+
			"Top Up ID: #%d\n"+
			"Nominal: Rp%.2f\n"+
			"Status: Rejected\n\n"+
			"Silakan periksa kembali data top up kamu "+
			"atau hubungi admin.\n\n"+
			"Rent Camera",
		user.Username,
		topUp.TopUpID,
		topUp.Amount,
	)

	if err := u.whatsappService.Send(
		user.Phone,
		message,
	); err != nil {
		fmt.Println(
			"WhatsApp notification failed:",
			err,
		)
	}

	return nil
}
