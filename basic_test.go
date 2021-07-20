package main

import (
	"context"
	"example/api"
	mocks "example/mocks"
	"example/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/status", nil)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/status", api.Status)
	r.ServeHTTP(w, req)

	fmt.Println("Body:", w.Body)
	assert.Equal(t, w.Code, 200)
}

func TestIntegrationTest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	require.NoError(t, err)
	defer client.Disconnect(ctx)

	api.Collection = client.Database("go-rest").Collection("messages")

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "69"}}

	api.GetMessageById(c)

	fmt.Println("Body:", w.Body)

	require.Equal(t, 200, w.Code)
	// assert.Equal(t, w.Code, 200) // or what value you need it to be
}

func TestUnitTest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockMessageRepo(mockCtrl)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	mockRepo.EXPECT().GetMessageById(c).Return(model.Message{Id: "mockedId", Date: 0, Payload: "Mocked Message"}, nil)

	res, _ := mockRepo.GetMessageById(c)

	require.Equal(t, "Mocked Message", res.Payload)
}
