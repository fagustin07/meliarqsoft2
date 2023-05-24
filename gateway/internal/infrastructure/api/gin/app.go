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
	"net/http"
	"strconv"
)

type MeliGinGateway struct {
	port                int
	userService         model.ICustomerService
	sellerService       model.ISellerService
	productService      model.IProductService
	findPurchaseService model.IFindPurchaseService
	makePurchaseService model.IMakePurchaseService
}

func NewMeliAPIGateway(
	port int,
	userService model.ICustomerService,
	sellerService services.SellerHttpSyncService,
	productService services.ProductHttpSyncService,
	findPurchaseService services.FindPurchaseHttpSyncService,
	makePurchaseService services.MakePurchaseHttpSyncService,
) *MeliGinGateway {
	return &MeliGinGateway{
		port:                port,
		userService:         userService,
		sellerService:       sellerService,
		productService:      productService,
		findPurchaseService: findPurchaseService,
		makePurchaseService: makePurchaseService,
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
	basePath.GET("/heartbeat", heartbeat)

	customerRoute := basePath.Group("/customers")
	customerRoute.POST("", handler.NewGinCustomerRegister(app.userService).Execute)
	customerRoute.GET("", handler.NewGinFindCustomer(app.userService).Execute)
	customerRoute.PUT("/:id", handler.NewGinUpdateCustomer(app.userService).Execute)
	customerRoute.DELETE("/:id", handler.NewGinUnregisterCustomer(app.userService).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.sellerService).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.sellerService).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.sellerService).Execute)
	sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.sellerService).Execute)

	productRoute := basePath.Group("/products")
	productRoute.POST("", handler.NewGinCreateProduct(app.productService).Execute)
	productRoute.GET("", handler.NewGinFindProduct(app.productService).Execute)
	productRoute.GET("/prices", handler.NewGinFilterProduct(app.productService).Execute)
	productRoute.PUT("/:id", handler.NewGinUpdateProduct(app.productService).Execute)
	productRoute.DELETE("/:id", handler.NewGinDeleteProduct(app.productService).Execute)

	purchaseRoute := basePath.Group("/purchases")
	purchaseRoute.POST("/products", handler.NewGinMakePurchase(app.makePurchaseService).Execute)
	purchaseRoute.GET("/products/:id", handler.NewGinFindPurchases(app.findPurchaseService).Execute)

	return route.Run(fmt.Sprintf(":%d", app.port))
}

func heartbeat(context *gin.Context) {
	context.JSON(http.StatusOK, "API GATEWAY IS WORKING!")
}
