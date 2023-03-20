package ports

import "github.com/gin-gonic/gin"

type ISellerService interface {
	Create(businessName string, email string) error
}

type ISellerRepository interface {
	Create(ID int, businessName string, email string) error
}

type ISellerHandler interface {
	Create(c *gin.Context) error
}
