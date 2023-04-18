package api

import (
	action2 "meliarqsoft2/internal/domain/application/command/action"
	query2 "meliarqsoft2/internal/domain/application/command/query"
)

type Events struct {
	RegisterSellerEvent   *action2.RegisterSellerEvent
	UpdateSellerEvent     *action2.UpdateSellerEvent
	UnregisterSellerEvent *action2.UnregisterSellerEvent
	FindSellerEvent       *query2.FindSellerEvent

	RegisterUserEvent   *action2.RegisterUserEvent
	UpdateUserEvent     *action2.UpdateUserEvent
	UnregisterUserEvent *action2.UnregisterUserEvent
	FindUserEvent       *query2.FindUserEvent

	CreateProductEvent *action2.CreateProductEvent
	UpdateProductEvent *action2.UpdateProductEvent
	DeleteProductEvent *action2.DeleteProductEvent
	FindProductEvent   *query2.FindProductEvent
	FilterProductEvent *query2.FilterProductEvent

	MakePurchaseEvent             *action2.MakePurchaseEvent
	FindPurchasesFromProductEvent *query2.FindPurchasesFromProductEvent
}

func NewEvents(registerSellerEvent *action2.RegisterSellerEvent, updateSellerEvent *action2.UpdateSellerEvent, unregisterSellerEvent *action2.UnregisterSellerEvent, findSellerEvent *query2.FindSellerEvent, registerUserEvent *action2.RegisterUserEvent, updateUserEvent *action2.UpdateUserEvent, unregisterUserEvent *action2.UnregisterUserEvent, findUserEvent *query2.FindUserEvent, createProductEvent *action2.CreateProductEvent, updateProductEvent *action2.UpdateProductEvent, deleteProductEvent *action2.DeleteProductEvent, findProductEvent *query2.FindProductEvent, filterProductEvent *query2.FilterProductEvent, makePurchaseEvent *action2.MakePurchaseEvent, findPurchasesFromProductEvent *query2.FindPurchasesFromProductEvent) *Events {
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
