package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type handler struct {
}

type payload struct {
	Summary    string
	StatusCode int
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var p payload

	switch req.Method {
	case "GET":
		p.Summary = fmt.Sprintf("GET %s", req.RequestURI)
	default:
		p.Summary = fmt.Sprintf("unknown method %s", req.Method)
		p.StatusCode = http.StatusBadRequest
	}

	body, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error during JSON marshaling: %v", err)))
		return
	}

	if p.StatusCode != 0 {
		w.WriteHeader(p.StatusCode)
	}

	_, err = w.Write(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during write to response: %v", err)
		return
	}
}

func main() {
	fmt.Println("Starting test HTTP server ...")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      &handler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
