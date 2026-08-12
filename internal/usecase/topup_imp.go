package usecase

import (
	"errors"
	"strings"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/repository"
)

type topUpUsecase struct {
	topUpRepository repository.TopUpRepository
	userRepository  repository.UserRepository
}

func NewTopUpUsecase(
	topUpRepository repository.TopUpRepository,
	userRepository repository.UserRepository,
) TopUpUsecase {
	return &topUpUsecase{
		topUpRepository: topUpRepository,
		userRepository:  userRepository,
	}
}

// =====================================================
// USER - CREATE TOP UP
// =====================================================

func (u *topUpUsecase) Create(
	userID int64,
	request dto.CreateTopUpRequest,
) (*dto.TopUpResponse, error) {

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

	return &dto.TopUpResponse{
		TopUpID: topUp.TopUpID,
		UserID:  topUp.UserID,
		Amount:  topUp.Amount,
		Status:  topUp.Status,
	}, nil
}

// =====================================================
// GET BY ID
// =====================================================

func (u *topUpUsecase) GetByID(
	id int64,
) (*dto.TopUpResponse, error) {

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

// =====================================================
// USER - GET BY USER ID
// =====================================================

func (u *topUpUsecase) GetByUserID(
	userID int64,
) ([]dto.TopUpResponse, error) {

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

// =====================================================
// ADMIN - GET ALL
// =====================================================

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

// =====================================================
// ADMIN - UPDATE STATUS
// =====================================================

func (u *topUpUsecase) UpdateStatus(
	id int64,
	status string,
) error {

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

// =====================================================
// ADMIN - APPROVE
// =====================================================

func (u *topUpUsecase) Approve(
	id int64,
) error {

	// Cari top up
	topUp, err := u.topUpRepository.GetByID(id)

	if err != nil {
		return errors.New("top up not found")
	}

	// Pastikan masih pending
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

	return nil
}

// =====================================================
// ADMIN - REJECT
// =====================================================

func (u *topUpUsecase) Reject(
	id int64,
) error {

	topUp, err := u.topUpRepository.GetByID(id)

	if err != nil {
		return errors.New("top up not found")
	}

	// Hanya pending yang boleh ditolak
	if topUp.Status != "pending" {
		return errors.New(
			"top up has already been processed",
		)
	}

	return u.topUpRepository.UpdateStatus(
		id,
		"failed",
	)
}
