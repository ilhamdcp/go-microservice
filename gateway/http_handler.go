package main

import (
	"net/http"
	"strconv"

	c "github.com/ilhamdcp/go-microservice/common"
	pb "github.com/ilhamdcp/go-microservice/common/api"
)

type HttpHandler struct {
	client pb.OrderServiceClient
}

func NewHttpHandler(client pb.OrderServiceClient) *HttpHandler {
	return &HttpHandler{client: client}
}

func (h *HttpHandler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.handleCreateOrder)
}

func (h *HttpHandler) handleCreateOrder(res http.ResponseWriter, req *http.Request) {
	customerId, err := strconv.ParseInt(req.PathValue("customerId"), 10, 64)
	if err = c.ReadJson(req, nil); err != nil {
		c.WriteError(res, http.StatusBadRequest, "Failed to retrieve customerId")
		return
	}
	var orderProducts []*pb.OrderProduct
	if err = c.ReadJson(req, &orderProducts); err != nil {
		c.WriteError(res, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(req.Context(), &pb.CreateOrderRequest{CustomerId: customerId, OrderProducts: orderProducts})
}
