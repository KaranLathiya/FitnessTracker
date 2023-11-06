package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"net/http"
	"time"
)


func DeleteMealDetails(w http.ResponseWriter, r *http.Request) {
	

	var meal models.Meal
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	
		RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details where user_id=$1 AND date=$2 AND meal_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), meal.MealType)
	
	if err != nil {
		databaseErrorMessage,databaseErrorCode := errors.DatabaseErrorShow(err)
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

func DeleteExerciseDetails(w http.ResponseWriter, r *http.Request) {
	

	var exercise models.Exercise
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	
		RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details where user_id=$1 AND date=$2 AND exercise_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), exercise.ExerciseType)
	
	if err != nil {
		databaseErrorMessage,databaseErrorCode := errors.DatabaseErrorShow(err)
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

func DeleteWeightDetails(w http.ResponseWriter, r *http.Request) {
	

	var weight models.Weight
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	
		RowsAffected, err = dal.MustExec("DELETE FROM public.weight_details where user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))
	
	if err != nil {
		databaseErrorMessage,databaseErrorCode := errors.DatabaseErrorShow(err)
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

func DeleteWaterDetails(w http.ResponseWriter, r *http.Request) {
	

	var water models.Water
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	
		RowsAffected, err = dal.MustExec("DELETE FROM public.water_details where user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))
	
	if err != nil {
		databaseErrorMessage,databaseErrorCode := errors.DatabaseErrorShow(err)
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
