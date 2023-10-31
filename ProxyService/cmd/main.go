package main

import (
	"ProxyService/cmd/handler"
	"ProxyService/cmd/pool"
	"ProxyService/cmd/service"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

func main() {
	if err := loadConfigs(); err != nil {
		log.Fatalf("Error while load config: %s", err.Error())
	}

	servicePool := pool.NewPool()

	for i := 0; i < viper.GetInt("service_count"); i++ {
		url, err := url.Parse(fmt.Sprintf("http://%s%d:%d", viper.GetString("host"), i+1, viper.GetInt("proxy_port")))

		if err != nil {
			log.Fatalf("Error while parse service url: %s", err.Error())
		}

		rp := httputil.NewSingleHostReverseProxy(url)
		servicePool.AddService(service.NewService(url, rp))

		log.Printf("New proxy service was added on url: %s\n", url)
	}

	handler := handler.NewHandler(servicePool)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: handler,
	}

	go CheckAlive(*servicePool)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while listen and serve: %s", err.Error())
	}
}

func loadConfigs() error {
	viper.SetConfigFile("./config/config.yml")
	return viper.ReadInConfig()
}

func CheckAlive(pool pool.Pool) {
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			log.Println("Start service alive check")
			pool.CheckServiceAlive()
			log.Println("Service alive check completed")
		}
	}
}
