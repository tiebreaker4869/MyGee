package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "this is index page")
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "this is hello page")
	})

	log.Fatal(r.Run(":8080"))
}
