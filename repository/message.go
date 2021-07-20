package repository

import (
	"example/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepo interface {
	GetMessageById(c *gin.Context) (model.Message, error)
	// GetAllMessages(c *context.Context)
	// PostMessage(c *context.Context)
}

type MessageRepoImpl struct {
	Collection *mongo.Collection
}

func (m *MessageRepoImpl) GetMessageById(c *gin.Context) (model.Message, error) {
	result := model.Message{}
	err := m.Collection.FindOne(c, bson.M{"id": c.Param("id")}).Decode(&result)

	return result, err
}
