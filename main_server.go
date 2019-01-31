// A basic HTTP server.
// By default, it serves the current working directory on port 8080.
//
// Note: This doesn't work in WebAssembly :( Has deadlock
package main

import (
  "io"
  "log"
  "net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello hello\n")
	}
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
