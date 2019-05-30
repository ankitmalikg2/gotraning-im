package main

import (
	"fmt"
	"io"
	"net/http"

	newrelic "github.com/newrelic/go-agent"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hello world")
	io.WriteString(w, "Hello world")
	w.Write([]byte("ankit malikg"))
}

func main() {
	config := newrelic.NewConfig("Your App Name", "a5ffe1304135caedc86cbf27b0696380a27cb791")
	app, err := newrelic.NewApplication(config)
	if err != nil {
		fmt.Println("Error:", err.Error)
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))

	// http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
