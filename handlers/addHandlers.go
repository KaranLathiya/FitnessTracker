package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"net/http"
	"time"
)

func AddUserProfileDetails(w http.ResponseWriter, r *http.Request) {

	var user models.Users
	db = dal.GetDB()
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	user.UserID, _ = session.Values["userID"].(int)
	if _, err := db.Exec(
		"INSERT INTO public.user_profile_details( user_id, age, gender, height, weight, health_goal, profile_photo) VALUES ( $1, $2, $3, $4, $5, $6, $7);", user.UserID, user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "User details Successfully added")
}

func AddExerciseDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var exercise models.Exercise
	db = dal.GetDB()
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}

	RowsAffected, err := dal.MustExec("INSERT INTO public.exercise_details( user_id, exercise_type, duration, calories_burned, date) VALUES ( $1, $2, $3, $4, $5);", session.Values["userID"], exercise.ExerciseType, exercise.Duration, exercise.CaloriesBurned, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}

func AddMealDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var meal models.Meal
	db = dal.GetDB()
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	RowsAffected, err := dal.MustExec(
		"INSERT INTO public.meal_details( user_id, meal_type, ingredients, calories_consumed, date) VALUES ( $1, $2, $3, $4, $5);", session.Values["userID"], meal.MealType, meal.Ingeredients, meal.CaloriesConsumed, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")
}
func AddWeightDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var weight models.Weight
	db = dal.GetDB()
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	RowsAffected, err := dal.MustExec(
		"INSERT INTO public.weight_details( user_id, daily_weight, date) VALUES ( $1, $2, $3);", session.Values["userID"], weight.DailyWeight, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}
func AddWaterDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var water models.Water
	db = dal.GetDB()
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	RowsAffected, err := dal.MustExec("INSERT INTO public.water_details( user_id, water_intake, date) VALUES ( $1, $2, $3);", session.Values["userID"], water.WaterIntake, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}
