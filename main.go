package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUserHandler)

	http.ListenAndServe(":8080", mux)

}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List User"))
}
