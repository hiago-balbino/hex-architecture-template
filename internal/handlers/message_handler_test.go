package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/domain"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/dto"
	"github.com/hiago-balbino/hex-architecture-template/internal/core/ports"
	"github.com/hiago-balbino/hex-architecture-template/pkg/apperrors"
	"github.com/hiago-balbino/hex-architecture-template/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateMessage_ShouldReturnErrorWhenBindingRequestParams(t *testing.T) {
	handler := setupHandler(nil)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/message").
		WithJSON(`{`).
		Expect().Status(http.StatusBadRequest).
		Body().Contains(apperrors.InvalidInput.Error())
}

func TestCreateMessage_ShouldReturnErrorWhenFailsToSetMessage(t *testing.T) {
	unexpectedError := errors.New("unexpected error")
	body := dto.CreateMessageRequest{Content: "message content"}

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Set", mock.Anything, body.Content).Return(domain.Message{}, unexpectedError)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/message").
		WithJSON(body).
		Expect().Status(http.StatusInternalServerError).
		Body().Contains(unexpectedError.Error())
}

func TestCreateMessage_ShouldSetMessageWithSuccess(t *testing.T) {
	body := dto.CreateMessageRequest{Content: "message content"}
	message := domain.NewMessage(uuid.NewString(), body.Content)

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Set", mock.Anything, message.Content).Return(message, nil)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	response := dto.CreateMessageResponse{}
	e := httpexpect.Default(t, server.URL)
	e.POST("/message").
		WithJSON(body).
		Expect().Status(http.StatusCreated).
		JSON().Decode(&response)

	assert.Equal(t, response.ID, message.ID)
}

func setupHandler(service ports.MessageServicer) *gin.Engine {
	handler := NewMessageHandler(service)
	server := Server{messagehdl: handler}
	router := server.setupRoutes()
	return router
}
