package gateway

import (
	"net/http"
	"tss_project/handler/gateway"

	"github.com/gorilla/mux"
)

func NewServer() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/keygen", gateway.KeygenHandler).Methods("POST")

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
