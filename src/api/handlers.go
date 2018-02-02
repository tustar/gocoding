package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

func GetFolder(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	panicErr(err)

	folder := queryFolder(id)
	json.NewEncoder(rw).Encode(folder)
}

func GetFolders(rw http.ResponseWriter, req *http.Request) {
	json.NewEncoder(rw).Encode(queryFolders())
}