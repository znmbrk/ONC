package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func teamGet(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from team;")

	if err != nil {
		http.Error(w, "Failed to query teams from the database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var team Team
	var teamSlice = []Team{}

	for rows.Next() {
		err := rows.Scan(&team.ID, &team.Name)
		if err != nil {
			http.Error(w, "Failed to scan team row", http.StatusInternalServerError)
			return
		}
		teamSlice = append(teamSlice, team)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teamSlice)
}

func teamPost(w http.ResponseWriter, r *http.Request) {
	var team = Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if team.ID == 0 {
		http.Error(w, "ID cannot be empty", http.StatusBadRequest)
		return
	}
	_, err = db.Exec("insert into team (id, name) values ($1, $2)", team.ID, team.Name)
	if err != nil {
		http.Error(w, "POST method failed to insert team", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

func specificTeamGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	var team = Team{}

	err := db.QueryRow("select * from team where id = $1", id).Scan(&team.ID, &team.Name)
	if err != nil {
		http.Error(w, "Failed to query team from database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

func specificTeamPut(w http.ResponseWriter, r *http.Request) {
	// URL + JSON --> GO --> DATABASE
	vars := mux.Vars(r)
	id := vars["ID"]
	var team = Team{}

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		http.Error(w, "Failed to decode the request body", http.StatusBadRequest)
		return
	}
	_, err = db.Exec("update team set name = $1 where id = $2", team.Name, id)
	if err != nil {
		http.Error(w, "PUT method failed to update team", http.StatusInternalServerError)
		return
	}
	team.ID, err = strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Couldn't convert id to integer", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)

}

func specificTeamDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	_, err := db.Exec("delete from team where id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete item from team", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"deleted": id})
}
