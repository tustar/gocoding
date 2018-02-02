package main

import "github.com/gorilla/mux"

func routeFolder(router *mux.Router)  {
	router.HandleFunc("/folders", GetFolders).Methods("GET")
	router.HandleFunc("/folders/{id}", GetFolder).Methods("GET")
}