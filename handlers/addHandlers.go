package handlers

import (
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"time"
)

// AddExerciseDetails example
//
// @tags Exercise
// @Security UserIDAuth
//	@Summary		Add a new exercise for today for today
//	@Description	add new exercise details with ExerciseType, Duration, CaloriesBurned, Date(default)
//	@ID				user-exercise-add
//	@Accept			json
//	@Produce		json
// @Param request body models.Exercise true "The input for add exercise"
//	@Success		200				string		"User details Successfully added"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/exercise/ [post]
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

// AddMealDetails example
//
// @tags Meal
// @Security UserIDAuth
//	@Summary		Add a new meal details for today
//	@Description	add new meal details with MealType, Ingredients, CaloriesConsumed, Date(default)
//	@ID				user-meal-add
//	@Accept			json
//	@Produce		json
// @Param request body models.Meal true "The input for add meal"
//	@Success		200				string		"User details Successfully added"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/meal/ [post]
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

// AddWeightDetails example
//
// @tags Weight
// @Security UserIDAuth
//	@Summary		Add weight for today
//	@Description	add weight details with in DailyWeight(in kg), Date(default)
//	@ID				user-weight-add
//	@Accept			json
//	@Produce		json
// @Param request body models.Weight true "The input for add weight"
//	@Success		200				string		"User details Successfully added"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/weight/ [post]
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

// AddWaterDetails example
//
// @tags Water
// @Security UserIDAuth
//	@Summary		Add water consumption for today
//	@Description	add water details with in WaterIntake(in litre), Date(default)
//	@ID				user-water-add
//	@Accept			json
//	@Produce		json
// @Param request body models.Water true "The input for add daily water"
//	@Success		200				string		"User details Successfully added"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/user/water/ [post]
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
