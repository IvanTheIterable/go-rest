package api

import (
	"example/repository"
	"go-rest/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

// GetMessageById Status @Summary Get one message
// @Description get message
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {string} string	"ok"
// @Router /{id} [get]
func GetMessageById(c *gin.Context) {
	result, err := (&repository.MessageRepoImpl{Collection: Collection}).GetMessageById(c)

	if err != nil {
		log.Print("Error:", err)
		c.AbortWithStatus(404)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetAllMessages GetMessageById Status @Summary Get all messages
// @Description get all messages
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /all [get]
func GetAllMessages(c *gin.Context) {
	var result []model.Message

	messages, _ := Collection.Find(c, bson.D{{}})

	err := messages.All(c, &result)

	if err != nil {
		log.Print(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// PostMessage @Summary Post message
// @Description Post new message
// @Produce  json
// @Param message body model.Message true "Message"
// @Success 201 {string} string	"created"
// @Router / [post]
func PostMessage(c *gin.Context) {
	result := model.Message{}

	if c.BindJSON(&result) != nil {
		c.AbortWithStatus(400)
	}

	_, err := Collection.InsertOne(c, result)

	if err != nil {
		log.Print(err)
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}
