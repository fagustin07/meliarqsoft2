package api

import (
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
)

type PurchaseEvents struct {
	MakePurchaseEvent             *action.MakePurchaseEvent
	FindPurchasesFromProductEvent *query.FindPurchasesFromProductEvent
	UndoPurchasesFromProductEvent *action.UndoPurchasesByProductEvent
}

func NewPurchaseEvents(undoPurchaseEvent *action.UndoPurchasesByProductEvent, makePurchaseEvent *action.MakePurchaseEvent, findPurchasesFromProductEvent *query.FindPurchasesFromProductEvent) *PurchaseEvents {
	return &PurchaseEvents{
		MakePurchaseEvent:             makePurchaseEvent,
		FindPurchasesFromProductEvent: findPurchasesFromProductEvent,
		UndoPurchasesFromProductEvent: undoPurchaseEvent,
	}
}
