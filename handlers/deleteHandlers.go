package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"net/http"
	"time"
)

func DeleteMealDetails(w http.ResponseWriter, r *http.Request) {

	mealType := r.FormValue("mealtype")
	//fmt.Fprintf(w, " %s\n", mealType)
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details where user_id=$1 AND date=$2 AND meal_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), mealType)

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

func DeleteExerciseDetails(w http.ResponseWriter, r *http.Request) {

	ExerciseType := r.FormValue("exercisetype")
	//fmt.Fprintf(w, " %s\n", ExerciseType)

	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details where user_id=$1 AND date=$2 AND exercise_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), ExerciseType)

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

func DeleteWeightDetails(w http.ResponseWriter, r *http.Request) {

	var RowsAffected int64

	RowsAffected, err = dal.MustExec("DELETE FROM public.weight_details where user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))

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

func DeleteWaterDetails(w http.ResponseWriter, r *http.Request) {

	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.water_details where user_id=$1 AND date=$2 ;", UserID.UserID, time.Now().Format("2006-01-02"))

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
