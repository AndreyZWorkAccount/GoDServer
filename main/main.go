package main

import (
	"fmt"
	"net/http"
)

func main() {

	var route = "/index"
	var port = 9292

	http.HandleFunc(route, func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming request.")
		rw.Header().Set("Server", "A Go Web Server")
		rw.WriteHeader(200)
		rw.Write([]byte("OK"))
	})

	fmt.Printf("Server starting at %s ...", route)

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
