package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//Greet says hello to name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//MyGreeterHandler shows how powerful are Go interfaces for DI, using a HTTP server
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// Excute go run greeter.go and see the outcome in http://localhost:5000
func main() {
	Greet(os.Stdout, "Eloide")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
