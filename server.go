package circle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const VERSION = 1

// Create a HTTP handler that parses incoming build hooks from CircleCI and
// emits the parsed Response structs to the given channel.
func NewHandler(rc chan *Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//tee := io.TeeReader(r.Body, os.Stdout)
		nd := json.NewDecoder(r.Body)
		defer r.Body.Close()
		resp := new(Response)
		err := nd.Decode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rc <- resp
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Server", fmt.Sprintf("circle-webhooks/v%d", VERSION))
		w.Write([]byte("thanks"))
	}
}
