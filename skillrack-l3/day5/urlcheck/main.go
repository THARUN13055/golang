package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type WebsiteResult struct {
	URL          string
	StatusCode   int
	IsReachable  bool
	HTTPSValid   bool
	ResponseTime time.Duration
	Error        string
}

func checkWebsite(url string) WebsiteResult {
	result := WebsiteResult{
		URL: url,
	}

	start := time.Now()

	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	resp, err := client.Get(url)
	result.ResponseTime = time.Since(start)

	if err != nil {
		result.Error = err.Error()
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.IsReachable = true

	if resp.TLS != nil {
		result.HTTPSValid = true
	}

	return result
}

func main() {
	websites := []string{
		"https://tefc.fit",
		"https://indishoppe.in",
		"https://admin.indishopp.in",
	}

	var wg sync.WaitGroup

	for _, website := range websites {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			result := checkWebsite(url)

			fmt.Printf("\nWebsite: %s\n", result.URL)
			fmt.Printf("Reachable: %v\n", result.IsReachable)
			fmt.Printf("HTTPS Working: %v\n", result.HTTPSValid)
			fmt.Printf("Status Code: %d\n", result.StatusCode)
			fmt.Printf("Response Time: %v\n", result.ResponseTime)

			if result.Error != "" {
				fmt.Printf("Error: %s\n", result.Error)
			}
		}(website)
	}

	wg.Wait()
}
