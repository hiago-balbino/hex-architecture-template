package dto

type CreateMessageRequest struct {
	Content string `json:"content"`
}

type CreateMessageResponse struct {
	ID string `json:"id"`
}

func BuildResponseCreateMessage(id string) CreateMessageResponse {
	return CreateMessageResponse{
		ID: id,
	}
}
