package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter();
	r.HandleFunc("/", serverHome)
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server Starting!</h1>"))
}
