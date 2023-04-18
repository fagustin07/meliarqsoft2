package gin

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"meliarqsoft2/docs"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin/handler"
	"net/http"
	"strconv"
)

type MeliGinApp struct {
	port   int
	events *api.Events
}

func NewMeliAPI(port int, events *api.Events) *MeliGinApp {
	return &MeliGinApp{
		port:   port,
		events: events,
	}
}

func (app MeliGinApp) Run() error {

	route := gin.Default()

	// swagger config
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "MELI - Arquitectura Hexagonal"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	basePath := route.Group("/api/v1")

	userRoute := basePath.Group("/users")
	userRoute.POST("", handler.NewGinUserRegister(app.events.RegisterUserEvent).Execute)
	userRoute.GET("", handler.NewGinFindUser(app.events.FindUserEvent).Execute)
	userRoute.PUT("/:id", handler.NewGinUpdateUser(app.events.UpdateUserEvent).Execute)
	userRoute.DELETE("/:id", handler.NewGinUnregisterUser(app.events.UnregisterUserEvent).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.events.RegisterSellerEvent).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.events.FindSellerEvent).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.events.UpdateSellerEvent).Execute)
	sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.events.UnregisterSellerEvent).Execute)

	productRoute := basePath.Group("/products")
	productRoute.POST("", handler.NewGinCreateProduct(app.events.CreateProductEvent).Execute)
	productRoute.GET("", handler.NewGinFindProduct(app.events.FindProductEvent).Execute)
	productRoute.PUT("/:id", handler.NewGinUpdateProduct(app.events.UpdateProductEvent).Execute)
	productRoute.DELETE("/:id", handler.NewGinDeleteProduct(app.events.DeleteProductEvent).Execute)
	productRoute.GET("/prices", handler.NewGinFilterProduct(app.events.FilterProductEvent).Execute)

	productRoute.POST("/purchases", handler.NewGinMakePurchase(app.events.MakePurchaseEvent).Execute)
	productRoute.GET("/:id/purchases", handler.NewGinFindPurchases(app.events.FindPurchasesFromProductEvent).Execute)

	return route.Run()
}

func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, bson.M{"message": "not implemented yet"})
}
