package handlers

import (
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"
)

func DeleteMealDetails(w http.ResponseWriter, r *http.Request) {
	mealType := r.FormValue("mealtype")
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details WHERE user_id=$1 AND date=$2 AND meal_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), mealType)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully deleted", w)
}

func DeleteExerciseDetails(w http.ResponseWriter, r *http.Request) {
	ExerciseType := r.FormValue("exercisetype")
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details WHERE user_id=$1 AND date=$2 AND exercise_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), ExerciseType)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully deleted", w)
}

func DeleteWeightDetails(w http.ResponseWriter, r *http.Request) {
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.weight_details WHERE user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully deleted", w)
}

func DeleteWaterDetails(w http.ResponseWriter, r *http.Request) {
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.water_details WHERE user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully deleted", w)
}
