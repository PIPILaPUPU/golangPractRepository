package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var names []string

func handleUser(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || index < 0 || index >= len(names) {
		http.Error(w, "Некорректный идентификатор", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(names[index]))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(names)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		return
	}

	names = append(names, name)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(len(names) - 1)))
}

func main() {
	r := chi.NewRouter()

	r.Get("/users", handleUsers)

	r.Get("/user/{id}", handleUser)

	r.Post("/user", handleAddUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}
