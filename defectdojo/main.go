package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blackaichi/go-defectdojo/defectdojo"
)

type ScanResult struct {
	FilePath    string `json:"file_path"`
	Severity    string `json:"severity"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	url := "http://13.201.28.37:8080"
	authorizedToken := "798b581a65f991dc728d2ae8745a77f6459078bc"

	dj, err := defectdojo.NewDojoClient(url, authorizedToken, nil)
	if err != nil {
		log.Fatalf("Failed to create DefectDojo client: %v", err)
	}
	cxt := context.Background()

	// Fetch engagements
	engagements, err := dj.Engagements.List(cxt, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve engagements: %v", err)
	}

	if engagements.Results != nil {
		for _, engagement := range *engagements.Results {
			var id int
			var name string

			if engagement.Id != nil {
				id = *engagement.Id
			}

			if engagement.Name != nil {
				name = *engagement.Name
			}

			fmt.Printf("Engagement ID: %d, Name: %s\n", id, name)
		}
	} else {
		fmt.Println("No engagements found")
	}

	filePath := "/home/tharun/axonaio/sample/paynpro/result/trivy_output.json"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %v", err)
	}

	// Create ImportScan struct
	importScan := &defectdojo.ImportScan{
		MinimumSeverity: ptrString("info"),
		Active:          ptrBool(true),
		Verified:        ptrBool(true),
		File:            ptrString(filePath),
	}

	// Use the Create method to import scan results
	importedScan, err := dj.ImportScan.Create(cxt, importScan)
	if err != nil {
		// Print error details
		fmt.Printf("Failed to import scan results: %v\n", err)
		if httpErr, ok := err.(defectdojo.HTTPError); ok {
			fmt.Printf("HTTP Status Code: %d\n", httpErr.StatusCode)
			fmt.Printf("Response Body: %s\n", httpErr.Body)
		}
	} else {
		fmt.Printf("Scan results imported successfully: %+v\n", importedScan)
	}
}

// Helper functions to create pointers
func ptrString(s string) *string {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}
