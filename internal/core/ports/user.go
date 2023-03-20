package ports

import "github.com/gin-gonic/gin"

type IUserService interface {
	Create(name string, email string) error
}

type IUserRepository interface {
	Create(ID int, name string, email string) error
}

type IUserHandler interface {
	Create(c *gin.Context) error
}
