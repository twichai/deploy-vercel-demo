package handler

import (
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Go Serverless Function!"))
}
