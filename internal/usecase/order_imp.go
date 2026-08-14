package usecase

import (
	"errors"
	"fmt"
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/repository"
	"p2-ip-kartikajls/internal/service"
)

type rentalOrderUsecase struct {
	rentalOrderRepository repository.RentalOrderRepository
	cameraRepository      repository.CameraRepository
	userRepository        repository.UserRepository
	whatsappService       service.WhatsAppService
}

func NewRentalOrderUsecase(
	rentalOrderRepository repository.RentalOrderRepository,
	cameraRepository repository.CameraRepository,
	userRepository repository.UserRepository,
	whatsappService service.WhatsAppService,
) RentalOrderUsecase {

	return &rentalOrderUsecase{
		rentalOrderRepository: rentalOrderRepository,
		cameraRepository:      cameraRepository,
		userRepository:        userRepository,
		whatsappService:       whatsappService,
	}
}

func (u *rentalOrderUsecase) CreateOrder(userID int64, req dto.CreateRentalOrderRequest) (*dto.RentalOrderResponse, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if req.CameraID <= 0 {
		return nil, errors.New("invalid camera id")
	}

	camera, err := u.cameraRepository.GetByID(req.CameraID)
	if err != nil {
		return nil, errors.New("camera not found")
	}

	if !camera.Available {
		return nil, errors.New("camera is not available")
	}

	totalAmount := camera.RentalCost

	order, err := u.rentalOrderRepository.CreateOrder(
		userID,
		req.CameraID,
		totalAmount,
	)
	if err != nil {
		return nil, err
	}

	// WhatsApp notification
	user, err := u.userRepository.GetByID(userID)
	if err == nil {

		message := fmt.Sprintf(
			"Halo %s,\n\n"+
				"Order rental kamera kamu berhasil dibuat.\n\n"+
				"Order ID: #%d\n"+
				"Kamera: %s\n"+
				"Total: Rp%.2f\n"+
				"Status: %s\n\n"+
				"Rent Camera",
			user.Username,
			order.RentalOrderID,
			camera.Name,
			order.TotalAmount,
			order.Status,
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
	}

	return order, nil
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
