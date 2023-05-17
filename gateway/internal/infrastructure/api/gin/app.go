package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"meliarqsoft2/docs"
	"meliarqsoft2/internal/domain/model"
	"meliarqsoft2/internal/infrastructure/api"
	"meliarqsoft2/internal/infrastructure/api/gin/handler"
	"meliarqsoft2/internal/infrastructure/services"
	"strconv"
)

type MeliGinApp struct {
	port           int
	events         *api.Events
	userService    model.IUserService
	sellerService  model.ISellerService
	productService model.IProductService
}

func NewMeliAPI(
	port int,
	events *api.Events,
	userService model.IUserService,
	sellerService services.SellerHttpSyncService,
	productService services.ProductHttpSyncService) *MeliGinApp {
	return &MeliGinApp{
		port:           port,
		events:         events,
		userService:    userService,
		sellerService:  sellerService,
		productService: productService,
	}
}

func (app MeliGinApp) Run() error {

	route := gin.Default()

	// swagger config
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "MELI - API GATEWAY"
	docs.SwaggerInfo.Description = "Proyecto de Arquitectura de Software 2, UNQ 2023s1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + strconv.Itoa(app.port)
	docs.SwaggerInfo.BasePath = "/api/v1"

	basePath := route.Group("/api/v1")

	userRoute := basePath.Group("/users")
	userRoute.POST("", handler.NewGinUserRegister(app.userService).Execute)
	userRoute.GET("", handler.NewGinFindUser(app.userService).Execute)
	userRoute.PUT("/:id", handler.NewGinUpdateUser(app.userService).Execute)
	userRoute.DELETE("/:id", handler.NewGinUnregisterUser(app.userService).Execute)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(app.sellerService).Execute)
	sellerRoute.GET("", handler.NewGinFindSeller(app.sellerService).Execute)
	sellerRoute.PUT("/:id", handler.NewGinUpdateSeller(app.sellerService).Execute)
	// TODO sellerRoute.DELETE("/:id", handler.NewGinUnregisterSeller(app.events.UnregisterSellerEvent).Execute)

	productRoute := basePath.Group("/products")
	// TODO productRoute.POST("", handler.NewGinCreateProduct(app.events.CreateProductEvent).Execute)
	productRoute.GET("", handler.NewGinFindProduct(app.productService).Execute)
	productRoute.GET("/prices", handler.NewGinFilterProduct(app.productService).Execute)
	productRoute.PUT("/:id", handler.NewGinUpdateProduct(app.productService).Execute)
	// TODO productRoute.DELETE("/:id", handler.NewGinDeleteProduct(app.events.DeleteProductEvent).Execute)

	// TODO productRoute.POST("/purchases", handler.NewGinMakePurchase(app.events.MakePurchaseEvent).Execute)
	productRoute.GET("/:id/purchases", handler.NewGinFindPurchases(app.events.FindPurchasesFromProductEvent).Execute)
	// TODO "delete" purchases

	return route.Run(fmt.Sprintf(":%d", app.port))
}
