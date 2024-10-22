package parser

import (
	"io/ioutil"
	"log"
	"reddit-go-api-generator/models"
	"strings"
	"testing"
)

// TestGenerateGoFunctions verifies that the generated functions match the expected output.
func TestGenerateGoFunctions(t *testing.T) {
	// Define dummy endpoints
	endpoints := []models.Endpoint{
		{
			ID:          "1",
			Method:      "GET",
			Path:        "/api/v1/me",
			Description: "Returns the identity of the user.",
			URLParams:   []string{},
			Payload:     []models.Input{},
			Response: []models.Output{
				{Name: "username", Type: "string", Description: "The username of the user"},
			},
			QueryParams: []models.Parameter{},
		},
		{
			ID:          "2",
			Method:      "POST",
			Path:        "/api/v1/collections/create_collection",
			Description: "Create a collection",
			URLParams:   []string{},
			Payload: []models.Input{
				{Name: "title", Type: "string", Description: "Title of the collection"},
				{Name: "description", Type: "string", Description: "Description of the collection"},
			},
			Response: []models.Output{
				{Name: "collection_id", Type: "string", Description: "The ID of the created collection"},
			},
			QueryParams: []models.Parameter{
				{Name: "lang", Type: "string", Description: "Language of the collection"},
			},
		},
	}

	// Generate functions
	functions := GenerateGoFunctions(endpoints)

	// Combine the functions into one output string
	generatedOutput := strings.Join(functions, "\n")

	// Expected output
	content, err := ioutil.ReadFile("./preview.txt")
	if err != nil {

		log.Fatal(err)
	}

	// Convert content to string
	expectedOutput := string(content)

	// Compare the generated output with the expected output
	if strings.TrimSpace(generatedOutput) != strings.TrimSpace(expectedOutput) {
		//t.Errorf("Generated output does not match expected output.\nExpected:\n%s\n\nGot:\n%s", expectedOutput, generatedOutput)
		t.Errorf("Generated output does not match expected output.")
	}
}
