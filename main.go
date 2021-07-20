package main

//go:generate mockgen -destination=mocks/mock_message_repository.go example/repository MessageRepo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"example/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-rest/docs"
)

var Db *mongo.Database

// @title go-rest example project
// @version 1.0
// @description This is a go-rest example project.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// require.NoError(t, err)
	defer client.Disconnect(ctx)

	api.Collection = client.Database("go-rest").Collection("messages")

	router := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/status", api.Status)
	router.GET("/:id", api.GetMessageById)
	router.GET("/all", api.GetAllMessages)
	router.POST("/", api.PostMessage)
	// router.GET("/test", api.Test)

	router.Run()
}
