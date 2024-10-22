package parser

import (
	"regexp"
)

func CleanColonPath(path string) string {
	re := regexp.MustCompile(`:([a-zA-Z0-9_]+)`)
	return re.ReplaceAllString(path, `{$1}`)
}
