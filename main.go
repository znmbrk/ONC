package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=zainmobarik dbname=on_call sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/team", teamGet).Methods("GET")
	r.HandleFunc("/team", teamPost).Methods("POST")
	r.HandleFunc("/team/{ID}", specificTeamGet).Methods("GET")
	r.HandleFunc("/team/{ID}", specificTeamPut).Methods("PUT")
	r.HandleFunc("/team/{ID}", specificTeamDelete).Methods("DELETE")
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
