package parser

import "strings"

// replacePeriodsWithDot replaces all periods in the path with the word "dot".
func replacePeriodsWithDot(path string) string {
	return strings.ReplaceAll(path, ".", "Dot")
}
func cleanPath(path string) string {
	return ensureCamelCase(replacePeriodsWithDot(removeDynamicFields(path)))
}

func splitAndTakeFirst(input string) string {
	// Check if there is a slash in the string
	if idx := strings.Index(input, "/"); idx != -1 {
		// Return the substring before the slash
		return strings.TrimSpace(input[:idx])
	}
	// Return the original string if no slash is present
	return strings.TrimSpace(input)
}

func RemoveInvalidCharacters(property string) string {
	// Fixes this weird one-off case: ('user',)
	newStr := splitAndTakeFirst(property)
	newStr = strings.ReplaceAll(newStr, "(", "")
	newStr = strings.ReplaceAll(newStr, ")", "")
	newStr = strings.ReplaceAll(newStr, "'", "")
	newStr = strings.ReplaceAll(newStr, ",", "")
	return newStr
}

func formatProperty(property string) string {
	return ensureNonKeyword(toLowerCamelCase(RemoveInvalidCharacters(property)))
}

func toCamelCaseFromSlash(str string) string {
	parts := strings.Split(str, "/")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func toCamelCaseFromSnakeCase(str string) string {
	parts := strings.Split(str, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func toCamelCaseFromKebabCase(str string) string {
	parts := strings.Split(str, "-")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	return strings.Join(parts, "")
}

func ensureCamelCase(str string) string {
	return toCamelCaseFromKebabCase(toCamelCaseFromSnakeCase(toCamelCaseFromSlash(str)))
}

// Converts a string to lowerCamelCase
func toLowerCamelCase(str string) string {
	camel := ensureCamelCase(str)
	return strings.ToLower(string(camel[0])) + camel[1:]
}

func ensureNonKeyword(str string) string {
	if str == "type" {
		str = "typeValue"
	}

	return str
}

// Converts a string to snake_case
func toSnakeCase(str string) string {
	return strings.ReplaceAll(strings.ToLower(RemoveInvalidCharacters(str)), " ", "_")
}
