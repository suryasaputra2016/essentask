package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suryasaputra2016/essentask/config"
	"github.com/suryasaputra2016/essentask/handlers"
	"github.com/suryasaputra2016/essentask/repo"
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

	userRepo := repo.NewuserRepo(db)
	userHandler := handlers.NewUserHandler(userRepo)

	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	router.HandleFunc("/register", userHandler.Register)
	// router.HandleFunc("/login", userHandler.Login)
	// router.HandleFunc("/logout", userHandler.Logout)

	port := os.Getenv("WEB_PORT")
	log.Printf("serving at %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!!")
}
