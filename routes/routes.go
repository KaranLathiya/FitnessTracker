package routes

import (
	"karanlathiya/FitnessTracker/handlers"
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
	r.Get("/user/profile/add", handleCORS(handlers.UserProfileAdd))
	r.Get("/user/profile/show", handleCORS(handlers.UserProfileShow))
	r.Get("/user/exercise", handleCORS(handlers.AddExerciseDetails))
	r.Get("/user/meal", handleCORS(handlers.AddMealDetails))
	r.Get("/user/weight", handleCORS(handlers.AddMealDetails))
	r.Get("/user/logout", handleCORS(handlers.Logout))

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("route does not exist"))
	})

	return r
}
