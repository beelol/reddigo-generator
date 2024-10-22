package parser

import (
	_ "embed"
	"fmt"
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

// GenerateGoFunctions Generates Go functions from a list of endpoints
func GenerateGoFunctions(endpoints []scraper.Endpoint) []string {
	var functions []string
	functions = append(functions, sdkHelpers)

	for _, endpoint := range endpoints {
		println(fmt.Sprintf("operating on %s", endpoint))
		println("reached buildFunctionName")

		funcName := buildFunctionName(endpoint)
		println("reached collectEnums")

		enums := collectEnums(endpoint, funcName)
		println("reached generateEnumDefinitions")

		enumDefs := generateEnumDefinitions(enums)
		println("reached generateResponseStruct")

		responseStruct := generateResponseStruct(endpoint, funcName)
		println("reached generateFunctionComment")

		comment := generateFunctionComment(endpoint, funcName)
		println("reached generateFunctionSignature")

		funcSignature := generateFunctionSignature(endpoint, funcName, enums)
		println("reached buildURL")

		urlBuild := buildURL(endpoint)
		println("reached buildPayload")

		payloadBuild := buildPayload(endpoint)
		println("reached buildQueryParams")

		queryParamsBuild := buildQueryParams(endpoint)
		println("reached buildRequest")

		requestBuild := buildRequest(endpoint, funcName)
		println("reached buildFunctionEnd")

		funcEnd := buildFunctionEnd(funcName)
		println("reached end")

		function := enumDefs + responseStruct + comment + funcSignature + urlBuild + payloadBuild + queryParamsBuild + requestBuild + funcEnd
		functions = append(functions, function)

		println(fmt.Sprintf("completed generating go func for %s", endpoint))
	}

	return functions
}
