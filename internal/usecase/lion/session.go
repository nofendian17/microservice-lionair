package lion

import (
	"context"
	"github.com/google/uuid"
	"lion/internal/model/lion"
	"lion/internal/model/response"
	"net/http"
	"time"
)

func (l *lionUseCase) SessionCreate(ctx context.Context, request lion.SessionCreateRequest) (*response.Response, error) {
	sessionID := uuid.New().String()

	binarySecurityToken, err := l.sessionCreate.GetBinarySecurityToken(
		ctx,
		request.ConversationID,
		sessionID,
		request.Username,
		request.Password,
	)
	if err != nil {
		return nil, err
	}

	ttl := time.Duration(l.config.Integration.SessionTTL) * time.Minute
	if err := l.redisClient.Set(ctx, sessionID, binarySecurityToken, ttl); err != nil {
		return nil, err
	}

	validUntil := time.Now().Add(ttl)
	data := &lion.SessionCreateResponse{
		SessionID:  sessionID,
		ValidUntil: validUntil.Format("2006-01-02T15:04:05Z"),
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Count:   nil,
		Data:    data,
	}, nil
}

func (l *lionUseCase) SessionClose(ctx context.Context, request lion.SessionCloseRequest) (*response.Response, error) {
	binarySecurityToken, err := l.redisClient.Get(ctx, request.SessionID)
	if err != nil {
		return nil, err
	}

	if err = l.sessionClose.Logout(ctx, request.ConversationID, binarySecurityToken); err != nil {
		return nil, err
	}

	if err = l.redisClient.Delete(ctx, request.SessionID); err != nil {
		return nil, err
	}

	return &response.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Count:   nil,
		Data:    nil,
	}, nil
}
