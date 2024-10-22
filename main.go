package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reddit-go-api-generator/parser"
	"reddit-go-api-generator/scraper"
)

// "https://www.reddit.com/dev/api/"

// CountEndpointsOnRedditAPI counts the number of endpoints on the Reddit API documentation page
//func CountEndpointsOnRedditAPI() (int, error) {
//	count := 0
//	c := colly.NewCollector()
//
//	// Count each endpoint section
//	c.OnHTML("div.endpoint", func(e *colly.HTMLElement) {
//		count++
//		// log.Printf("Endpoint found: %s %s", e.ChildText("h3 span.method"), e.ChildText("h3 em.placeholder"))
//	})
//
//	// Visit the API documentation page
//	err := c.Visit(scraper.RedditAPIUrl)
//	if err != nil {
//		log.Fatalf("Error visiting URL: %v", err)
//		return 0, err
//	}
//
//	return count, nil
//}

// setupSDKDirectory creates the directory structure for the SDK if it does not exist
func setupSDKDirectory(basePath string) (string, error) {
	// Define the SDK directory path
	sdkDir := filepath.Join(basePath)

	// Create the SDK directory if it does not exist
	if err := os.MkdirAll(sdkDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create SDK directory structure: %w", err)
	}

	// Return the final path where the generated file should be placed
	return filepath.Join(sdkDir, "reddigo.go"), nil
}

func main() {
	endpointsData, err := scraper.ScrapeRedditAPI(0,
		func(s string) {
			//log.Printf("Targeted: %s", s)
		},
		func(s string) {
			//log.Printf("Processed: %s", s)
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Successfully scraped %d endpointsData", len(endpointsData))
	}

	finalFunctions := parser.GenerateGoFunctions(endpointsData)

	for _, function := range finalFunctions {
		println(function)
	}

	sdkPath := collectSDKFilePath()
	outputFileFinal, err := setupSDKDirectory(sdkPath)

	if err != nil {
		log.Fatalf("Error setting up SDK directory: %v", err)
	}

	err = writeToFile(outputFileFinal, finalFunctions)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Iterate through the collected endpoints and print them
	for _, endpoint := range endpointsData {
		fmt.Printf("Endpoint: %+v\n", endpoint)
	}

	log.Println("Successfully wrote generated functions to reddit_api_sdk.go")

	os.Exit(0)
}

// writeToFile writes the generated Go functions to a file
func writeToFile(filename string, functions []string) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	// Write each function to the file
	for _, function := range functions {
		_, err := file.WriteString(function + "\n\n")
		if err != nil {
			return fmt.Errorf("could not write to file: %w", err)
		}
	}

	return nil
}

func collectSDKFilePath() string {
	// Define a flag to collect the SDK base path
	sdkPath := flag.String("o", "reddigo", "Specify the base path for the SDK directory")

	// Parse the flags
	flag.Parse()

	// Check if the -sdkPath flag is provided and has a value
	if *sdkPath == "" {
		log.Fatal("Error: You must provide an SDK base path with -sdkPath")
	}

	//// Construct the SDK file path
	//filePath, err := setupSDKDirectory(*sdkPath)
	//if err != nil {
	//	log.Fatalf("Error setting up SDK directory: %v", err)
	//}
	//
	//// Print the path for confirmation
	//fmt.Printf("SDK file path: %s\n", filePath)

	//return &filePath

	return *sdkPath
}
