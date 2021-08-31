package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/troydai/blocks/tserver/handler"
)

type payload struct {
	Summary    string
	StatusCode int
}

func main() {
	fmt.Println("Starting test HTTP server ...")

	h, wait := handler.NewHandler()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		s.ListenAndServe()
	}()
	defer StopServer(s)

	wait()
}

func StopServer(s *http.Server) {
	fmt.Println("Exiting the server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
