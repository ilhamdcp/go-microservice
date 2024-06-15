package main

import "net/http"

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.handleCreateOrder)
}

func (h *HttpHandler) handleCreateOrder(res http.ResponseWriter, req *http.Request) {
}
