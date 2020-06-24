package slack

import (
	"testing"
)

const (
	apiToken = "xoxs-123456789123-123456789123-123456789123-2943a567bc05bc66ca6201dbc5f00bec3f774a47b1b94289a2ae8e79834c21a5"
	text     = `"api_token":"` + apiToken + `","hello":"100"}`
)

func TestReadToken(t *testing.T) {
	token, err := readToken(text)
	if err != nil {
		t.Fatal(err)
	}
	want := apiToken
	if want != token {
		t.Fatalf("want %s but got %s", want, token)
	}
}
