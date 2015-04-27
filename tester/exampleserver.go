package main

import (
	"fmt"
	"net/http"
	"os"

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
	go handleCircleData(respChan)
	http.HandleFunc("/v1/builds", handler)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
