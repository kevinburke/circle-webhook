package circle

import (
	"encoding/json"
	"os"
	"testing"
)

func TestPayloadParse(t *testing.T) {
	f, err := os.Open("data/payload.json")
	if err != nil {
		t.Fatal(err)
	}
	r := new(Response)
	nd := json.NewDecoder(f)
	err = nd.Decode(r)
	if err != nil {
		t.Fatal(err)
	}
	if r.Payload.AuthorEmail != "kev@inburke.com" {
		t.Errorf("Expected AuthorEmail to be 'kev@inburke.com', was \"%s\"", r.Payload.AuthorEmail)
	}

	if r.Payload.BuildNum != 2 {
		t.Errorf("Expected BuildNum to be 2, was %d", r.Payload.BuildNum)
	}
}
