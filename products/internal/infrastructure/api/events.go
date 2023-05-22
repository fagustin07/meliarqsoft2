package api

import (
	"meliarqsoft2/internal/domain/application/action"
	"meliarqsoft2/internal/domain/application/query"
)

type Events struct {
	CreateProductEvent     *action.CreateProductEvent
	UpdateProductEvent     *action.UpdateProductEvent
	DeleteProductEvent     *action.DeleteProductEvent
	FindProductEvent       *query.FindProductEvent
	FilterProductEvent     *query.FilterProductEvent
	RestoreProductsEvent   *action.RestoreProductsEvent
	DeleteProductsBySeller *action.DeleteProductsBySeller
	FindProductById        *query.FindProductById
}

func NewEvents(createProductEvent *action.CreateProductEvent, updateProductEvent *action.UpdateProductEvent, deleteProductEvent *action.DeleteProductEvent, findProductEvent *query.FindProductEvent, filterProductEvent *query.FilterProductEvent, restoreProductsEvent *action.RestoreProductsEvent, deleteBySellerEvent *action.DeleteProductsBySeller, findById *query.FindProductById) *Events {
	return &Events{
		CreateProductEvent:     createProductEvent,
		UpdateProductEvent:     updateProductEvent,
		DeleteProductEvent:     deleteProductEvent,
		FindProductEvent:       findProductEvent,
		FilterProductEvent:     filterProductEvent,
		RestoreProductsEvent:   restoreProductsEvent,
		DeleteProductsBySeller: deleteBySellerEvent,
		FindProductById:        findById,
	}

}
