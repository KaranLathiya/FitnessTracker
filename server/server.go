package main

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	// "github.com/go-chi/chi/v5/middleware"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
	defer db.Close()
}
