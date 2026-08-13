package repository

import (
	"errors"

	"gorm.io/gorm"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
)

type rentalOrderDetailRepository struct {
	db *gorm.DB
}

func NewRentalOrderDetailRepository(
	db *gorm.DB,
) RentalOrderDetailRepository {
	return &rentalOrderDetailRepository{
		db: db,
	}
}

func (r *rentalOrderDetailRepository) CreateDetail(orderID int64, req dto.CreateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error) {

	subtotal := req.RentalCost * float64(req.RentalDays)

	detail := entity.RentalOrderDetail{
		RentalOrderID: orderID,
		CameraID:      req.CameraID,
		RentalCost:    req.RentalCost,
		RentalDays:    req.RentalDays,
		Subtotal:      subtotal,
	}

	if err := r.db.Create(&detail).Error; err != nil {
		return nil, err
	}

	return &dto.RentalOrderDetailResponse{
		RentalDetailID: detail.RentalDetailID,
		RentalOrderID:  detail.RentalOrderID,
		CameraID:       detail.CameraID,
		RentalCost:     detail.RentalCost,
		RentalDays:     detail.RentalDays,
		Subtotal:       detail.Subtotal,
		CreatedAt:      detail.CreatedAt,
	}, nil
}

func (r *rentalOrderDetailRepository) GetDetailByID(detailID int64) (*dto.RentalOrderDetailResponse, error) {

	var detail entity.RentalOrderDetail

	err := r.db.
		Where("rental_detail_id = ?", detailID).
		First(&detail).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rental order detail not found")
		}

		return nil, err
	}

	return &dto.RentalOrderDetailResponse{
		RentalDetailID: detail.RentalDetailID,
		RentalOrderID:  detail.RentalOrderID,
		CameraID:       detail.CameraID,
		RentalCost:     detail.RentalCost,
		RentalDays:     detail.RentalDays,
		Subtotal:       detail.Subtotal,
		CreatedAt:      detail.CreatedAt,
	}, nil
}

func (r *rentalOrderDetailRepository) GetDetailsByOrderID(orderID int64) ([]dto.RentalOrderDetailResponse, error) {

	var details []entity.RentalOrderDetail

	err := r.db.
		Where("rental_order_id = ?", orderID).
		Order("rental_detail_id ASC").
		Find(&details).Error

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.RentalOrderDetailResponse,
		0,
		len(details),
	)

	for _, detail := range details {

		responses = append(
			responses,
			dto.RentalOrderDetailResponse{
				RentalDetailID: detail.RentalDetailID,
				RentalOrderID:  detail.RentalOrderID,
				CameraID:       detail.CameraID,
				RentalCost:     detail.RentalCost,
				RentalDays:     detail.RentalDays,
				Subtotal:       detail.Subtotal,
				CreatedAt:      detail.CreatedAt,
			},
		)
	}

	return responses, nil
}

func (r *rentalOrderDetailRepository) GetAllDetails() ([]dto.RentalOrderDetailResponse, error) {

	var details []entity.RentalOrderDetail

	err := r.db.
		Order("rental_detail_id ASC").
		Find(&details).Error

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.RentalOrderDetailResponse,
		0,
		len(details),
	)

	for _, detail := range details {

		responses = append(
			responses,
			dto.RentalOrderDetailResponse{
				RentalDetailID: detail.RentalDetailID,
				RentalOrderID:  detail.RentalOrderID,
				CameraID:       detail.CameraID,
				RentalCost:     detail.RentalCost,
				RentalDays:     detail.RentalDays,
				Subtotal:       detail.Subtotal,
				CreatedAt:      detail.CreatedAt,
			},
		)
	}

	return responses, nil
}

func (r *rentalOrderDetailRepository) UpdateDetail(detailID int64, req dto.UpdateRentalOrderDetailRequest) (*dto.RentalOrderDetailResponse, error) {

	var detail entity.RentalOrderDetail

	err := r.db.
		Where("rental_detail_id = ?", detailID).
		First(&detail).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rental order detail not found")
		}

		return nil, err
	}

	detail.CameraID = req.CameraID
	detail.RentalCost = req.RentalCost
	detail.RentalDays = req.RentalDays
	detail.Subtotal = req.RentalCost * float64(req.RentalDays)

	if err := r.db.Save(&detail).Error; err != nil {
		return nil, err
	}

	return &dto.RentalOrderDetailResponse{
		RentalDetailID: detail.RentalDetailID,
		RentalOrderID:  detail.RentalOrderID,
		CameraID:       detail.CameraID,
		RentalCost:     detail.RentalCost,
		RentalDays:     detail.RentalDays,
		Subtotal:       detail.Subtotal,
		CreatedAt:      detail.CreatedAt,
	}, nil
}

func (r *rentalOrderDetailRepository) DeleteDetail(detailID int64) error {

	result := r.db.
		Delete(
			&entity.RentalOrderDetail{},
			"rental_detail_id = ?",
			detailID,
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("rental order detail not found")
	}

	return nil
}
