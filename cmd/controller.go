package cmd

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var n, _ = strconv.Atoi(r.FormValue("n"))
	var result = FibonacciRecursion(n)

	respondWithJSON(w, http.StatusOK, result)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
