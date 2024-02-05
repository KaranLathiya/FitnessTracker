package handlers

import (
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	var RowsAffected int64
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	RowsAffected, err = dal.MustExec("UPDATE public.user_details SET email=$2, fullname=$3, age=$4, gender=$5, height=$6, weight=$7, health_goal=$8, profile_photo=$9 WHERE user_id=$1 ;", UserID.UserID, user.Email, user.FullName, user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully updated", w)
}

func UpdateMealDetails(w http.ResponseWriter, r *http.Request) {
	var meal models.Meal
	_, err = dataReadFromBody(r, &meal)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.meal_details SET ingredients=$1, calories_consumed=$2  WHERE user_id=$3 AND date=$4 AND meal_type=$5;", meal.Ingeredients, meal.CaloriesConsumed, UserID.UserID, time.Now().Format("2006-01-02"), meal.MealType)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully updated", w)
}

func UpdateExerciseDetails(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	_, err = dataReadFromBody(r, &exercise)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.exercise_details SET duration=$1, calories_burned=$2  WHERE user_id=$3 AND date=$4 AND exercise_type=$5;", exercise.Duration, exercise.CaloriesBurned, UserID.UserID, time.Now().Format("2006-01-02"), exercise.ExerciseType)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully updated", w)
}

func UpdateWeightDetails(w http.ResponseWriter, r *http.Request) {
	var weight models.Weight
	_, err = dataReadFromBody(r, &weight)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.weight_details SET daily_weight=$1  WHERE user_id=$2 AND date=$3 ;", weight.DailyWeight, UserID.UserID, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully updated", w)
}

func UpdateWaterDetails(w http.ResponseWriter, r *http.Request) {
	var water models.Water
	_, err = dataReadFromBody(r, &water)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("UPDATE public.water_details SET water_intake=$1  WHERE user_id=$2 AND date=$3 ;", water.WaterIntake, UserID.UserID, time.Now().Format("2006-01-02"))
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User details Successfully updated", w)
}

func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	var password models.ChangePassword
	var RowsAffected int64
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &password)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	if password.CurrentPassword == password.NewPassword {
		response.MessageShow(400, "current password and new password can't be same", w)
		return
	}
	var currentHashedpassword string
	errIfNoRows := db.QueryRow("SELECT password FROM public.user_details WHERE user_id=$1", UserID.UserID).Scan(&currentHashedpassword)
	if errIfNoRows != nil {
		if errIfNoRows.Error() == "sql: no rows in result set" {
			response.MessageShow(401, "Email id doesn't exist", w)
			return
		}
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(currentHashedpassword), []byte(password.CurrentPassword))
	if err != nil {
		response.MessageShow(401, "Wrong password", w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password.NewPassword), 14)
	password.NewPassword = string(bytes)
	RowsAffected, err = dal.MustExec("UPDATE public.user_details SET password=$1 WHERE user_id=$2;", password.NewPassword, UserID.UserID)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	if RowsAffected == 0 {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	response.MessageShow(200, "User password successfully updated", w)
}
