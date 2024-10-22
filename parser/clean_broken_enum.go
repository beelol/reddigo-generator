package parser

import (
	"fmt"
	"strings"
)

func FormatOneOfEnum(description string) string {
	// Find the portion of the description starting with "one of ("
	startIndex := strings.Index(description, "one of (")
	if startIndex == -1 {
		// Return the description as is if the format doesn't match
		return description
	}

	// Extract everything after "one of ("
	substring := description[startIndex+8:]

	// Find the closing parenthesis
	endIndex := strings.Index(substring, ")")
	if endIndex == -1 {
		// Return the description as is if no closing parenthesis is found
		return description
	}

	// Extract the content within the parentheses
	oneOfContent := substring[:endIndex]

	// Check if the oneOfContent starts with backtick and comma
	if strings.HasPrefix(oneOfContent, "`,") {
		// Remove leading backtick and comma
		oneOfContent = oneOfContent[2:]

		// Remove trailing backticks if they exist
		oneOfContent = strings.TrimSuffix(oneOfContent, "`")

		// Replace commas with ", " for proper formatting
		oneOfContent = strings.ReplaceAll(oneOfContent, ",", ", ")

		// Reconstruct the description with formatted oneOf content
		return fmt.Sprintf("%sone of (%s)%s", description[:startIndex], oneOfContent, description[startIndex+endIndex+9:])
	}

	// If it doesn't match the specific pattern, return the description as is
	return description
}
