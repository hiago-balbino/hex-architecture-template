package handlers

import (
	"github.com/gin-gonic/gin"
	service "github.com/hiago-balbino/hex-architecture-template/internal/core/usecases/message"
	repo "github.com/hiago-balbino/hex-architecture-template/internal/repositories/message"
	"github.com/hiago-balbino/hex-architecture-template/pkg/identifier"
)

type Server struct {
	messagehdl messageHandler
}

func NewServer() Server {
	uuidGenerator := identifier.NewUUIDGenerator()
	messageRepository := repo.NewMessageRepository()
	messageService := service.NewMessageService(uuidGenerator, messageRepository)
	messageHandler := NewMessageHandler(messageService)

	return Server{
		messagehdl: messageHandler,
	}
}

func (s Server) Start() {
	router := s.setupRoutes()
	router.Run(":8080")
}

func (s Server) setupRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/message", s.messagehdl.createMessage)
	router.GET("/message/:id", s.messagehdl.getMessage)
	router.GET("/messages", s.messagehdl.getMessages)
	router.DELETE("/message/:id", s.messagehdl.deleteMessage)
	return router
}
