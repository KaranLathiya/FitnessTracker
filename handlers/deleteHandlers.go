package handlers

import (
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"
)

// DeleteMealDetails example
//
// @tags Meal
// @Security UserIDAuth
//	@Summary		delete meal details of today
//	@Description	delete meal details of today with MealType
//	@ID				user-meal-delete
//	@Accept			json
//	@Produce		json
// @Param   mealtype     query     string     true  "mealtype for which want to delete details"     example("breakfast")
//	@Success		200				string		"User details Successfully deleted"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/meal [delete]
func DeleteMealDetails(w http.ResponseWriter, r *http.Request) {
	var mealType models.MealType
	mealType.MealType = r.FormValue("mealtype")
	err = validate.Struct(mealType)
	if err != nil {
		fmt.Println("Error in passing data through json")
		response.MessageShow(400, "Invalid data", w)
		return
	}
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details WHERE user_id=$1 AND date=$2 AND meal_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), mealType.MealType)
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

// DeleteExerciseDetails example
//
// @tags Exercise
// @Security UserIDAuth
//	@Summary		delete exercise details of today
//	@Description	delete exercise details of today with ExerciseType
//	@ID				user-exercise-delete
//	@Accept			json
//	@Produce		json
// @Param   exercisetype     query     string     true  "exercisetype for which want to delete details"     example("weight_lifting")
//	@Success		200				string		"User details Successfully deleted"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/exercise [delete]
func DeleteExerciseDetails(w http.ResponseWriter, r *http.Request) {
	var exerciseType models.ExerciseType
	exerciseType.ExerciseType = r.FormValue("exercisetype")
	err = validate.Struct(exerciseType)
	if err != nil {
		fmt.Println("Error in passing data through json")
		response.MessageShow(400, "Invalid data", w)
		return
	}
	exerciseType.ExerciseType = r.FormValue("exercisetype")
	var RowsAffected int64
	RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details WHERE user_id=$1 AND date=$2 AND exercise_type=$3;", UserID.UserID, time.Now().Format("2006-01-02"), exerciseType.ExerciseType)
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

// DeleteWeightDetails example
//
// @tags Weight
// @Security UserIDAuth
//	@Summary		delete weight details of today
//	@Description	delete weight details of today 
//	@ID				user-weight-delete
//	@Accept			json
//	@Produce		json
//	@Success		200				string		"User details Successfully deleted"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/weight [delete]
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

// DeleteWaterDetails example
//
// @tags Water
// @Security UserIDAuth
//	@Summary		delete water details of today
//	@Description	delete water details of today 
//	@ID				user-water-delete
//	@Accept			json
//	@Produce		json
//	@Success		200				string		"User details Successfully deleted"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/water [delete]
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
