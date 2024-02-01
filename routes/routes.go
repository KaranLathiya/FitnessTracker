package routes

import (
	// "encoding/json"
	"karanlathiya/FitnessTracker/handlers"
	// "karanlathiya/FitnessTracker/models"
	"net/http"

	"github.com/go-chi/chi"
)

// func handleCORS(handler http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Set CORS headers
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Content-Type", "json/application")
// 		// Handle preflight requests
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}

// 		// Call the original handler
// 		handler(w, r)
// 	}
// }

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {

		r.Use(handlers.HandleCORS)
		r.Post("/signup", handlers.UserSignup)
		r.Post("/login", handlers.UserLogin)
		r.Post("/forgot-password", handlers.ForgotPassword)
		r.Post("/verify-otp", handlers.VerifyOTP)
		
		r.Route("/user", func(r chi.Router) {
			r.Use(handlers.Authentication)

			r.Put("/change-password", handlers.UpdateUserPassword)

			r.Route("/profile", func(r chi.Router) {
				r.Put("/", handlers.UpdateUserProfileDetails)
				r.Get("/", handlers.FetchUserProfileDetails)
			})

			r.Route("/exercise", func(r chi.Router) {
				r.Post("/", handlers.AddExerciseDetails)
				// r.Get("/", handlers.FetchExerciseDetails)
				r.Put("/", handlers.UpdateExerciseDetails)
				r.Delete("/", handlers.DeleteExerciseDetails)
			})
			r.Route("/meal", func(r chi.Router) {
				r.Post("/", handlers.AddMealDetails)
				// r.Get("/", handlers.FetchMealDetails)
				r.Put("/", handlers.UpdateMealDetails)
				r.Delete("/", handlers.DeleteMealDetails)
			})
			r.Route("/weight", func(r chi.Router) {
				r.Post("/", handlers.AddWeightDetails)
				// r.Get("/", handlers.FetchWeightDetails)
				r.Put("/", handlers.UpdateWeightDetails)
				r.Delete("/", handlers.DeleteWeightDetails)
			})
			r.Route("/water", func(r chi.Router) {
				r.Post("/", handlers.AddWaterDetails)
				// r.Get("/", handlers.FetchWaterDetails)
				r.Put("/", handlers.UpdateWaterDetails)
				r.Delete("/", handlers.DeleteWaterDetails)
			})
			r.Get("/yearly-weight-details", handlers.FetchYearlyWeightDetails)
			r.Get("/yearly-caloriesburned-details", handlers.FetchYearlyCaloriesBurnedDetails)
			// r.Get("/water-intake-of-month", handlers.FetchWaterIntakeMonthly)
			r.Get("/alldetails", handlers.FetchAllDetails)
		})

		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(405)
			w.Write([]byte("wrong method"))
		})
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(404)
			w.Write([]byte("route does not exist"))
		})

	})

	return r
}
