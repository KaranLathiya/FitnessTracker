package routes

import (
	// "encoding/json"
	"karanlathiya/FitnessTracker/handlers"
	// "karanlathiya/FitnessTracker/models"
	"net/http"

	"github.com/go-chi/chi"
)


func handleCORS(handler http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "json/application")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the original handler
		handler(w, r)
	}
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/user/signup", handleCORS(handlers.UserSignup))
	r.Post("/user/login", handleCORS(handlers.UserLogin))
	
	r.Get("/user/profile/add", handleCORS(handlers.AddUserProfileDetails))
	r.Get("/user/exercise/add", handleCORS(handlers.AddExerciseDetails))
	r.Get("/user/meal/add", handleCORS(handlers.AddMealDetails))
	r.Get("/user/weight/add", handleCORS(handlers.AddWeightDetails))
	r.Get("/user/water/add", handleCORS(handlers.AddWaterDetails))

	r.Get("/user/profile/show", handleCORS(handlers.UserProfileShow))
	
	r.Put("/user/profile/update", handleCORS(handlers.EditUserProfileDetails))
	r.Put("/user/exercise/update", handleCORS(handlers.EditExerciseDetails))
	r.Put("/user/meal/update", handleCORS(handlers.EditMealDetails))
	r.Put("/user/weight/update", handleCORS(handlers.EditWeightDetails))
	r.Put("/user/water/update", handleCORS(handlers.EditWaterDetails))
	
	r.Delete("/user/profile/delete", handleCORS(handlers.EditUserProfileDetails))
	r.Delete("/user/exercise/delete", handleCORS(handlers.EditExerciseDetails))
	r.Delete("/user/meal/delete", handleCORS(handlers.EditMealDetails))
	r.Delete("/user/weight/delete", handleCORS(handlers.EditWeightDetails))
	r.Delete("/user/water/delete", handleCORS(handlers.EditWaterDetails))
	
	r.Get("/user/logout", handleCORS(handlers.Logout))

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("wrong method"))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	return r
}
