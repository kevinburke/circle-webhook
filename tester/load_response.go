package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kevinburke/circle-webhook"
)

func main() {
	f, err := os.Open("payload.json")
	if err != nil {
		log.Fatal(err)
	}
	var r circle.Response
	nd := json.NewDecoder(f)
	nd.Decode(&r)
	//fmt.Println(r.Payload)
	fmt.Println(r.Payload.Owners[0])
}
