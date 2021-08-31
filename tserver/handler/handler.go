package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/troydai/blocks/tserver/data"
)

type (
	handler struct {
		closec chan struct{}
	}

	Wait func()
)

func NewHandler() (http.Handler, Wait) {
	h := &handler{
		closec: make(chan struct{}),
	}
	w := func() {
		<-h.closec
	}

	return h, w
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var p data.Payload

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

func (h *handler) close() {
	h.closec <- struct{}{}
}
