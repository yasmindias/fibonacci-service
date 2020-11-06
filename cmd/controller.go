package cmd

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.FormValue("n"))
	if err == nil {
		result := FibonacciRecursion(n)
		respondWithJSON(w, http.StatusOK, result)
	} else {
		respondWithJSON(w, http.StatusBadRequest, "Error: input must be an integer")
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
