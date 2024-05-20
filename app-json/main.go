package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/resonantChaos22/toolkit"
)

type RequestBody struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponseBody struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

var tools toolkit.Tools

func main() {
	mux := routes()
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()
	api := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.Handle("/api/", http.StripPrefix("/api", api))

	api.HandleFunc("/receive-post", handleReceivePost)
	api.HandleFunc("/remote-service", handleRemoteService)
	api.HandleFunc("/simulated-service", handleSimulatedService)

	return mux
}

func handleReceivePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Called Receive Post API")
	reqBody := new(RequestBody)
	err := tools.ReadJSON(w, r, reqBody)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}
	resPayload := ResponseBody{
		Message: fmt.Sprintf("Hit the handler okay, now sending the response - %s", reqBody.Message),
	}
	tools.WriteJSON(w, http.StatusAccepted, &resPayload)
}

func handleRemoteService(w http.ResponseWriter, r *http.Request) {
	log.Println("Called Remote Service API")
	reqBody := new(RequestBody)
	err := tools.ReadJSON(w, r, reqBody)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	_, statusCode, err := tools.PushJSONToRemote("http://localhost:8081/api/simulated-service", reqBody)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	resPayload := ResponseBody{
		Message:    fmt.Sprintf("Hit the handler okay, now sending the response - %s", reqBody.Message),
		StatusCode: statusCode,
	}
	tools.WriteJSON(w, http.StatusAccepted, &resPayload)
}

func handleSimulatedService(w http.ResponseWriter, r *http.Request) {
	log.Println("Called Simulated Service API")
	payload := ResponseBody{
		Message: "OK",
	}

	_ = tools.WriteJSON(w, http.StatusOK, payload)
}
