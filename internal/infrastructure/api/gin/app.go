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
	port   int         `required:"true" default:"8080"`
	events *api.Events `required:"true"`
}

func NewMeliAPI(port int, events *api.Events) *MeliGinApp {
	return &MeliGinApp{
		port:   port,
		events: events,
	}
}

func (app MeliGinApp) Run(events *api.Events) error {

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
	userRoute.POST("", handler.NewGinUserRegister(events.RegisterUserEvent).Execute)
	userRoute.GET("", NotImplemented)
	userRoute.PUT("/:id", NotImplemented)
	userRoute.DELETE("/:id", NotImplemented)

	sellerRoute := basePath.Group("/sellers")
	sellerRoute.POST("", handler.NewGinRegisterSeller(events.RegisterSellerEvent).Execute)
	sellerRoute.PUT("/:id", NotImplemented)
	sellerRoute.DELETE("/:id", NotImplemented)
	sellerRoute.GET("", NotImplemented)

	productRoute := basePath.Group("/products")
	productRoute.POST("", NotImplemented)
	productRoute.PUT("/:id", NotImplemented)
	productRoute.DELETE("/:id", NotImplemented)
	productRoute.GET("", NotImplemented)
	productRoute.GET("/prices", NotImplemented)
	productRoute.POST("/purchases", NotImplemented)
	productRoute.GET("/:id/purchases", NotImplemented)

	return route.Run()
}

func NotImplemented(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, bson.M{"message": "not implemented yet"})
}
