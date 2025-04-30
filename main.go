package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suryasaputra2016/essentask/config"
)

func main() {
	db, err := config.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = config.PrepareTables(db)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	port := os.Getenv("WEB_PORT")
	log.Printf("serving at %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!!")
}
