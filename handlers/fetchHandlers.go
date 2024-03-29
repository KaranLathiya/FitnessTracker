package handlers

import (
	"encoding/json"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"
)

// FetchUserProfileDetails example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		fetch user profile details 
//	@Description	fetch user profile with Email, FullName, Age, Gender, Height, Weight, HealthGoal, ProfilePhoto
//	@ID				user-profiledetails
//	@Produce		json
//	@Success		200		{object}	models.AllDetails	
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/profile [get]
func FetchUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	db := dal.GetDB()
	var user models.Users
	rows, err := db.Query("SELECT email, fullname, age, gender, height, weight, health_goal, profile_photo FROM public.user_details WHERE user_id=$1", UserID.UserID)
	// errIfZeroRows := db.QueryRow("select email, fullname from public.user_registration_details where user_id=$1", UserID.UserID).Scan(&user.Email, &user.FullName)
	if err != nil {
		response.MessageShow(500, "Internal Server Error", w)
		return
	}
	for rows.Next() {
		err := rows.Scan(&user.Email, &user.FullName, &user.Age, &user.Gender, &user.Height, &user.Weight, &user.HealthGoal, &user.ProfilePhoto)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
	}
	defer rows.Close()
	userProfileDetails, _ := json.MarshalIndent(user, "", "  ")
	w.Write(userProfileDetails)
}

// FetchAllDetails example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		fetch all details 
//	@Description	fetch all details by date 
//	@ID				user-alldetails-fetch
//	@Produce		json
// @Param   date     query     string     true  "date for which want to see details"     example("2024-01-01")
//	@Success		200		{object}	models.AllDetails	
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/alldetails [get]
func FetchAllDetails(w http.ResponseWriter, r *http.Request) {
	// allDetailsMap := make(map[string]interface{})
	date := r.FormValue("date")
	dateLayout := "2006-01-02"
	_, err := time.Parse(dateLayout, date)
	if err != nil {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	exercise, err := fetchExerciseDetails(date)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	meal, err := fetchMealDetails(date)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	weight, err := fetchWeightDetails(date)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	water, err := fetchWaterDetails(date)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	// allDetailsMap["waterDetails"] = water
	// allDetailsMap["weightDetails"] = weight
	// allDetailsMap["exerciseDetails"] = exercise
	// allDetailsMap["mealDetails"] = meal
	var allDetails models.AllDetails
	allDetails.ExerciseDetails = exercise
	allDetails.MealDetails = meal
	if weight.DailyWeight == 0 {
		allDetails.WeightDetails = nil
	} else {
	allDetails.WeightDetails = &weight
	}
	if weight.DailyWeight == 0 {
		allDetails.WaterDetails =  nil
	} else {
	allDetails.WaterDetails = &water
	}
	userAlldetails, _ := json.MarshalIndent(allDetails, "", "  ")
	w.Write(userAlldetails)
}

func fetchExerciseDetails(date string) ([]models.Exercise, error) {
	db := dal.GetDB()
	var exercise []models.Exercise
	rows, err := db.Query("SELECT exercise_type, duration, calories_burned FROM public.exercise_details WHERE user_id=$1 AND date=$2", UserID.UserID, date)
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

func fetchMealDetails(date string) ([]models.Meal, error) {
	db := dal.GetDB()
	var meal []models.Meal
	rows, err := db.Query("SELECT meal_type, ingredients, calories_consumed FROM public.meal_details WHERE user_id=$1 AND date=$2", UserID.UserID, date)
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

func fetchWeightDetails(date string) (models.Weight, error) {
	db := dal.GetDB()
	var weight models.Weight
	rows, err := db.Query("SELECT daily_weight FROM public.weight_details WHERE user_id=$1 AND date=$2 ", UserID.UserID, date)
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
		return weight, err
	}
	defer rows.Close()
	_, _ = json.MarshalIndent(weight, "", "  ")
	return weight, err
}

func fetchWaterDetails(date string) (models.Water, error) {
	db := dal.GetDB()
	var water models.Water
	rows, err := db.Query("SELECT water_intake FROM public.water_details WHERE user_id=$1 AND date=$2", UserID.UserID, date)
	if err != nil {
		return water, err
	}
	i := 0
	for rows.Next() {
		err := rows.Scan(&water.WaterIntake)
		if err != nil {
			// fmt.Println("Error scanning row:", err)
			return water, err
		}
		i += 1
	}
	defer rows.Close()
	if i == 0 {
		return water, err
	}
	_, _ = json.MarshalIndent(water, "", "  ")
	return water, err
}
