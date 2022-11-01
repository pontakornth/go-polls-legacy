package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "not found"})
}
func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
