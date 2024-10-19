package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Output struct {
	Name        string
	Description string
	Type        string
}

type Parameter struct {
	Name        string
	Description string
	Type        string
}

type Input struct {
	Name        string
	Description string
	Type        string
}

type Endpoint struct {
	ID          string
	Method      string
	Path        string
	Description string
	URLParams   []string
	Payload     []Input
	Response    []Output
	QueryParams []Parameter
}

// ScrapeRedditAPI scrapes the Reddit API documentation
func ScrapeRedditAPI(limit int) ([]Endpoint, error) {
	var endpoints []Endpoint
	c := colly.NewCollector()
	count := 0

	c.OnHTML("div.endpoint", func(e *colly.HTMLElement) {
		if count >= limit {
			return // Stop processing if we've reached the limit
		}

		method := e.ChildText("h3 span.method")

		path := extractCleanPath(e)

		id := method + " " + path

		description := e.ChildText("div.md p")
		if description == "" {
			description = "No description available"
		}

		urlParams := extractURLParams(e)
		payload := extractPayload(e)
		response := extractResponse(e)
		queryParams := extractQueryParams(e)

		endpoint := Endpoint{
			ID:          id,
			Method:      method,
			Path:        path,
			Description: description,
			URLParams:   urlParams,
			Payload:     payload,
			Response:    response,
			QueryParams: queryParams,
		}

		log.Printf("Processed endpoint: %s %s", method, path)
		endpoints = append(endpoints, endpoint)
		count++
	})

	err := c.Visit("https://www.reddit.com/dev/api/")

	if err != nil {
		log.Fatalf("Error visiting URL: %v", err)
		return nil, err
	}

	return endpoints, nil
}

// Extract URL parameters from placeholders in the path
func extractURLParams(e *colly.HTMLElement) []string {
	var urlParams []string
	e.ForEach("h3 em.placeholder", func(_ int, em *colly.HTMLElement) {
		urlParams = append(urlParams, em.Text)
	})
	return urlParams
}

// Extract payload parameters for POST/PATCH requests
func extractPayload(e *colly.HTMLElement) []Input {
	var inputs []Input
	e.ForEach("table.parameters tr.json-model", func(_ int, tr *colly.HTMLElement) {
		payload := tr.ChildText("td pre code")
		if payload != "" {
			lines := strings.Split(payload, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "{" || line == "}" || line == "" {
					continue
				}

				name, inputType, description := parsePayloadLine(line)
				input := Input{
					Name:        name,
					Description: description,
					Type:        inputType,
				}
				inputs = append(inputs, input)
			}
		}
	})
	return inputs
}

// Parses a line of payload and returns the name, type, and description
func parsePayloadLine(line string) (name, inputType, description string) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return "", "", ""
	}

	name = strings.TrimSpace(strings.Trim(parts[0], `"`))
	description = strings.TrimSpace(parts[1])

	inputType = determineType(description)

	return name, inputType, description
}

// Extract response parameters from the table
func extractResponse(e *colly.HTMLElement) []Output {
	var response []Output
	e.ForEach("table.parameters tbody tr", func(_ int, tr *colly.HTMLElement) {
		paramName := tr.ChildText("th")
		paramDesc := tr.ChildText("td p")
		paramType := determineType(paramDesc)

		response = append(response, Output{
			Name:        paramName,
			Description: paramDesc,
			Type:        paramType,
		})
	})
	return response
}

// Extract query parameters if present
func extractQueryParams(e *colly.HTMLElement) []Parameter {
	var queryParams []Parameter
	e.ForEach("table.parameters tbody tr", func(_ int, tr *colly.HTMLElement) {
		if tr.ChildText("th") == "after" || tr.ChildText("th") == "before" || tr.ChildText("th") == "count" || tr.ChildText("th") == "limit" {
			paramName := tr.ChildText("th")
			paramDesc := tr.ChildText("td p")
			paramType := determineType(paramDesc)

			queryParams = append(queryParams, Parameter{
				Name:        paramName,
				Description: paramDesc,
				Type:        paramType,
			})
		}
	})
	return queryParams
}

// Determine the type of a parameter based on its description
func determineType(description string) string {
	originalDesc := description
	description = strings.ToLower(description)

	switch {
	case strings.Contains(description, "boolean"):
		return "bool"
	case strings.Contains(description, "integer"):
		return "int"
	case strings.Contains(description, "string"):
		return "string"
	case strings.Contains(description, "valid url"):
		return "string"
	case strings.Contains(description, "a valid email"):
		return "string"
	case strings.Contains(description, "fullname"):
		return "string"
	case strings.Contains(description, "one of"):
		// Extract the content between parentheses as the enum values
		start := strings.Index(originalDesc, "one of (")
		end := strings.Index(originalDesc, ")")
		if start != -1 && end != -1 && end > start {
			options := originalDesc[start+8 : end]
			// Return formatted enum type with options
			return fmt.Sprintf("enum(%s)", options)
		}
	case strings.Contains(description, "expand"):
		return "bool"
	case strings.Contains(description, "alphanumeric"):
		return "string"
	case strings.Contains(description, "characters"):
		return "string"
	// Retain the original cases that may not have been included before
	case strings.Contains(description, "comma-separated"):
		return "string"
	case strings.Contains(description, "optional") && strings.Contains(description, "boolean"):
		return "bool"
	case strings.Contains(description, "optional") && strings.Contains(description, "integer"):
		return "int"
	}

	return "interface{}"
}

// Extract path from the h3 element, excluding oauth-scope-list and other elements
func extractCleanPath(e *colly.HTMLElement) string {
	h3 := e.DOM.Find("h3")

	// Remove any oauth-scope-list and api-badge elements from the h3
	h3.Find("span.oauth-scope-list").Remove()
	h3.Find("a").Remove()

	cleanPath := strings.TrimSpace(h3.Text())
	cleanPath = strings.TrimSpace(strings.Replace(cleanPath, e.ChildText("h3 span.method"), "", 1))

	return cleanPath
}
