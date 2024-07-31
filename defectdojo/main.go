package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/blackaichi/go-defectdojo/defectdojo"
)

func main() {
	url := "http://13.201.28.37:8080"
	authorizedToken := "798b581a65f991dc728d2ae8745a77f6459078bc"

	// Create DefectDojo client
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

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("The error is %v\n", err)
		return
	}

	// Define the relative file path
	filePath := "result/trivy_output.json"

	// Join the current working directory with the file path
	fullPath := filepath.Join(cwd, filePath)

	// Check if the file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %v", err)
	}

	// Create ImportScan struct
	importScan := &defectdojo.ImportScan{
		MinimumSeverity:   ptrString("Info"),
		Active:            ptrBool(true),
		Verified:          ptrBool(true),
		File:              ptrString(filePath),
		EngagementId:      ptrInt(17),
		ScanType:          ptrString("Trivy Scan"),
		Environment:       ptrString("Development"),
		AutoCreateContext: ptrBool(true),
		ProductName:       ptrString("myproject"),
		ProductTypeName:   ptrString("Research and Development"),
	}

	// Debugging: print importScan details
	fmt.Printf("ImportScan details: %+v\n", importScan)

	// Use the Create method to import scan results
	importedScan, err := dj.ImportScan.Create(cxt, importScan)
	if err != nil {
		// Print error details
		fmt.Printf("Failed to import scan results: %v\n", err)
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

func ptrInt(i int) *int {
	return &i
}
