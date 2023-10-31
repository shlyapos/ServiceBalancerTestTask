package handler

import "net/http"

type Handler struct {
	reqCounter          int
	reqCounterPerSecond int
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetReqCounter() int {
	return h.reqCounter
}

func (h *Handler) GetReqCounterPerSecond() int {
	temp := h.reqCounterPerSecond
	h.reqCounterPerSecond = 0

	return temp
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	h.reqCounter++
	h.reqCounterPerSecond++
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
