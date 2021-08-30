package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type handler struct {
	closec chan struct{}
}

type payload struct {
	Summary    string
	StatusCode int
}

func newHandler() *handler {
	return &handler{
		closec: make(chan struct{}),
	}
}

func (h *handler) CloseChannel() <-chan struct{} {
	return h.closec
}

func (h *handler) close() {
	h.closec <- struct{}{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var p payload

	switch req.Method {
	case "GET":
		p.Summary = fmt.Sprintf("GET %s", req.RequestURI)
	case "POST":
		p.Summary = fmt.Sprintf("POST %s", req.RequestURI)
		if req.RequestURI == "/close" {
			h.close()
		}
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

	h := newHandler()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		s.ListenAndServe()
	}()

	<-h.CloseChannel()
	fmt.Println("Exiting the server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
