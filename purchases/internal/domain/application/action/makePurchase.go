package action

import (
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/dto"
)

type MakePurchaseEvent struct {
	repository          model.IPurchaseRepository
	notificationService model.INotificationService
}

func NewMakePurchaseEvent(repo model.IPurchaseRepository, service model.INotificationService) *MakePurchaseEvent {
	return &MakePurchaseEvent{
		repository:          repo,
		notificationService: service,
	}
}

func (actionEvent MakePurchaseEvent) Execute(purchaseReq *dto.CreatePurchase) (*model.Purchase, error) {
	_, err := actionEvent.repository.Create(purchaseReq.Purchase)

	if err != nil {
		return nil, err
	}

	_ = actionEvent.notificationService.Send(purchaseReq.SellerNotification)
	_ = actionEvent.notificationService.Send(purchaseReq.BuyerNotification)

	return purchaseReq.Purchase, nil
}
