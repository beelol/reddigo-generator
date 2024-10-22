package parser

import (
	"fmt"
	"log"
	"reddit-go-api-generator/models"
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
func buildFunctionName(endpoint models.Endpoint) string {
	method := strings.Title(strings.ToLower(endpoint.Method))
	name := cleanAPIPath(endpoint.Path)
	name = cleanPath(name)
	return fmt.Sprintf("%s%s", method, name)
}

// Helper function to safely remove /api/v1/ or /api/ from the path
func cleanAPIPath(path string) string {
	// Remove /api/v1/ if it exists
	path = strings.Replace(path, "/api/v1/", "/", 1)
	// Remove /api/ if it exists (only if /api/v1/ wasn't present)
	path = strings.Replace(path, "/api/", "/", 1)
	return path
}

// Helper function to generate the response struct if needed
func generateResponseStruct(endpoint models.Endpoint, funcName string) string {
	if len(endpoint.Response) == 0 {
		return ""
	}
	responseStructName := getResponseStructName(funcName, endpoint.Response)

	structDef := fmt.Sprintf("// %s represents the response for %s %s\n", responseStructName, endpoint.Method, endpoint.Path)
	structDef += fmt.Sprintf("type %s struct {\n", responseStructName)
	for _, resp := range endpoint.Response {
		fieldName := strings.Title(toCamelCaseFromSnakeCase(resp.Name))

		if fieldName == "Type" {
			fieldName = "TypeValue" // Avoid Go keyword conflict
		}

		fieldType := adjustEnumType(resp.Type)
		jsonTag := toSnakeCase(resp.Name)

		// Format the description as a multi-line comment if it contains multiple lines
		description := formatFieldDescription(resp.Description)

		structDef += fmt.Sprintf("\t%s %s `json:\"%s\"` %s\n", fieldName, fieldType, jsonTag, description)
	}
	structDef += "}\n\n"
	return structDef
}

func getResponseStructName(funcName string, response []models.Output) string {
	if len(response) == 0 {
		println(fmt.Sprintf("%s has no response body", funcName))
		return "any"
	}

	responseStructName := fmt.Sprintf("%sResponse", funcName)
	return responseStructName
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
func generateFunctionComment(endpoint models.Endpoint, funcName string) string {
	return fmt.Sprintf(`/*
%s makes a %s request to %s
ID: %s
Description: %s
*/
`, funcName, endpoint.Method, endpoint.Path, endpoint.ID, endpoint.Description)
}

func generateFunctionSignature(endpoint models.Endpoint, funcName string, enums []models.Enum) string {
	log.Printf("Generating function signature for: %s", funcName)

	params := collectFunctionParameters(endpoint)

	log.Printf("Parameters collected: %v", params)

	responseStructName := getResponseStructName(funcName, endpoint.Response)

	return fmt.Sprintf("func (sdk *ReddiGoSDK) %s(%s) (%s, error) {\n", funcName, strings.Join(params, ", "), responseStructName)
}

// Helper function to collect parameters for the function signature
func collectFunctionParameters(endpoint models.Endpoint) []string {
	var params []string
	paramSet := make(map[string]bool) // A set to track existing parameter names

	// Add dynamic parts (e.g., subreddit, where)
	dynamicFields := extractDynamicFields(endpoint.Path)
	for _, field := range dynamicFields {
		paramName := formatProperty(field)
		if !paramSet[paramName] { // Only add if it hasn't been added yet
			params = append(params, fmt.Sprintf("%s string", paramName))
			paramSet[paramName] = true
		}
	}

	// Add URL parameters, payload, and query parameters
	for _, param := range endpoint.URLParams {
		paramName := formatProperty(param)
		if !paramSet[paramName] {
			params = append(params, fmt.Sprintf("%s string", paramName))
			paramSet[paramName] = true
		}
	}
	for _, payload := range endpoint.Payload {
		paramName := formatProperty(payload.Name)
		paramType := adjustEnumType(payload.Type)
		if !paramSet[paramName] {
			params = append(params, fmt.Sprintf("%s %s", paramName, paramType))
			paramSet[paramName] = true
		}
	}
	for _, queryParam := range endpoint.QueryParams {
		paramName := formatProperty(queryParam.Name)
		if !paramSet[paramName] {
			params = append(params, fmt.Sprintf("%s string", paramName))
			paramSet[paramName] = true
		}
	}

	return params
}

// Helper function to build the URL using dynamic fields and parameters
func buildURL(endpoint models.Endpoint) string {
	urlPattern := transformDynamicFields(endpoint.Path)

	dynamicFields := extractDynamicFields(endpoint.Path)

	urlBuild := ""

	if len(dynamicFields) > 0 {
		urlBuild = fmt.Sprintf("\treqUrl := fmt.Sprintf(\"%s\"", urlPattern)

		for _, field := range dynamicFields {
			urlBuild += fmt.Sprintf(", %s", toLowerCamelCase(field))
		}
		urlBuild += ")"
	} else {
		urlBuild = fmt.Sprintf("\treqUrl := \"%s\"", urlPattern)
	}

	urlBuild += "\n"
	return urlBuild
}

func extractDynamicFields(path string) []string {
	var fields []string
	for {
		start := strings.Index(path, "{")
		if start == -1 {
			break
		}
		end := strings.Index(path[start:], "}")
		if end == -1 {
			break
		}
		field := path[start+1 : start+end]
		fields = append(fields, field)
		// Move 'path' forward beyond the closing brace
		path = path[start+end+1:]
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
func buildPayload(endpoint models.Endpoint) string {
	if len(endpoint.Payload) == 0 {
		return ""
	}

	// Check if the payload should be treated as the entire JSON body
	if len(endpoint.Payload) == 1 && strings.ToLower(endpoint.Payload[0].Name) == "json" {
		payloadName := toLowerCamelCase(endpoint.Payload[0].Name)

		return fmt.Sprintf("\tpayload := %s\n", payloadName)
	}

	payloadBuild := "\tpayload := map[string]interface{}{\n"
	for _, payload := range endpoint.Payload {
		//payloadBuild += fmt.Sprintf("\t\t\"%s\": %s,\n", toSnakeCase(payload.Name), toLowerCamelCase(payload.Name))

		paramName := formatProperty(payload.Name)
		jsonName := toSnakeCase(payload.Name)

		payloadBuild += fmt.Sprintf("\t\t\"%s\": %s,\n", jsonName, paramName)
	}

	payloadBuild += "\t}\n"
	return payloadBuild
}

// Helper function to build query parameters
func buildQueryParams(endpoint models.Endpoint) string {
	if len(endpoint.QueryParams) == 0 {
		return ""
	}
	queryParamsBuild := "\tqueryParams := urlpkg.Values{}\n"
	for _, queryParam := range endpoint.QueryParams {
		queryParamsBuild += fmt.Sprintf("\tqueryParams.Add(\"%s\", %s)\n", toSnakeCase(queryParam.Name), toLowerCamelCase(queryParam.Name))
	}
	queryParamsBuild += "\treqUrl += \"?\" + queryParams.Encode()\n"
	return queryParamsBuild
}

// Helper function to construct the request using MakeRequest from ReddiGoSDK
func buildRequest(endpoint models.Endpoint, funcName string) string {
	requestBuild := fmt.Sprintf("\t// Construct the request for %s method\n", endpoint.Method)

	// Build the URL using the helper function
	// requestBuild += buildURL(endpoint)

	responseName := getResponseStructName(funcName, endpoint.Response)
	newInstanceOfResponseStr := ""

	if responseName == "any" {
		newInstanceOfResponseStr = "nil"
	} else {
		newInstanceOfResponseStr = fmt.Sprintf("%s{}", responseName)
	}

	// Variable to hold the body for the MakeRequest function
	bodyVar := "nil"

	// Add headers and body for POST/PUT methods
	if endpoint.Method == "POST" || endpoint.Method == "PATCH" || endpoint.Method == "PUT" {
		requestBuild += "\tjsonPayload, err := jsonpkg.Marshal(payload)\n"
		requestBuild += "\tif err != nil {\n\t\treturn " + newInstanceOfResponseStr + ", err\n\t}\n"
		//requestBuild += "\treq.Header.Set(\"Content-Type\", \"application/json\")\n"
		//requestBuild += "\treq.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))\n"
		bodyVar = "bytes.NewBuffer(jsonPayload)"
	}

	requestBuild += fmt.Sprintf("\tresp, err := sdk.MakeRequest(\"%s\", reqUrl, %s)\n", endpoint.Method, bodyVar)
	requestBuild += "\tif err != nil {\n\t\treturn " + newInstanceOfResponseStr + ", err\n\t}\n"
	requestBuild += "\tdefer resp.Body.Close()\n"
	requestBuild += fmt.Sprintf("\tvar response %s\n", responseName)
	requestBuild += "\tif err := jsonpkg.NewDecoder(resp.Body).Decode(&response); err != nil {\n"
	requestBuild += "\t\treturn " + newInstanceOfResponseStr + ", err\n\t}\n"

	return requestBuild
}

// Helper function to close the function with a return statement
func buildFunctionEnd(funcName string) string {
	return fmt.Sprintf("\treturn response, nil\n}\n\n")
}
