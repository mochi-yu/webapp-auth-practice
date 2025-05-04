package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/webapp-auth-practice/entities"
	"github.com/mochi-yu/webapp-auth-practice/usecase"
)

type IMessageController interface {
	PostMessage(c *gin.Context)
	GetMessages(c *gin.Context)
}

type MessageController struct {
	mu usecase.IMessageUsecase
}

func NewMessageController(mu usecase.IMessageUsecase) IMessageController {
	return &MessageController{mu: mu}
}

func (mc *MessageController) PostMessage(c *gin.Context) {
	var req entities.MessageRequest
	c.BindJSON(&req)

	_, err := mc.mu.Create(c.Request.Context(), req.ToEntMessage())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message created successfully"})
}

func (mc *MessageController) GetMessages(c *gin.Context) {
	msgs, err := mc.mu.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": msgs})
}
