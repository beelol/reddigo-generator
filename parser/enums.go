package parser

import (
	"fmt"
	"reddit-go-api-generator/models"
	"strings"
)

// Collect enums from Payload, Response, and QueryParams
func collectEnums(endpoint models.Endpoint, funcName string) []models.Enum {
	var enums []models.Enum

	// Collect enums from Payload
	for _, payload := range endpoint.Payload {
		if strings.HasPrefix(payload.Type, "enum(") {
			enumValues := extractEnumValues(payload.Type)
			enumName := fmt.Sprintf("%s%sEnum", funcName, toCamelCaseFromSnakeCase(payload.Name))
			enums = append(enums, models.Enum{Name: enumName, Values: enumValues})
		}
	}

	// Collect enums from Response
	for _, resp := range endpoint.Response {
		if strings.HasPrefix(resp.Type, "enum(") {
			enumValues := extractEnumValues(resp.Type)
			enumName := fmt.Sprintf("%s%sEnum", funcName, toCamelCaseFromSnakeCase(resp.Name))
			enums = append(enums, models.Enum{Name: enumName, Values: enumValues})
		}
	}

	// Collect enums from QueryParams
	for _, param := range endpoint.QueryParams {
		if strings.HasPrefix(param.Type, "enum(") {
			enumValues := extractEnumValues(param.Type)
			enumName := fmt.Sprintf("%s%sEnum", funcName, toCamelCaseFromSnakeCase(param.Name))
			enums = append(enums, models.Enum{Name: enumName, Values: enumValues})
		}
	}

	return enums
}

// Extract values from an enum type string
func extractEnumValues(enumStr string) []string {
	values := strings.TrimPrefix(strings.TrimSuffix(enumStr, ")"), "enum(")
	return strings.Split(values, ", ")
}

// Generate Go definitions for enums
func generateEnumDefinitions(enums []models.Enum) string {
	var enumDefs string

	for _, enum := range enums {
		enumDefs += fmt.Sprintf("type %s string\n\n", enum.Name)
		enumDefs += "const (\n"
		for _, value := range enum.Values {
			//newIdentifierName := strings.ReplaceAll(enum.Name, "-", "Minus")

			identifier := strings.ReplaceAll(fmt.Sprintf("%s%s", enum.Name, strings.Title(value)), "-", "Minus")
			enumDefs += fmt.Sprintf("\t%s %s = \"%s\"\n", identifier, enum.Name, value)
		}
		enumDefs += ")\n\n"
	}

	return enumDefs
}

// Adjust the type to string for function parameters if it's an enum
func adjustEnumType(typeStr string) string {
	if strings.HasPrefix(typeStr, "enum(") {
		return "string"
	}
	return typeStr
}

// Helper function to build the URL using dynamic fields and parameters

// The rest of the helper functions remain similar to the previous version
