package repository

import (
	"p2-ip-kartikajls/internal/dto"
)

type RentalOrderDetailRepository interface {
	CreateDetail(orderID int64, req dto.CreateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error)
	GetDetailByID(detailID int64) (*dto.RentalOrderDetailResponse, error)
	GetDetailsByOrderID(orderID int64) ([]dto.RentalOrderDetailResponse, error)
	GetAllDetails() ([]dto.RentalOrderDetailResponse, error)
	UpdateDetail(detailID int64, req dto.UpdateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error)
	DeleteDetail(detailID int64) error
}
