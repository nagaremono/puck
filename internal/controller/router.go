package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("service is healtyyyy"))
		w.WriteHeader(http.StatusOK)
	})

	return r
}
