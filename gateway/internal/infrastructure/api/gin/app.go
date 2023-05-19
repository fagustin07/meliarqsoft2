package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meliarqsoft2/docs"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api/gin/handler"
	"meliarqsoft2/internal/infrastructure/services"
	"strconv"
)

type MeliGinGateway struct {
	port                int
	userService         model.ICustomerService
	sellerService       model.ISellerService
	productService      model.IProductService
	findPurchaseService model.IFindPurchaseService
}

func NewMeliAPIGateway(
	port int,
	userService model.ICustomerService,
	sellerService services.SellerHttpSyncService,
	productService services.ProductHttpSyncService,
	findPurchaseService services.FindPurchaseHttpSyncService,
) *MeliGinGateway {
	return &MeliGinGateway{
		port:                port,
		userService:         userService,
		sellerService:       sellerService,
		productService:      productService,
		findPurchaseService: findPurchaseService,
	}
}

func (app MeliGinGateway) Run() error {

	route := gin.Default()

	// swagger config
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "MELI - API GATEWAY"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	basePath := route.Group("/api/v1")

	customerRoute := basePath.Group("/customers")
	customerRoute.POST("", handler.NewGinCustomerRegister(app.userService).Execute)
	customerRoute.GET("", handler.NewGinFindCustomer(app.userService).Execute)
	customerRoute.PUT("/:id", handler.NewGinUpdateCustomer(app.userService).Execute)
	customerRoute.DELETE("/:id", handler.NewGinUnregisterCustomer(app.userService).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.sellerService).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.sellerService).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.sellerService).Execute)
	// TODO: si hacemos la coreografia de borrar a un vendedor, el borrar un producto con
	// las compras, quedan encapsulados dentro de la coreografia.
	// sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.events.UnregisterSellerEvent).Execute)

	productRoute := basePath.Group("/products")
	// TODO: coreo que chequea si existe el productor y luego crea el producto
	// productRoute.POST("", handler.NewGinCreateProduct(app.events.CreateProductEvent).Execute)
	productRoute.GET("", handler.NewGinFindProduct(app.productService).Execute)
	productRoute.GET("/prices", handler.NewGinFilterProduct(app.productService).Execute)
	productRoute.PUT("/:id", handler.NewGinUpdateProduct(app.productService).Execute)
	// TODO productRoute.DELETE("/:id", handler.NewGinDeleteProduct(app.events.DeleteProductEvent).Execute)

	// TODO: COMPRA DISTRIBUIDA, modelar el orquestador
	// productRoute.POST("/purchases", handler.NewGinMakePurchase(app.events.MakePurchaseEvent).Execute)
	productRoute.GET("/:id/purchases", handler.NewGinFindPurchases(app.findPurchaseService).Execute)
	// TODO "delete" purchases

	return route.Run(fmt.Sprintf(":%d", app.port))
}
