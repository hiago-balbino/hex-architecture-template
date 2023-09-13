package handlers

import (
	"github.com/gin-gonic/gin"
	service "github.com/hiago-balbino/hex-architecture-template/internal/core/usecases/message"
	repo "github.com/hiago-balbino/hex-architecture-template/internal/repositories/message"
)

type Server struct {
	messagehdl messageHandler
}

func NewServer() Server {
	messageRepository := repo.NewMessageRepository()
	messageService := service.NewMessageService(messageRepository)
	messageHandler := NewMessageHandler(messageService)

	return Server{
		messagehdl: messageHandler,
	}
}

func (s Server) Start() {
	router := gin.Default()
	router.POST("/message", s.messagehdl.createMessage)
	router.GET("/message/:id", s.messagehdl.getMessage)
	router.GET("/messages", s.messagehdl.getMessages)
	router.DELETE("/message/:id", s.messagehdl.deleteMessage)
	router.Run(":8080")
}
