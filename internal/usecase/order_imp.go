package usecase

import (
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/repository"
)

type rentalOrderUsecase struct {
	rentalOrderRepository repository.RentalOrderRepository
}

func NewRentalOrderUsecase(
	rentalOrderRepository repository.RentalOrderRepository,
) RentalOrderUsecase {

	return &rentalOrderUsecase{
		rentalOrderRepository: rentalOrderRepository,
	}
}

func (u *rentalOrderUsecase) CreateOrder(userID int64, req dto.CreateRentalOrderRequest) (*dto.RentalOrderResponse, error) {

	return u.rentalOrderRepository.CreateOrder(
		userID,
		req,
	)
}

func (u *rentalOrderUsecase) GetOrderByID(orderID int64) (*dto.RentalOrderResponse, error) {

	return u.rentalOrderRepository.GetOrderByID(orderID)
}

func (u *rentalOrderUsecase) GetOrdersByUserID(userID int64) ([]dto.RentalOrderResponse, error) {

	return u.rentalOrderRepository.GetOrdersByUserID(userID)
}

func (u *rentalOrderUsecase) GetAllOrders() ([]dto.RentalOrderResponse, error) {

	return u.rentalOrderRepository.GetAllOrders()
}

func (u *rentalOrderUsecase) UpdateOrderStatus(orderID int64, status string) error {

	return u.rentalOrderRepository.UpdateOrderStatus(
		orderID,
		status,
	)
}

func (u *rentalOrderUsecase) DeleteOrder(orderID int64) error {

	return u.rentalOrderRepository.DeleteOrder(orderID)
}
