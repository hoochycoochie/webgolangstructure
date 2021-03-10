package api

import (
	v1 "financeapp/internal/api/v1"
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter provide handler API SERVICE
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
