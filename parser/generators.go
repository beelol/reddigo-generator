package parser

import (
	"fmt"
	"reddit-go-api-generator/scraper"
	"strings"
)

// Helper function to remove dynamic fields from URLs for function and type names
func removeDynamicFields(path string) string {
	// Replace placeholders like [subreddit] and {where}
	cleanPath := strings.ReplaceAll(path, "[", "")
	cleanPath = strings.ReplaceAll(cleanPath, "]", "")
	cleanPath = strings.ReplaceAll(cleanPath, "{", "")
	cleanPath = strings.ReplaceAll(cleanPath, "}", "")
	return cleanPath
}

// Helper function to create the function name in camel case, handling placeholders
func buildFunctionName(endpoint scraper.Endpoint) string {
	method := strings.Title(strings.ToLower(endpoint.Method))
	path := cleanPath(endpoint.Path)
	return fmt.Sprintf("%s%s", method, path)
}

// Helper function to generate the response struct if needed
func generateResponseStruct(endpoint scraper.Endpoint, funcName string) string {
	if len(endpoint.Response) == 0 {
		return ""
	}
	responseStructName := fmt.Sprintf("%sResponse", funcName)
	structDef := fmt.Sprintf("// %s represents the response for %s %s\n", responseStructName, endpoint.Method, endpoint.Path)
	structDef += fmt.Sprintf("type %s struct {\n", responseStructName)
	for _, resp := range endpoint.Response {
		fieldName := strings.Title(toCamelCaseFromUnderscore(resp.Name))
		fieldType := resp.Type
		jsonTag := toSnakeCase(resp.Name)

		// Format the description as a multi-line comment if it contains multiple lines
		description := formatFieldDescription(resp.Description)

		structDef += fmt.Sprintf("\t%s %s `json:\"%s\"` %s\n", fieldName, fieldType, jsonTag, description)
	}
	structDef += "}\n\n"
	return structDef
}

// Helper function to format the field description for multi-line comments
func formatFieldDescription(description string) string {
	lines := strings.Split(description, "\n")
	if len(lines) > 1 {
		// Format as a multi-line comment
		comment := "/* " + lines[0]
		for _, line := range lines[1:] {
			comment += "\n\t" + line
		}
		comment += " */"
		return comment
	}
	// Format as a single-line comment if there's only one line
	return fmt.Sprintf("// %s", description)
}

// Helper function to generate function comments
func generateFunctionComment(endpoint scraper.Endpoint, funcName string) string {
	return fmt.Sprintf(`/*
%s makes a %s request to %s
ID: %s
Description: %s
*/
`, funcName, endpoint.Method, endpoint.Path, endpoint.ID, endpoint.Description)
}

// Helper function to generate the function signature as a method of ReddigoSDK
// func generateFunctionSignature(endpoint scraper.Endpoint, funcName string) string {
// 	params := collectFunctionParameters(endpoint)
// 	return fmt.Sprintf("func (sdk *ReddigoSDK) %s(%s) (%sResponse, error) {\n", funcName, strings.Join(params, ", "), funcName)
// }

func generateFunctionSignature(endpoint scraper.Endpoint, funcName string, enums []Enum) string {
	params := collectFunctionParameters(endpoint, enums)
	return fmt.Sprintf("func (sdk *ReddigoSDK) %s(%s) (%sResponse, error) {\n", funcName, strings.Join(params, ", "), funcName)
}

// Helper function to collect parameters for the function signature
func collectFunctionParameters(endpoint scraper.Endpoint, enums []Enum) []string {
	var params []string
	// Add dynamic parts (e.g., subreddit, where)
	dynamicFields := extractDynamicFields(endpoint.Path)
	for _, field := range dynamicFields {
		params = append(params, fmt.Sprintf("%s string", toLowerCamelCase(field)))
	}
	// Add URL parameters, payload, and query parameters
	for _, param := range endpoint.URLParams {
		params = append(params, fmt.Sprintf("%s string", toLowerCamelCase(param)))
	}
	for _, payload := range endpoint.Payload {

		paramType := adjustEnumType(payload.Type)

		params = append(params, fmt.Sprintf("%s %s", toLowerCamelCase(payload.Name), paramType))
	}
	for _, queryParam := range endpoint.QueryParams {
		params = append(params, fmt.Sprintf("%s string", toLowerCamelCase(queryParam.Name)))
	}
	return params
}

