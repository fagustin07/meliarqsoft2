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

	RegisterUserEvent   *action.RegisterUserEvent
	UpdateUserEvent     *action.UpdateUserEvent
	UnregisterUserEvent *action.UnregisterUserEvent
	FindUserEvent       *query.FindUserEvent
	ExistUserEvent      *query.ExistUser

	SendNotificationEvent *action.SendNotificationEvent
}

func NewEvents(registerSellerEvent *action.RegisterSellerEvent, updateSellerEvent *action.UpdateSellerEvent, unregisterSellerEvent *action.UnregisterSellerEvent, findSellerEvent *query.FindSellerEvent, existSeller *query.ExistSeller, registerUserEvent *action.RegisterUserEvent, updateUserEvent *action.UpdateUserEvent, unregisterUserEvent *action.UnregisterUserEvent, findUserEvent *query.FindUserEvent, existUser *query.ExistUser, sendNotificationEvent *action.SendNotificationEvent) *Events {
	return &Events{
		RegisterSellerEvent:   registerSellerEvent,
		UpdateSellerEvent:     updateSellerEvent,
		UnregisterSellerEvent: unregisterSellerEvent,
		FindSellerEvent:       findSellerEvent,
		ExistSellerEvent:      existSeller,

		RegisterUserEvent:   registerUserEvent,
		UpdateUserEvent:     updateUserEvent,
		UnregisterUserEvent: unregisterUserEvent,
		FindUserEvent:       findUserEvent,
		ExistUserEvent:      existUser,

		SendNotificationEvent: sendNotificationEvent,
	}

}
