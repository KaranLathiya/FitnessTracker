package main

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	_ "karanlathiya/FitnessTracker/docs"
	"karanlathiya/FitnessTracker/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	httpSwagger "github.com/swaggo/http-swagger"
	// "github.com/go-chi/chi/v5/middleware"
)

//	@title			Fitnesstracker API
//	@version		1.0
//	@description	This is a sample Fitnessstracker server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host		localhost:8080

// @securitydefinitions.apikey UserIDAuth
// @in header
// @name Authorization

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
	if err != nil {
		panic(err)
	}
	// dal.InitDB(db)
	//Routing
	
	defer db.Close()
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
