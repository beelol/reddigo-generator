package parser

import "strings"

func toCamelCaseFromSlash(str string) string {
	parts := strings.Split(str, "/")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func toCamelCaseFromUnderscore(str string) string {
	parts := strings.Split(str, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func ensureCamelCase(str string) string {
	return toCamelCaseFromUnderscore(toCamelCaseFromSlash(str))
}

// Converts a string to lowerCamelCase
func toLowerCamelCase(str string) string {
	camel := toCamelCaseFromUnderscore(str)
	return strings.ToLower(string(camel[0])) + camel[1:]
}

// Converts a string to snake_case
func toSnakeCase(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "_")
}
