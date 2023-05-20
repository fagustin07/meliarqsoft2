package api

import (
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
)

type Events struct {
	RegisterSellerEvent   *action.RegisterSellerEvent
	UpdateSellerEvent     *action.UpdateSellerEvent
	UnregisterSellerEvent *action.UnregisterSellerEvent
	FindSellerEvent       *query.FindSellerEvent
	ExistSellerEvent      *query.ExistSeller

	RegisterCustomerEvent   *action.RegisterCustomerEvent
	UpdateCustomerEvent     *action.UpdateCustomerEvent
	UnregisterCustomerEvent *action.UnregisterCustomerEvent
	FindCustomerEvent       *query.FindCustomerEvent
	ExistCustomerEvent      *query.ExistCustomer

	SendNotificationEvent *action.SendNotificationEvent
}

func NewEvents(registerSellerEvent *action.RegisterSellerEvent, updateSellerEvent *action.UpdateSellerEvent, unregisterSellerEvent *action.UnregisterSellerEvent, findSellerEvent *query.FindSellerEvent, existSeller *query.ExistSeller, registerCustomerEvent *action.RegisterCustomerEvent, updateCustomerEvent *action.UpdateCustomerEvent, unregisterCustomerEvent *action.UnregisterCustomerEvent, findCustomerEvent *query.FindCustomerEvent, existCustomer *query.ExistCustomer, sendNotificationEvent *action.SendNotificationEvent) *Events {
	return &Events{
		RegisterSellerEvent:   registerSellerEvent,
		UpdateSellerEvent:     updateSellerEvent,
		UnregisterSellerEvent: unregisterSellerEvent,
		FindSellerEvent:       findSellerEvent,
		ExistSellerEvent:      existSeller,

		RegisterCustomerEvent:   registerCustomerEvent,
		UpdateCustomerEvent:     updateCustomerEvent,
		UnregisterCustomerEvent: unregisterCustomerEvent,
		FindCustomerEvent:       findCustomerEvent,
		ExistCustomerEvent:      existCustomer,

		SendNotificationEvent: sendNotificationEvent,
	}

}
