package main

import (
	"log"
	"os"

	"reddit-go-api-generator/progress"
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
	count, err := CountEndpointsOnRedditAPI()
	println(count)

	if err != nil {
		log.Fatal(err)
		return
	}

	p := progress.NewProgram(count)

	go func() {

		_, err := scraper.ScrapeRedditAPI(0, func(endpoint string) {
			// p.Send(progress.SetCurrentEndpoint(endpoint))
		}, func(endpoint string) {
			p.Send(progress.IncrementProgress())
		})

		if err != nil {
			log.Fatalf("Error scraping API: %v", err)
		}

		// finalFunctions := parser.GenerateGoFunctions(endpointsData)

		// p.Send(progress.HideProgressBar())
		// log.Print(finalFunctions[0])
		// p.Send(progress.Message{Type: "done"})
		os.Exit(1)
	}()

	if err := p.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Start TUI and listen for progress updates
	// if err := p.Start(); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, function := range finalFunctions {
	// 	log.Print(function)
	// }

	return

	// Load progress from file or initialize
	// store.LoadProgress()

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
