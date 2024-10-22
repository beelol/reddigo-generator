package models

type RedditAPIField struct {
	Name        string
	Description string
	Type        string
}

type Endpoint struct {
	ID          string
	Method      string
	Path        string
	Description string
	URLParams   []string
	Payload     []Input
	Response    []Output
	QueryParams []Parameter
}

type Input struct {
	Name        string
	Description string
	Type        string
}

type Output struct {
	Name        string
	Description string
	Type        string
}

type Parameter struct {
	Name        string
	Description string
	Type        string
}

// Struct to represent enums
type Enum struct {
	Name   string
	Values []string
}
