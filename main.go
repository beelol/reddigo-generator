package main

import (
	"log"

	"reddit-go-api-generator/parser"
	"reddit-go-api-generator/scraper"
)

func main() {
	endpointsData, err := scraper.ScrapeRedditAPI()

	if err != nil {
		log.Fatalf("Error scraping API: %v", err)
	}

	// for _, endpoint := range endpointsData {
	// 	log.Printf("Endpoint: %s, Method: %s, Path: %s, Description: %s, URL Params: %v, Payload: %v, Response: %v, Query Params: %v",
	// 		endpoint.ID, endpoint.Method, endpoint.Path, endpoint.Description, endpoint.URLParams, endpoint.Payload, endpoint.Response, endpoint.QueryParams)
	// }

	finalFunctions := parser.GenerateGoFunctions(endpointsData)
	for _, function := range finalFunctions {
		log.Print(function)
	}

	return

	// Load progress from file or initialize
	// store.LoadProgress()

	// Initialize Bubbletea TUI
	// p := progress.NewProgram(store.GetEndpoints())

	// // Start TUI and listen for progress updates
	// if err := p.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	// // Scrape Reddit API Documentation
	// endpoints, err := scraper.ScrapeRedditAPI()
	// if err != nil {
	// 	log.Fatalf("Error scraping API: %v", err)
	// }

	// // Parse and process each endpoint
	// for _, endpoint := range endpoints {
	// 	schema, err := parser.ParseEndpoint(endpoint)
	// 	if err != nil {
	// 		log.Printf("Error parsing endpoint %s: %v", endpoint.ID, err)
	// 		continue
	// 	}

	// 	// Generate Go function for the endpoint
	// 	if err := generator.GenerateFunction(schema); err != nil {
	// 		log.Printf("Error generating function for %s: %v", endpoint.ID, err)
	// 		continue
	// 	}

	// 	// Update progress and save state
	// 	store.MarkCompleted(endpoint.ID)
	// 	store.SaveProgress()
	// }

	// log.Println("All endpoints processed successfully!")
}
