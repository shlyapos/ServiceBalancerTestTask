package main

import (
	"BackendService/cmd/handler"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func main() {
	if err := loadConfigs(); err != nil {
		log.Fatalf("Error while load configs: %s", err.Error())
	}

	handler := handler.NewHandler()
	router := newRouter(handler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: router,
	}

	log.Println("Start service")

	go loggingRequests(handler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while listen and serve: %s", err.Error())
	}
}

func loadConfigs() error {
	viper.SetConfigFile("./config/config.yml")
	return viper.ReadInConfig()
}

func newRouter(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/ping", handler.Ping)
		r.Get("/info", handler.Info)
	})

	return router
}

func loggingRequests(handler *handler.Handler) {
	ticker := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			log.Printf("Requests count: %d, [Avg per second: %d]\n", handler.GetReqCounter(), handler.GetReqCounterPerSecond())
		}
	}
}
