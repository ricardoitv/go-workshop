package project

import (
	"encoding/json"
	"fmt"
	"go-workshop-practical-me/cmd/openai"
	"io"
	"net/http"
	"os"
)

// an interface from net/http or something like that
type HelloWorldHandler struct {
	OpenAIClient openai.OpenAI
}

type request struct {
	Transcript string `json:"transcript"`
}
type response struct {
	Summary string `json:"summary"`
}

func NewHelloWorldHandler(token string) HelloWorldHandler {
	openAi := openai.NewOpenAI(token)
	return HelloWorldHandler{
		OpenAIClient: openAi,
	}
}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var requestMessage request
	err := json.NewDecoder(r.Body).Decode(&requestMessage)
	// see https://stackoverflow.com/a/32718077
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		return
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if requestMessage.Transcript == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// calling openai
	summary, err := h.OpenAIClient.Summarise(requestMessage.Transcript)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	responseMessage := response{Summary: summary}
	err = json.NewEncoder(w).Encode(responseMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
