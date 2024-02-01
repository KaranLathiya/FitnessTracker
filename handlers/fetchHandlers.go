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
	db := dal.GetDB()
	var user models.Users
	rows, err := db.Query("select email, fullname, age, gender, height, weight, health_goal, profile_photo from public.user_details where user_id=$1", UserID.UserID)
	// errIfZeroRows := db.QueryRow("select email, fullname from public.user_registration_details where user_id=$1", UserID.UserID).Scan(&user.Email, &user.FullName)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	for rows.Next() {
		err := rows.Scan(&user.Email, &user.FullName, &user.Age, &user.Gender, &user.Height, &user.Weight, &user.HealthGoal, &user.ProfilePhoto)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("null"))
			return
		}
	}
	defer rows.Close()

	user_data, _ := json.MarshalIndent(user, "", "  ")
	w.Write(user_data)
}

func FetchAllDetails(w http.ResponseWriter, r *http.Request) {
	allDetailsMap := make(map[string]interface{})
	date := r.FormValue("date")
	//fmt.Println(date)
	exercise, err := fetchExerciseDetails(date)
	if err != nil {
		errors.MessageShow(500, err.Error(), w)
	}
	meal, err := fetchMealDetails(date)
	if err != nil {
		errors.MessageShow(500, err.Error(), w)
	}
	weight, err := fetchWeightDetails(date)
	if err != nil {
		errors.MessageShow(500, err.Error(), w)
	}
	water, err := fetchWaterDetails(date)
	if err != nil {
		errors.MessageShow(500, err.Error(), w)
	}
	allDetailsMap["waterDetails"] = water
	allDetailsMap["weightDetails"] = weight
	allDetailsMap["exerciseDetails"] = exercise
	allDetailsMap["mealDetails"] = meal
	allDetails, _ := json.Marshal(allDetailsMap)
	w.Write(allDetails)
}

func fetchExerciseDetails(date string) (interface{}, error) {
	db := dal.GetDB()

	var exercise []models.Exercise
	rows, err := db.Query("select exercise_type, duration, calories_burned from public.exercise_details where user_id=$1 AND date=$2", UserID.UserID, date)
	if err != nil {
		return exercise, err
	}
	i := 0
	for rows.Next() {
		emptyExercise := models.Exercise{}
		exercise = append(exercise, emptyExercise)
		err := rows.Scan(&exercise[i].ExerciseType, &exercise[i].Duration, &exercise[i].CaloriesBurned)
		if err != nil {
			return exercise, err
		}
		i += 1
	}
	defer rows.Close()
	_, _ = json.MarshalIndent(exercise, "", "  ")
	return exercise, err
}

func fetchMealDetails(date string) (interface{}, error) {
	db := dal.GetDB()
	var meal []models.Meal
	rows, err := db.Query("select meal_type, ingredients, calories_consumed from public.meal_details where user_id=$1 and date=$2", UserID.UserID, date)
	if err != nil {
		return meal, err
	}
	i := 0
	for rows.Next() {
		emptyMeal := models.Meal{}
		meal = append(meal, emptyMeal)
		err := rows.Scan(&meal[i].MealType, &meal[i].Ingeredients, &meal[i].CaloriesConsumed)
		if err != nil {
			return meal, err
		}
		i += 1
	}
	defer rows.Close()
	_, _ = json.MarshalIndent(meal, "", "  ")
	return meal, err
}
func fetchWeightDetails(date string) (interface{}, error) {
	db := dal.GetDB()
	var weight models.Weight
	rows, err := db.Query("select daily_weight from public.weight_details where user_id=$1 AND date=$2 ", UserID.UserID, date)
	if err != nil {
		return weight, err
	}

	i := 0
	for rows.Next() {
		err := rows.Scan(&weight.DailyWeight)
		if err != nil {
			return weight, err
		}
		i += 1
	}
	if i == 0 {
		return nil, err
	}
	defer rows.Close()

	_, _ = json.MarshalIndent(weight, "", "  ")
	return weight, err
}

func fetchWaterDetails(date string) (interface{}, error) {
	db := dal.GetDB()
	var water models.Water
	rows, err := db.Query("select water_intake from public.water_details where user_id=$1 AND date=$2", UserID.UserID, date)
	if err != nil {
		return water, err
	}

	i := 0
	for rows.Next() {
		err := rows.Scan(&water.WaterIntake)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return water, err
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		return nil, err
	}
	_, _ = json.MarshalIndent(water, "", "  ")
	return water, err
}
