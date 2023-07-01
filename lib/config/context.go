package config

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// Request context is injected into each request and contains request specific information
type RequestContext struct {
	UserID uuid.UUID
}

func GetRequestContext(ctx context.Context) *RequestContext {
	return ctx.Value("_request_context").(*RequestContext)
}

func CreateRequestContext(ctx context.Context, payload *RequestContext) context.Context {
	return context.WithValue(ctx, "_request_context", payload)
}
