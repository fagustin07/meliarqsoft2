package api

import (
	"meliarqsoft2/internal/application/command/action"
	"meliarqsoft2/internal/application/command/query"
)

type Events struct {
	RegisterSellerEvent   *action.RegisterSellerEvent
	UpdateSellerEvent     *action.UpdateSellerEvent
	UnregisterSellerEvent *action.UnregisterSellerEvent
	FindSellerEvent       *query.FindSellerEvent

	RegisterUserEvent   *action.RegisterUserEvent
	UpdateUserEvent     *action.UpdateUserEvent
	UnregisterUserEvent *action.UnregisterUserEvent
	FindUserEvent       *query.FindUserEvent

	CreateProductEvent *action.CreateProductEvent
	UpdateProductEvent *action.UpdateProductEvent
	DeleteProductEvent *action.DeleteProductEvent
	FindProductEvent   *query.FindProductEvent
	FilterProductEvent *query.FilterProductEvent

	MakePurchaseEvent             *action.MakePurchaseEvent
	FindPurchasesFromProductEvent *query.FindPurchasesFromProductEvent
}

func NewEvents(registerSellerEvent *action.RegisterSellerEvent, updateSellerEvent *action.UpdateSellerEvent, unregisterSellerEvent *action.UnregisterSellerEvent, findSellerEvent *query.FindSellerEvent, registerUserEvent *action.RegisterUserEvent, updateUserEvent *action.UpdateUserEvent, unregisterUserEvent *action.UnregisterUserEvent, findUserEvent *query.FindUserEvent, createProductEvent *action.CreateProductEvent, updateProductEvent *action.UpdateProductEvent, deleteProductEvent *action.DeleteProductEvent, findProductEvent *query.FindProductEvent, filterProductEvent *query.FilterProductEvent, makePurchaseEvent *action.MakePurchaseEvent, findPurchasesFromProductEvent *query.FindPurchasesFromProductEvent) *Events {
	return &Events{
		RegisterSellerEvent:   registerSellerEvent,
		UpdateSellerEvent:     updateSellerEvent,
		UnregisterSellerEvent: unregisterSellerEvent,
		FindSellerEvent:       findSellerEvent,

		RegisterUserEvent:   registerUserEvent,
		UpdateUserEvent:     updateUserEvent,
		UnregisterUserEvent: unregisterUserEvent,
		FindUserEvent:       findUserEvent,

		CreateProductEvent: createProductEvent,
		UpdateProductEvent: updateProductEvent,
		DeleteProductEvent: deleteProductEvent,
		FindProductEvent:   findProductEvent,
		FilterProductEvent: filterProductEvent,

		MakePurchaseEvent:             makePurchaseEvent,
		FindPurchasesFromProductEvent: findPurchasesFromProductEvent,
	}

}
