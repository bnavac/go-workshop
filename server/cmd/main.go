package main

import (
	"app/app/test"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("/", app.Hello)
	http.HandleFunc("/headers", app.Headers)
	var port int = 3000
	//password := app.helper("asdf")
	log.Printf("starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
