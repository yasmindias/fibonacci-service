package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/yasmindias/fibonacci-service/cmd"
)

func main() {
	runHTTPServer()
}

func runHTTPServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib", controller.Get).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
