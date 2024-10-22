package parser

import (
	"testing"
)

func TestCleanColonPath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/api/mod/conversations/:conversation_id", "/api/mod/conversations/{conversation_id}"},
		{"/api/user/:user_id/profile/:profile_id", "/api/user/{user_id}/profile/{profile_id}"},
		{"/api/item/:item_id/details", "/api/item/{item_id}/details"},
		{"/api/:version/status/:status_code", "/api/{version}/status/{status_code}"},
		{"/api/no_params", "/api/no_params"},
	}

	for _, test := range tests {
		output := CleanColonPath(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected '%s' but got '%s'", test.input, test.expected, output)
		}
	}
}
