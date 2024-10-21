package parser

import (
	_ "embed"
	"reddit-go-api-generator/scraper"
)

//go:embed sdk_helpers.txt
var sdkHelpers string

// // Generates Go functions from a list of endpoints
// func GenerateGoFunctions(endpoints []scraper.Endpoint) []string {
// 	var functions []string

// 	functions = append(functions, sdkHelpers)

// 	for _, endpoint := range endpoints {
// 		funcName := buildFunctionName(endpoint)
// 		responseStruct := generateResponseStruct(endpoint, funcName)
// 		comment := generateFunctionComment(endpoint, funcName)
// 		funcSignature := generateFunctionSignature(endpoint, funcName)
// 		urlBuild := buildURL(endpoint)
// 		payloadBuild := buildPayload(endpoint)
// 		queryParamsBuild := buildQueryParams(endpoint)
// 		requestBuild := buildRequest(endpoint, funcName)
// 		funcEnd := buildFunctionEnd(funcName)

// 		function := responseStruct + comment + funcSignature + urlBuild + payloadBuild + queryParamsBuild + requestBuild + funcEnd
// 		functions = append(functions, function)
// 	}

// 	return functions
// }

// Generates Go functions from a list of endpoints
func GenerateGoFunctions(endpoints []scraper.Endpoint) []string {
	var functions []string
	functions = append(functions, sdkHelpers)

	for _, endpoint := range endpoints {
		funcName := buildFunctionName(endpoint)
		enums := collectEnums(endpoint, funcName)
		enumDefs := generateEnumDefinitions(enums)
		responseStruct := generateResponseStruct(endpoint, funcName)
		comment := generateFunctionComment(endpoint, funcName)
		funcSignature := generateFunctionSignature(endpoint, funcName, enums)
		urlBuild := buildURL(endpoint)
		payloadBuild := buildPayload(endpoint)
		queryParamsBuild := buildQueryParams(endpoint)
		requestBuild := buildRequest(endpoint, funcName)
		funcEnd := buildFunctionEnd(funcName)

		function := enumDefs + responseStruct + comment + funcSignature + urlBuild + payloadBuild + queryParamsBuild + requestBuild + funcEnd
		functions = append(functions, function)
	}

	return functions
}
