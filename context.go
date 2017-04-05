package main

import (
	"context"
	"net/http"
)

// ContextHandler adds a context to HTTPHandlers
type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

// ContextHandlerFunc types the ContextHandler interface method
type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

// ContextHandlerFuncts implement the ContextHandler interface
func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

// ContextAdapter allows adaptation into different types of handlers
type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

// ContextAdapter conforms to an HTTPHandler
func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}
