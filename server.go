package circle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const VERSION = 1

func NewHandler(rc chan *Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nd := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var resp *Response
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
