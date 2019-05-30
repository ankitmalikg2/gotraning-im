package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
)

func getpeople(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get ")
	io.WriteString(w, "functions")
	w.Write([]byte("ankit malikg"))
}

func main() {
	config := newrelic.NewConfig("mux application", "a5ffe1304135caedc86cbf27b0696380a27cb791")
	app, err := newrelic.NewApplication(config)
	if err != nil {
		fmt.Println("Error:", err.Error)
	}

	router := mux.NewRouter()
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/ad", getpeople)).Methods("GET")

	// http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))
	// http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8080", router)
}
