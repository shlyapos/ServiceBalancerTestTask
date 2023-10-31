package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	ticker := time.NewTicker(time.Millisecond * 100)

	requestIndex := 1

	for {
		select {
		case <-ticker.C:
			doRequest(client, requestIndex)
			requestIndex++
		}
	}
}

func doRequest(client *http.Client, requestNumber int) {
	req, err := http.NewRequest("GET", "http://localhost:9000/info", nil)

	if err != nil {
		log.Printf("Error while create request: %s\n", err.Error())
		return
	}

	res, err := client.Do(req)

	if err != nil {
		log.Printf("Error while request: %s\n", err.Error())
		return
	}

	log.Printf("Request - %d, response status: %s", requestNumber, res.Status)

	res.Body.Close()
}
