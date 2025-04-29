package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryasaputra2016/essentask/config"
)

func main() {
	db, err := config.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	println("serving at 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
