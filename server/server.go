package main

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/routes"
	"log"
	"net/http"
	"os"
	// "github.com/go-chi/chi/v5/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	fmt.Println("Server started")
	db, err := dal.Connect()
	errors.CheckErr(err)
	dal.InitDB(db)
	
	//Routing
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, nil))
	defer db.Close()
}