// Helper function to build the URL using dynamic fields and parameters
func buildURL(endpoint scraper.Endpoint) string {
	urlPattern := transformDynamicFields(endpoint.Path)
	urlBuild := fmt.Sprintf("\turl := fmt.Sprintf(\"%s\"", urlPattern)
	for _, field := range extractDynamicFields(endpoint.Path) {
		urlBuild += fmt.Sprintf(", %s", toLowerCamelCase(field))
	}
	urlBuild += ")\n"
	return urlBuild
}

// Helper function to extract dynamic fields from a URL path
func extractDynamicFields(path string) []string {
	var fields []string
	// Extract only fields enclosed in curly braces
	start := strings.Index(path, "{")
	for start != -1 {
		end := strings.Index(path[start:], "}")
		if end != -1 {
			field := path[start+1 : start+end]
			fields = append(fields, field)
			start = strings.Index(path[start+end:], "{")
		} else {
			break
		}
	}
	return fields
}

// Helper function to transform dynamic fields into fmt.Sprintf placeholders
func transformDynamicFields(path string) string {
	cleanPath := path
	dynamicFields := extractDynamicFields(path)
	for _, field := range dynamicFields {
		cleanPath = strings.ReplaceAll(cleanPath, "{"+field+"}", "%s")
	}
	return cleanPath
}

// Helper function to prepare payload
func buildPayload(endpoint scraper.Endpoint) string {
	if len(endpoint.Payload) == 0 {
		return ""
	}

	payloadBuild := "\tpayload := map[string]interface{}{\n"
	for _, payload := range endpoint.Payload {
		payloadBuild += fmt.Sprintf("\t\t\"%s\": %s,\n", toSnakeCase(payload.Name), toLowerCamelCase(payload.Name))
	}

	payloadBuild += "\t}\n"
	return payloadBuild
}

// Helper function to build query parameters
func buildQueryParams(endpoint scraper.Endpoint) string {
	if len(endpoint.QueryParams) == 0 {
		return ""
	}
	queryParamsBuild := "\tqueryParams := url.Values{}\n"
	for _, queryParam := range endpoint.QueryParams {
		queryParamsBuild += fmt.Sprintf("\tqueryParams.Add(\"%s\", %s)\n", toSnakeCase(queryParam.Name), toLowerCamelCase(queryParam.Name))
	}
	queryParamsBuild += "\turl += \"?\" + queryParams.Encode()\n"
	return queryParamsBuild
}

// Helper function to construct the request using MakeRequest from ReddigoSDK
func buildRequest(endpoint scraper.Endpoint, funcName string) string {
	requestBuild := fmt.Sprintf("\t// Construct the request for %s method\n", endpoint.Method)

	// Build the URL using the helper function
	requestBuild += buildURL(endpoint)

	// Add headers and body for POST/PUT methods
	if endpoint.Method == "POST" || endpoint.Method == "PATCH" || endpoint.Method == "PUT" {
		requestBuild += "\tjsonPayload, err := json.Marshal(payload)\n"
		requestBuild += "\tif err != nil {\n\t\treturn " + funcName + "Response{}, err\n\t}\n"
		requestBuild += "\treq.Header.Set(\"Content-Type\", \"application/json\")\n"
		requestBuild += "\treq.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))\n"
	}

	requestBuild += "\treq, err := sdk.MakeRequest(\"" + endpoint.Method + "\", url, nil)\n"
	requestBuild += "\tif err != nil {\n\t\treturn " + funcName + "Response{}, err\n\t}\n"
	requestBuild += "\tdefer resp.Body.Close()\n"
	requestBuild += fmt.Sprintf("\tvar response %sResponse\n", funcName)
	requestBuild += "\tif err := json.NewDecoder(resp.Body).Decode(&response); err != nil {\n"
	requestBuild += "\t\treturn " + funcName + "Response{}, err\n\t}\n"

	return requestBuild
}

// Helper function to close the function with a return statement
func buildFunctionEnd(funcName string) string {
	return fmt.Sprintf("\treturn response, nil\n}\n\n")
}
