package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configurar la conexión a MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Comprobar la conexión a la base de datos
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a MongoDB")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		newUUID, _ := uuid.NewUUID()
		c.JSON(200, gin.H{
			"uuid": newUUID,
		})
	})
	otherErr := r.Run()
	if otherErr != nil {
		return
	}
}
