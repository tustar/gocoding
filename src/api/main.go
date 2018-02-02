package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Api
	router := mux.NewRouter()
	routeFolder(router)
	http.ListenAndServe(":4000", router)
}



