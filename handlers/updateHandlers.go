package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"net/http"
	"time"
)

func UpdateUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	var RowsAffected int64
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	RowsAffected, err = dal.MustExec("UPDATE public.user_details set  email=$2, fullname=$3, age=$4, gender=$5, height=$6, weight=$7, health_goal=$8, profile_photo=$9  where user_id=$1 ;", UserID.UserID, user.Email, user.FullName, user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto)

	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		errors.MessageShow(400, "Invalid data", w)
		return
	}

	fmt.Println(RowsAffected)
	errors.MessageShow(200, "User details Successfully updated", w)
}
func UpdateMealDetails(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64

	RowsAffected, err = dal.MustExec("UPDATE public.meal_details set ingredients=$1, calories_consumed=$2  where user_id=$3 AND date=$4 AND meal_type=$5;", meal.Ingeredients, meal.CaloriesConsumed, UserID.UserID, time.Now().Format("2006-01-02"), meal.MealType)

	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		errors.MessageShow(400, "Invalid data", w)
		return
	}
	fmt.Println(RowsAffected)
	errors.MessageShow(200, "User details Successfully updated", w)
}

func UpdateExerciseDetails(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.exercise_details set duration=$1, calories_burned=$2  where user_id=$3 AND date=$4 AND exercise_type=$5;", exercise.Duration, exercise.CaloriesBurned, UserID.UserID, time.Now().Format("2006-01-02"), exercise.ExerciseType)

	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		errors.MessageShow(400, "Invalid data", w)
		return
	}
	fmt.Println(RowsAffected)
	errors.MessageShow(200, "User details Successfully updated", w)
}

func UpdateWeightDetails(w http.ResponseWriter, r *http.Request) {
	var weight models.Weight
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.weight_details set daily_weight=$1  where user_id=$2 AND date=$3 ;", weight.DailyWeight, UserID.UserID, time.Now().Format("2006-01-02"))

	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		errors.MessageShow(400, "Invalid data", w)
		return
	}
	fmt.Println(RowsAffected)
	errors.MessageShow(200, "User details Successfully updated", w)
}

func UpdateWaterDetails(w http.ResponseWriter, r *http.Request) {
	var water models.Water
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.water_details set water_intake=$1  where user_id=$2 AND date=$3 ;", water.WaterIntake, UserID.UserID, time.Now().Format("2006-01-02"))

	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		errors.MessageShow(400, "Invalid data", w)
		return
	}
	fmt.Println(RowsAffected)
	errors.MessageShow(200, "User details Successfully updated", w)
}
