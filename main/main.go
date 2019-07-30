package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/index", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming request.")
		rw.Header().Set("Server", "A Go Web Server")
		rw.WriteHeader(200)
		rw.Write([]byte("OK"))
	})

	fmt.Println("Server starting...")
	fmt.Println(http.ListenAndServe(":9292", nil))
}
