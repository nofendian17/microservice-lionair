package lion

type SessionCreateRequest struct {
	ConversationID string `json:"conversationID" validate:"required"`
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

type SessionCreateResponse struct {
	SessionID  string `json:"sessionID"`
	ValidUntil string `json:"validUntil"`
}

type SessionCloseRequest struct {
	ConversationID string `json:"conversationID" validate:"required"`
	SessionID      string `json:"sessionID" validate:"required,uuid"`
}
