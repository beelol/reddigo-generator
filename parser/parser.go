package parser

import (
	"fmt"
	"reddit-go-api-generator/scraper"
	"strings"
)

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

// Generates Go functions from a list of endpoints
func GenerateGoFunctions(endpoints []scraper.Endpoint) []string {
	var functions []string

	for _, endpoint := range endpoints {
		// Create function name
		funcName := fmt.Sprintf("%s%s", endpoint.Method, ensureCamelCase(endpoint.Path))

		// Generate the response struct if needed
		responseStruct := ""
		if len(endpoint.Response) > 0 {
			responseStructName := fmt.Sprintf("%sResponse", funcName)
			responseStruct = fmt.Sprintf("// %s represents the response for %s %s\n", responseStructName, endpoint.Method, endpoint.Path)
			responseStruct += fmt.Sprintf("type %s struct {\n", responseStructName)
			for _, resp := range endpoint.Response {
				fieldName := strings.Title(toCamelCaseFromUnderscore(resp.Name))
				fieldType := resp.Type
				jsonTag := toSnakeCase(resp.Name)
				responseStruct += fmt.Sprintf("\t%s %s `json:\"%s\"` // %s\n", fieldName, fieldType, jsonTag, resp.Description)
			}
			responseStruct += "}\n\n"
		}

		// Generate function comments as multiline
		comment := fmt.Sprintf(`/*
%s makes a %s request to %s
ID: %s
Description: %s
*/
`, funcName, endpoint.Method, endpoint.Path, endpoint.ID, endpoint.Description)

		// Generate function signature
		funcSignature := fmt.Sprintf("func %s(", funcName)

		// Collect parameters for the function signature
		params := []string{}
		for _, param := range endpoint.URLParams {
			params = append(params, fmt.Sprintf("%s string", toLowerCamelCase(param)))
		}
		for _, payload := range endpoint.Payload {
			params = append(params, fmt.Sprintf("%s %s", toLowerCamelCase(payload.Name), payload.Type))
		}
		for _, queryParam := range endpoint.QueryParams {
			params = append(params, fmt.Sprintf("%s string", toLowerCamelCase(queryParam.Name)))
		}

		// Add parameters to function signature
		funcSignature += strings.Join(params, ", ")
		funcSignature += fmt.Sprintf(") (%sResponse, error) {\n", funcName)

		// Build the URL using URL parameters
		urlBuild := fmt.Sprintf("\turl := fmt.Sprintf(\"%s\"", endpoint.Path)
		if len(endpoint.URLParams) > 0 {
			for _, param := range endpoint.URLParams {
				urlBuild += fmt.Sprintf(", %s", toLowerCamelCase(param))
			}
		}
		urlBuild += ")\n"

		// Prepare payload and query parameters
		payloadBuild := ""
		if len(endpoint.Payload) > 0 {
			payloadBuild += "\tpayload := map[string]interface{}{\n"
			for _, payload := range endpoint.Payload {
				payloadBuild += fmt.Sprintf("\t\t\"%s\": %s,\n", toSnakeCase(payload.Name), toLowerCamelCase(payload.Name))
			}
			payloadBuild += "\t}\n"
		}

		queryParamsBuild := ""
		if len(endpoint.QueryParams) > 0 {
			queryParamsBuild += "\tqueryParams := url.Values{}\n"
			for _, queryParam := range endpoint.QueryParams {
				queryParamsBuild += fmt.Sprintf("\tqueryParams.Add(\"%s\", %s)\n", toSnakeCase(queryParam.Name), toLowerCamelCase(queryParam.Name))
			}
			queryParamsBuild += "\turl += \"?\" + queryParams.Encode()\n"
		}

		// Construct the request based on the method
		requestBuild := fmt.Sprintf("\t// Construct the request for %s method\n", endpoint.Method)
		requestBuild += "\tclient := &http.Client{}\n"
		requestBuild += fmt.Sprintf("\treq, err := http.NewRequest(\"%s\", url, nil)\n", endpoint.Method)
		requestBuild += "\tif err != nil {\n\t\treturn " + funcName + "Response{}, err\n\t}\n"

		// Add headers and body for POST/PUT methods
		if endpoint.Method == "POST" || endpoint.Method == "PATCH" || endpoint.Method == "PUT" {
			requestBuild += "\tjsonPayload, err := json.Marshal(payload)\n"
			requestBuild += "\tif err != nil {\n\t\treturn " + funcName + "Response{}, err\n\t}\n"
			requestBuild += "\treq.Header.Set(\"Content-Type\", \"application/json\")\n"
			requestBuild += "\treq.Body = io.NopCloser(bytes.NewBuffer(jsonPayload))\n"
		}

		// Execute the request
		requestBuild += "\tresp, err := client.Do(req)\n"
		requestBuild += "\tif err != nil {\n\t\treturn " + funcName + "Response{}, err\n\t}\n"
		requestBuild += "\tdefer resp.Body.Close()\n"

		// Decode the response
		requestBuild += fmt.Sprintf("\tvar response %sResponse\n", funcName)
		requestBuild += "\tif err := json.NewDecoder(resp.Body).Decode(&response); err != nil {\n"
		requestBuild += "\t\treturn " + funcName + "Response{}, err\n\t}\n"

		// Function closing with return statement
		funcEnd := fmt.Sprintf("\treturn response, nil\n}\n\n")

		// Compile everything together
		function := responseStruct + comment + funcSignature + urlBuild + payloadBuild + queryParamsBuild + requestBuild + funcEnd
		functions = append(functions, function)
	}

	return functions
}
