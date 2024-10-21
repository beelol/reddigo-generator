package main

import (
	"fmt"
	"log"
	"os"

	"reddit-go-api-generator/parser"
	"reddit-go-api-generator/scraper"

	"github.com/gocolly/colly"
)

// CountEndpointsOnRedditAPI counts the number of endpoints on the Reddit API documentation page
func CountEndpointsOnRedditAPI() (int, error) {
	count := 0
	c := colly.NewCollector()

	// Count each endpoint section
	c.OnHTML("div.endpoint", func(e *colly.HTMLElement) {
		count++
		// log.Printf("Endpoint found: %s %s", e.ChildText("h3 span.method"), e.ChildText("h3 em.placeholder"))
	})

	// Visit the API documentation page
	err := c.Visit("https://www.reddit.com/dev/api/")
	if err != nil {
		log.Fatalf("Error visiting URL: %v", err)
		return 0, err
	}

	return count, nil
}

func main() {

	// Initialize Bubbletea TUI
	// count, err := CountEndpointsOnRedditAPI()
	// println(count)

	// if err != nil {
	// 	log.Fatal(err)

	// 	return
	// }

	endpointsData, err := scraper.ScrapeRedditAPI(11, func(endpoint string) {
		// p.Send(progress.SetCurrentEndpoint(endpoint))
	}, func(endpoint string) {
		// p.Send(progress.IncrementProgress())
	})

	if err != nil {
		log.Fatalf("Error scraping API: %v", err)
	}

	finalFunctions := parser.GenerateGoFunctions(endpointsData)

	err = writeToFile("reddit_api_sdk.go", finalFunctions)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Println("Successfully wrote generated functions to reddit_api_sdk.go")

	return
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
