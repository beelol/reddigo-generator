package generator

import (
	"fmt"
	"os"

	"reddit-go-api-generator/parser"
)

func GenerateFunction(schema parser.Schema) error {
	functionStr := fmt.Sprintf(`func %s() {
        // Endpoint: %s %s
        // Parameters: %+v
    }`, schema.ID, schema.Method, schema.Path, schema.Params)

	file, err := os.OpenFile("reddit_sdk.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(functionStr); err != nil {
		return err
	}

	return nil
}
