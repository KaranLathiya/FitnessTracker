package handlers

import (
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UpdateUserProfileDetails example
//
// @tags UserProfile
// @Security UserIDAuth
//	@Summary		update user profile details 
//	@Description	update user profile details like Email, FullName, Age, Gender, Height, Weight, HealthGoal, ProfilePhoto
//	@ID				user-profile-update
//	@Accept			json
//	@Produce		json
// @Param request body models.Users true "The input for update the user profile details"
//	@Success		200				string		"User details Successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/profile/ [put]
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

// UpdateMealDetails example
//
// @tags Meal
// @Security UserIDAuth
//	@Summary		update meal details of today
//	@Description	update meal details like Ingredients, MealType, CaloriesConsumed
//	@ID				user-meal-update
//	@Accept			json
//	@Produce		json
// @Param request body models.Meal true "The input for update the meal details"
//	@Success		200				string		"User details Successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/meal/ [put]
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

// UpdateExerciseDetails example
//
// @tags Exercise
// @Security UserIDAuth
//	@Summary		update exercise details of today
//	@Description	update exercise details like Duration, ExerciseType, CaloriesBurned
//	@ID				user-exercise-update
//	@Accept			json
//	@Produce		json
// @Param request body models.Exercise true "The input for update the exercise details"
//	@Success		200				string		"User details Successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/exercise/ [put]
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

// UpdateWeightDetails example
//
// @tags Weight
// @Security UserIDAuth
//	@Summary		update weight details of today
//	@Description	update daily weight details 
//	@ID				user-weight-update
//	@Accept			json
//	@Produce		json
// @Param request body models.Weight true "The input for update the weight details"
//	@Success		200				string		"User details Successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/weight/ [put]
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

// UpdateWaterDetails example
//
// @tags Water
// @Security UserIDAuth
//	@Summary		update water details of today
//	@Description	update daily water details 
//	@ID				user-water-update
//	@Accept			json
//	@Produce		json
// @Param request body models.Water true "The input for update the water details"
//	@Success		200				string		"User details Successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/water/ [put]
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

// UpdateUserPassword example
//
// @tags User
// @Security UserIDAuth
//	@Summary		set new password for user 
//	@Description	set new password with CurrentPassword, NewPassword 
//	@ID				user-password-update
//	@Accept			json
// @Param request body models.UpdateUserPassword true "The input for set new password"
//	@Produce		json
//	@Success		200				string		"User password successfully updated"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data / current password and new password can't be same"
//	@Failure		401		{object}	models.Message	"Email id doesn't exist"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/change-password [post]
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	var updateUserPassword models.UpdateUserPassword
	var RowsAffected int64
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &updateUserPassword)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	if updateUserPassword.CurrentPassword == updateUserPassword.NewPassword {
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
	err := bcrypt.CompareHashAndPassword([]byte(currentHashedpassword), []byte(updateUserPassword.CurrentPassword))
	if err != nil {
		response.MessageShow(401, "Wrong password", w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(updateUserPassword.NewPassword), 14)
	updateUserPassword.NewPassword = string(bytes)
	RowsAffected, err = dal.MustExec("UPDATE public.user_details SET password=$1 WHERE user_id=$2;", updateUserPassword.NewPassword, UserID.UserID)
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
