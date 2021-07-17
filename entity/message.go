package entity

type MessageResponse struct {
	Message string `json:"message"`
}

func Message(msg string) MessageResponse {
	return MessageResponse{msg}
}
