package main

import (
	"database/sql"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUserHandler)

	http.ListenAndServe(":8080", mux)

}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
}
