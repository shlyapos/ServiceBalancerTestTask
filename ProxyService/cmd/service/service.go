package service

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Service struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func NewService(url *url.URL, reverseProxy *httputil.ReverseProxy) *Service {
	return &Service{
		URL:          url,
		Alive:        true,
		ReverseProxy: reverseProxy,
	}
}

func (s *Service) SetAlive(aliveStatus bool) {
	s.mux.Lock()
	s.Alive = aliveStatus
	s.mux.Unlock()
}

func (s *Service) IsAlive() bool {
	s.mux.RLock()
	aliveStatus := s.Alive
	s.mux.RUnlock()

	return aliveStatus
}
