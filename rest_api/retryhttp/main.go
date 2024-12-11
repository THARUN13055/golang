package main

import (
	"log"
	"net/http"

	"github.com/avast/retry-go"
)

func SendRequest() error {
	var resp *http.Response
	err := retry.Do(
		func() error {
			var err error
			resp, err = http.Get("www.google.com")
			return err
		},
		retry.Attempts(3),
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retrying request after error: %v", err)
		}),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	err := SendRequest()
	if err != nil {
		log.Fatalf("Request failed after retries: %v", err)
	} else {
		log.Println("Request succeeded")
	}
}