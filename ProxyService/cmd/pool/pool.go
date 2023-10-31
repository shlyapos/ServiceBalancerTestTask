package pool

import (
	"ProxyService/cmd/service"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Pool struct {
	services []*service.Service
	current  int
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) NextIndex() int {
	p.current += 1
	return p.current % len(p.services)
}

func (p *Pool) GetNextPeer() *service.Service {
	nextIdx := p.NextIndex()
	l := len(p.services) + nextIdx

	for i := nextIdx; i < l; i++ {
		idx := i % len(p.services)

		if p.services[idx].IsAlive() {
			if i != nextIdx {
				p.current = idx
			}

			return p.services[idx]
		}
	}

	return nil
}

func (p *Pool) AddService(newService *service.Service) {
	p.services = append(p.services, newService)
}

func isServiceAlive(url *url.URL) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url.String() + "/ping")

	if err != nil {
		return false
	}

	log.Printf("Service request status: %s\n", resp.Status)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func (p *Pool) CheckServiceAlive() {
	for _, s := range p.services {
		status := "up"
		alive := isServiceAlive(s.URL)
		s.SetAlive(alive)

		if !alive {
			status = "down"
		}

		log.Printf("%s [%s]\n", s.URL, status)
	}
}
