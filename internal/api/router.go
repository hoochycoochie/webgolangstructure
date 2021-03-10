package api

import (
	"net/http"

	v1 "github.com/hoochycoochie/webgolangstructure/internal/api/v1"

	"github.com/gorilla/mux"
)

//NewRouter provide handler API SERVICE
func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil
}
