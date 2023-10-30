package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"net/http"
	"time"
)

func EditUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var user models.Users

	var RowsAffected int64
	if r.Method == "PUT" {
		_, err = dataReadFromBody(r, &user)
		if err != nil {
			errorMessage := errorShow(400, err.Error())
			w.WriteHeader(400)
			w.Write(errorMessage)
			return
		}
		RowsAffected, err = dal.MustExec("UPDATE public.user_profile_details set  age=$2, gender=$3, height=$4, weight=$5, health_goal=$6, profile_photo=$7  where user_id=$1 ;", session.Values["userID"], user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.user_profile_details where user_id=$1 ;", session.Values["userID"])
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}
func EditMealDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var meal models.Meal
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.meal_details set ingredients=$1, calories_consumed=$2  where user_id=$3 AND date=$4 AND meal_type=$5;", meal.Ingeredients, meal.CaloriesConsumed, session.Values["userID"],time.Now().Format("2006-01-02"), meal.MealType)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details where user_id=$1 AND date=$2 AND meal_type=$3;", session.Values["userID"], time.Now().Format("2006-01-02"), meal.MealType)
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditExerciseDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var exercise models.Exercise
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.exercise_details set duration=$1, calories_burned=$2  where user_id=$3 AND date=$4 AND exercise_type=$5;", exercise.Duration, exercise.CaloriesBurned, session.Values["userID"], time.Now().Format("2006-01-02"), exercise.ExerciseType)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details where user_id=$1 AND date=$2 AND exercise_type=$3;", session.Values["userID"], time.Now().Format("2006-01-02"), exercise.ExerciseType)
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditWeightDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var weight models.Weight
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.weight_details set daily_weight=$2  where user_id=$2 AND date=$3 ;", weight.DailyWeight, session.Values["userID"], time.Now().Format("2006-01-02"))
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.weight_details where user_id=$1 AND date=$2 ;", session.Values["userID"], time.Now().Format("2006-01-02"))
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditWaterDetails(w http.ResponseWriter, r *http.Request) {
	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	var water models.Water
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.water_details set water_intake=$2  where user_id=$2 AND date=$3 ;", water.WaterIntake, session.Values["userID"], time.Now().Format("2006-01-02"))
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.water_details where user_id=$1 AND date=$2 ;", session.Values["userID"], time.Now().Format("2006-01-02"))
	}
	if err != nil {
		databaseErrorMessage := databaseErrorShow(err)
		errorMessage := errorShow(409, databaseErrorMessage)
		w.WriteHeader(409)
		w.Write(errorMessage)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}
