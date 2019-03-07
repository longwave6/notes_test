package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "longwave_api/config"
	. "longwave_api/dao"
	. "longwave_api/routes"
)

var config = Config{}
var dao = GeneralDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", AllNotesEndPoint).Methods("GET")
	r.HandleFunc("/notes", CreateNoteEndPoint).Methods("POST")
	r.HandleFunc("/notes", UpdateNoteEndPoint).Methods("PUT")
	r.HandleFunc("/notes", DeleteNoteEndPoint).Methods("DELETE")
	r.HandleFunc("/notes/{id}", FindNoteEndpoint).Methods("GET")
	r.HandleFunc("/users", AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", FindUserEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
