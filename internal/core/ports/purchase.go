package ports

import (
	"github.com/gin-gonic/gin"
	"meliarqsoft2/internal/core/domain"
)

type IPurchaseService interface {
	Create(IDProduct int, IDUser int) error
	Delete(IDProduct int, IDUser int) error
	Find(IDProduct int, IDUser int) (domain.Product, domain.User)
}

type IPurchaseRepository interface {
	Create(IDProduct int, IDUser int) error
	Delete(IDProduct int, IDUser int) error
	Find(IDProduct int, IDUser int) (domain.Product, domain.User)
}

type IPurchaseHandler interface {
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Find(c *gin.Context) (domain.Product, domain.User)
}
