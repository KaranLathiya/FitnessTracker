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

	date := r.FormValue("date")
	//fmt.Fprintf(w, " %s\n", date)
	//fmt.Println(date)
	rows, err := db.Query("select exercise_type, duration, calories_burned, date from public.exercise_details where user_id=$1 AND date=$2", UserID.UserID, date)
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
	date := r.FormValue("date")
	//fmt.Fprintf(w, " %s\n", date)
	//fmt.Println(date)
	rows, err := db.Query("select meal_type, ingredients, calories_consumed, date from public.meal_details where user_id=$1 and date=$2", UserID.UserID, date)
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
	date := r.FormValue("date")
	//fmt.Fprintf(w, " %s\n", date)
	//fmt.Println(date)
	rows, err := db.Query("select daily_weight, date from public.weight_details where user_id=$1 AND date=$2 ", UserID.UserID, date)
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
	
	weight_data, _ := json.MarshalIndent(weight, "", "  ")
	w.Write(weight_data)
}

func FetchWaterDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	date := r.FormValue("date")
	//fmt.Fprintf(w, " %s\n", date)
	//fmt.Println(date)
	rows, err := db.Query("select water_intake, date from public.water_details where user_id=$1 AND date=$2", UserID.UserID, date)
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
	
	water_data, _ := json.MarshalIndent(water, "", "  ")
	w.Write(water_data)
}
func FetchWaterIntakeMonthly(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select water_intake, date from public.water_details where user_id=$1 AND date >= NOW() - INTERVAL '30 days' order by date desc", UserID.UserID)
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

func FetchAllDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	allDetailsMap := make(map[string]interface{})
	db = dal.GetDB()
	date := r.FormValue("date")
	//fmt.Println(date)
	rows, err := db.Query("select water_intake, date from public.water_details where user_id=$1 AND date=$2", UserID.UserID, date)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var water models.Water
	for rows.Next() {
		err := rows.Scan(&water.WaterIntake, &water.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
	}
	defer rows.Close()
	rows, err = db.Query("select exercise_type, duration, calories_burned, date from public.exercise_details where user_id=$1 AND date=$2", UserID.UserID, date)
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

	rows, err = db.Query("select meal_type, ingredients, calories_consumed, date from public.meal_details where user_id=$1 and date=$2", UserID.UserID, date)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var meal []models.Meal
	i = 0
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

	rows, err = db.Query("select daily_weight, date from public.weight_details where user_id=$1 AND date=$2 ", UserID.UserID, date)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var weight models.Weight
	for rows.Next() {
		err := rows.Scan(&weight.DailyWeight, &weight.Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
	}
	defer rows.Close()
	_, _ = json.MarshalIndent(water, "", "  ")
	_, _ = json.MarshalIndent(weight, "", "  ")
	_, _ = json.MarshalIndent(exercise, "", "  ")
	_, _ = json.MarshalIndent(meal, "", "  ")
	allDetailsMap["water"] = water
	allDetailsMap["weight"] = weight
	allDetailsMap["exercise"] = exercise
	allDetailsMap["meal"] = meal
	allDetails, _ := json.Marshal(allDetailsMap)
	
	w.Write(allDetails)

}


