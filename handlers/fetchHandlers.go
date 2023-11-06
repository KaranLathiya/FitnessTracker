package handlers

import (
	"encoding/json"
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"net/http"
)

func FetchUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	var user models.Users
	rows, err := db.Query("select age, gender, height, weight, health_goal, profile_photo from public.user_details where user_id=$1", UserID.UserID)
	// errIfZeroRows := db.QueryRow("select email, fullname from public.user_registration_details where user_id=$1", UserID.UserID).Scan(&user.Email, &user.FullName)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&user.Age, &user.Gender, &user.Height, &user.Weight, &user.HealthGoal, &user.ProfilePhoto)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		w.Write([]byte("[]"))
		return
	}
	user_data, _ := json.MarshalIndent(user, "", "  ")
	w.Write(user_data)
}

func FetchExerciseDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select exercise_type, duration, calories_burned, date from public.exercise_details where user_id=$1 AND date >= NOW() - INTERVAL '7 days'order by date desc", UserID.UserID)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var exercise []models.Exercise
	i := 0
	for rows.Next() {
		emptyExercise := models.Exercise{}
		exercise = append(exercise, emptyExercise)
		err := rows.Scan(&exercise[i].ExerciseType, &exercise[i].Duration, &exercise[i].CaloriesBurned, &exercise[i].Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		w.Write([]byte("[]"))
		return
	}
	exercise_data, _ := json.MarshalIndent(exercise, "", "  ")
	w.Write(exercise_data)
}
func FetchMealDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select meal_type, ingredients, calories_consumed, date from public.meal_details where user_id=$1 and date >= NOW() - INTERVAL '7 days'order by date desc", UserID.UserID)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var meal []models.Meal
	i := 0
	for rows.Next() {
		emptyMeal := models.Meal{}
		meal = append(meal, emptyMeal)
		err := rows.Scan(&meal[i].MealType, &meal[i].Ingeredients, &meal[i].CaloriesConsumed, &meal[i].Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		w.Write([]byte("[]"))
		return
	}
	fmt.Println(meal)
	meal_data, _ := json.MarshalIndent(meal, "", "  ")
	w.Write(meal_data)

}
func FetchWeightDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select daily_weight, date from public.weight_details where user_id=$1 AND date >= NOW() - INTERVAL '7 days'order by date desc ", UserID.UserID)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var weight []models.Weight
	i := 0
	for rows.Next() {
		emptyWeight := models.Weight{}
		weight = append(weight, emptyWeight)
		err := rows.Scan(&weight[i].DailyWeight, &weight[i].Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		w.Write([]byte("[]"))
		return
	}
	weight_data, _ := json.MarshalIndent(weight, "", "  ")
	w.Write(weight_data)
}

func FetchWaterDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select water_intake, date from public.water_details where user_id=$1 AND date >= NOW() - INTERVAL '7 days'order by date desc", UserID.UserID)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var water []models.Water
	i := 0
	for rows.Next() {
		emptyWater := models.Water{}
		water = append(water, emptyWater)
		err := rows.Scan(&water[i].WaterIntake, &water[i].Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		w.Write([]byte("[]"))
		return
	}
	water_data, _ := json.MarshalIndent(water, "", "  ")
	w.Write(water_data)
}
