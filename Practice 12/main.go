package main

import (
	"fmt"
	"net/http"
)

func serverHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working......")
	w.Write([]byte("<h1>Server Starting!</h1>"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome)
	fmt.Println("Server is listening on port 8080...")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
