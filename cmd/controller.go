package cmd

import (
	"encoding/json"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, r.FormValue("n"))
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
