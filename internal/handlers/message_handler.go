package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/dto"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/ports"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
)

type messageHandler struct {
	service ports.MessageServicer
}

func NewMessageHandler(service ports.MessageServicer) messageHandler {
	return messageHandler{
		service: service,
	}
}

func (h messageHandler) createMessage(c *gin.Context) {
	var messageReqDto dto.CreateMessageRequest
	err := c.BindJSON(&messageReqDto)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.Join(apperrors.InvalidInput, err).Error()})
		return
	}

	message, err := h.service.Set(c.Request.Context(), messageReqDto.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, dto.BuildResponseCreateMessage(message.ID))
}

func (h messageHandler) getMessage(c *gin.Context) {
	messageID := c.Param("id")

	message, err := h.service.Get(c.Request.Context(), messageID)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseGetMessage(message))
}

func (h messageHandler) getMessages(c *gin.Context) {
	messages, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dto.BuildResponseGetMessages(messages))
}
