package main

import (
	"fmt"
	"go-workshop-practical-me/cmd/project"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	handler := project.NewHelloWorldHandler(os.Getenv("OPENAI_TOKEN"))
	mux.Handle("/", handler)
	log.Print("Server started and listening on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
