package usecase

import (
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/repository"
)

type rentalOrderDetailUsecase struct {
	repository repository.RentalOrderDetailRepository
}

func NewRentalOrderDetailUsecase(
	repository repository.RentalOrderDetailRepository,
) RentalOrderDetailUsecase {
	return &rentalOrderDetailUsecase{
		repository: repository,
	}
}

func (u *rentalOrderDetailUsecase) CreateDetail(orderID int64, req dto.CreateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error) {

	return u.repository.CreateDetail(
		orderID,
		req,
	)
}

func (u *rentalOrderDetailUsecase) GetDetailByID(detailID int64) (*dto.RentalOrderDetailResponse, error) {

	return u.repository.GetDetailByID(detailID)
}

func (u *rentalOrderDetailUsecase) GetDetailsByOrderID(orderID int64) ([]dto.RentalOrderDetailResponse, error) {

	return u.repository.GetDetailsByOrderID(orderID)
}

func (u *rentalOrderDetailUsecase) GetAllDetails() ([]dto.RentalOrderDetailResponse, error) {

	return u.repository.GetAllDetails()
}

func (u *rentalOrderDetailUsecase) UpdateDetail(detailID int64, req dto.UpdateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error) {

	return u.repository.UpdateDetail(
		detailID,
		req,
	)
}

func (u *rentalOrderDetailUsecase) DeleteDetail(detailID int64) error {

	return u.repository.DeleteDetail(detailID)
}
