package handler

import (
	"ProxyService/cmd/pool"
	"net/http"
)

type Handler struct {
	pool *pool.Pool
}

func NewHandler(newPool *pool.Pool) *Handler {
	return &Handler{
		pool: newPool,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if peer := h.pool.GetNextPeer(); peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	} else {
		http.Error(w, "Error: service not available", http.StatusServiceUnavailable)
	}

}
