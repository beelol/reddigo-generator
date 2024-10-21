package parser

func cleanPath(path string) string {
	return ensureCamelCase(removeDynamicFields(path))
}
