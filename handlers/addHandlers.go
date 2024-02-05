package handlers

import (
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"
)

func AddUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	_, err := dal.MustExec("UPDATE public.user_details SET age=$2, gender=$3, height=$4, weight=$5, health_goal=$6, profile_photo=$7  WHERE user_id=$1 ;", UserID.UserID, user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "User details Successfully added", w)
}

func AddExerciseDetails(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	_, err := dal.MustExec("INSERT INTO public.exercise_details( user_id, exercise_type, duration, calories_burned, date) VALUES ( $1, $2, $3, $4, $5);", UserID.UserID, exercise.ExerciseType, exercise.Duration, exercise.CaloriesBurned, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "User details Successfully added", w)
}

func AddMealDetails(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	_, err := dal.MustExec(
		"INSERT INTO public.meal_details( user_id, meal_type, ingredients, calories_consumed, date) VALUES ( $1, $2, $3, $4, $5);", UserID.UserID, meal.MealType, meal.Ingeredients, meal.CaloriesConsumed, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "User details Successfully added", w)
}

func AddWeightDetails(w http.ResponseWriter, r *http.Request) {
	var weight models.Weight
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	_, err := dal.MustExec(
		"INSERT INTO public.weight_details( user_id, daily_weight, date) VALUES ( $1, $2, $3);", UserID.UserID, weight.DailyWeight, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "User details Successfully added", w)
}

func AddWaterDetails(w http.ResponseWriter, r *http.Request) {
	var water models.Water
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	_, err := dal.MustExec("INSERT INTO public.water_details( user_id, water_intake, date) VALUES ( $1, $2, $3);", UserID.UserID, water.WaterIntake, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "User details Successfully added", w)
}
