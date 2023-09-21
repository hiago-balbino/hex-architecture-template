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

func TestGetMessage_ShouldReturnErrorWhenMessageNotFound(t *testing.T) {
	messageID := uuid.NewString()

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Get", mock.Anything, messageID).Return(domain.Message{}, apperrors.NotFound)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.GET("/message/{id}").
		WithPath("id", messageID).
		Expect().Status(http.StatusNotFound).
		Body().Contains(apperrors.NotFound.Error())
}

func TestGetMessage_ShouldReturnErrorWhenFailsToGetMessage(t *testing.T) {
	messageID := uuid.NewString()
	unexpectedError := errors.New("unexpected error")

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Get", mock.Anything, messageID).Return(domain.Message{}, unexpectedError)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.GET("/message/{id}").
		WithPath("id", messageID).
		Expect().Status(http.StatusInternalServerError).
		Body().Contains(unexpectedError.Error())
}

func TestGetMessage_ShouldReturnMessageWithSuccedd(t *testing.T) {
	message := domain.NewMessage(uuid.NewString(), "message content")

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Get", mock.Anything, message.ID).Return(message, nil)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	response := dto.GetMessageResponse{}
	e := httpexpect.Default(t, server.URL)
	e.GET("/message/{id}").
		WithPath("id", message.ID).
		Expect().Status(http.StatusOK).
		JSON().Decode(&response)

	assert.Equal(t, response.ID, message.ID)
	assert.Equal(t, response.Content, message.Content)
}

func TestGetMessages_ShouldReturnErrorWhenFailsToGetMessages(t *testing.T) {
	unexpectedError := errors.New("unexpected error")

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("GetAll", mock.Anything).Return([]domain.Message{}, unexpectedError)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.GET("/messages").
		Expect().Status(http.StatusInternalServerError).
		Body().Contains(unexpectedError.Error())
}

func TestGetMessages_ShouldReturnMessagesWithSuccess(t *testing.T) {
	messages := []domain.Message{
		domain.NewMessage(uuid.NewString(), "message content 1"),
		domain.NewMessage(uuid.NewString(), "message content 2"),
	}

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("GetAll", mock.Anything).Return(messages, nil)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)

	response := []dto.GetMessageResponse{}
	e := httpexpect.Default(t, server.URL)
	e.GET("/messages").
		Expect().Status(http.StatusOK).
		JSON().Decode(&response)

	assert.Equal(t, len(response), len(messages))
	assert.Equal(t, messages[0].ID, response[0].ID)
	assert.Equal(t, messages[0].Content, response[0].Content)
	assert.Equal(t, messages[1].ID, response[1].ID)
	assert.Equal(t, messages[1].Content, response[1].Content)
}

func TestDeleteMessage_ShouldReturnNoContentWhenMessageNotFound(t *testing.T) {
	messageID := uuid.NewString()

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Delete", mock.Anything, messageID).Return(apperrors.NotFound)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.DELETE("/message/{id}").
		WithPath("id", messageID).
		Expect().Status(http.StatusNoContent).
		Body().IsEmpty()
}

func TestDeleteMessage_ShouldReturnErrorWheFailsToDeleteMessage(t *testing.T) {
	messageID := uuid.NewString()
	unexpectedError := errors.New("unexpected error")

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Delete", mock.Anything, messageID).Return(unexpectedError)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.DELETE("/message/{id}").
		WithPath("id", messageID).
		Expect().Status(http.StatusInternalServerError).
		Body().Contains(unexpectedError.Error())
}

func TestDeleteMessage_ShouldReturnNoContentWhenDeleteMessageWithSuccess(t *testing.T) {
	messageID := uuid.NewString()

	serviceMock := new(mocks.MessageServicerMock)
	serviceMock.On("Delete", mock.Anything, messageID).Return(nil)

	handler := setupHandler(serviceMock)
	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.DELETE("/message/{id}").
		WithPath("id", messageID).
		Expect().Status(http.StatusNoContent).
		Body().IsEmpty()
}

func setupHandler(service ports.MessageServicer) *gin.Engine {
	handler := NewMessageHandler(service)
	server := Server{messagehdl: handler}
	router := server.setupRoutes()
	return router
}
