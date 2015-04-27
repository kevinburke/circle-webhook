# Webhooks from Circle

This is a library that makes it really easy to parse Circle webhooks.

### Payload Structure

The webhook response structure is defined in `payload.go`. It's a little bit
incomplete, because the webhooks I received had several `null` values. Still,
most of the important fields should be there.

### Responding to Incoming Webhooks

Create a channel that can send and receive Response objects:

```go
respChan := make(chan *circle.Response)
```

Create a handler function that takes a receive-only channel and parses it
however you like. Here's a simple one that prints the build author's email.
Note, this currently will block the server thread until parsing is done.

```go
func handleCircleData(cs <-chan *circle.Response) {
	resp := <-cs
	fmt.Println(resp.Payload.AuthorEmail)
}
```

Putting it all together, and creating a server:

```go
import (
    "fmt"
    "net/http"

	"github.com/gorilla/handlers"
    "github.com/kevinburke/circle-webhook"
)

func handleCircleData(cs <-chan *circle.Response) {
	resp := <-cs
	fmt.Println(resp.Payload.AuthorEmail)
}

func main() {
	respChan := make(chan *circle.Response)
	handler := circle.NewHandler(respChan)
	// Launch it in a goroutine, so it doesn't block
	go handleCircleData(respChan)
	http.HandleFunc("/v1/builds", handler)
	// LoggingHandler will log a line for each incoming request
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
```

That's it! Then use `ngrok` locally or point your build to
yourdomain.com:8080/v1/builds to get some sweet Circle webhook action.
