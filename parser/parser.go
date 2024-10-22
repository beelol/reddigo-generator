package parser

import (
	_ "embed"
	"reddit-go-api-generator/models"
)

//go:embed sdk_helpers.txt
var sdkHelpers string

// GenerateGoFunctions Generates Go functions from a list of endpoints
func GenerateGoFunctions(endpoints []models.Endpoint) []string {
	var functions []string
	functions = append(functions, sdkHelpers)

	for _, endpoint := range endpoints {
		//println(fmt.Sprintf("operating on %s", endpoint))
		//println("reached buildFunctionName")

		funcName := buildFunctionName(endpoint)
		//println("reached collectEnums")

		enums := collectEnums(endpoint, funcName)
		//println("reached generateEnumDefinitions")

		enumDefs := generateEnumDefinitions(enums)
		//println("reached generateResponseStruct")

		responseStruct := generateResponseStruct(endpoint, funcName)
		//println("reached generateFunctionComment")

		comment := generateFunctionComment(endpoint, funcName)
		//println("reached generateFunctionSignature")

		funcSignature := generateFunctionSignature(endpoint, funcName, enums)
		//println("reached buildURL")

		urlBuild := buildURL(endpoint)
		//println("reached buildPayload")

		payloadBuild := buildPayload(endpoint)
		//println("reached buildQueryParams")

		queryParamsBuild := buildQueryParams(endpoint)
		//println("reached buildRequest")

		requestBuild := buildRequest(endpoint, funcName)
		//println("reached buildFunctionEnd")

		funcEnd := buildFunctionEnd(funcName)
		//println("reached end")

		function := enumDefs + responseStruct + comment + funcSignature + urlBuild + payloadBuild + queryParamsBuild + requestBuild + funcEnd
		functions = append(functions, function)

		//println(fmt.Sprintf("completed generating go func for %s", endpoint))
	}

	return functions
}
