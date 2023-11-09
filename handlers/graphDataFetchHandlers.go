package handlers

import (
	"encoding/json"
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"net/http"
)

func FetchYearlyWeightDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	var yearlyWeight []models.YearlyWeight
	date := r.FormValue("date")
	year := date[:4]
	rows, err := db.Query("select date_part('month', date) ,avg(daily_weight) FROM public.weight_details where user_id=$1 and date_part('year', date)=$2 GROUP BY date_part('month', date);", UserID.UserID, year)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	i := 0
	for rows.Next() {
		emptyYearlyWeight := models.YearlyWeight{}
		yearlyWeight = append(yearlyWeight, emptyYearlyWeight)
		err := rows.Scan(&yearlyWeight[i].Month, &yearlyWeight[i].AverageMonthlyWeight)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()

	yearlyWeightDeatils, _ := json.MarshalIndent(yearlyWeight, "", "  ")
	w.Write(yearlyWeightDeatils)
}

func FetchYearlyCaloriesBurnedDetails(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	var yearlyCaloriesBurned []models.YearlyCaloriesBurned
	date := r.FormValue("date")
	year := date[:4]
	rows, err := db.Query("select date_part('month', date) ,avg(calories_burned) FROM public.exercise_details where user_id=$1 and date_part('year', date)=$2 GROUP BY date_part('month', date);", UserID.UserID, year)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	i := 0
	for rows.Next() {
		emptyYearlyCaloriesBurned := models.YearlyCaloriesBurned{}
		yearlyCaloriesBurned = append(yearlyCaloriesBurned, emptyYearlyCaloriesBurned)
		err := rows.Scan(&yearlyCaloriesBurned[i].Month, &yearlyCaloriesBurned[i].AverageMonthlyCaloriesBurned)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()

	yearlyCaloriesBurnedDeatils, _ := json.MarshalIndent(yearlyCaloriesBurned, "", "  ")
	w.Write(yearlyCaloriesBurnedDeatils)
}
func FetchWaterIntakeMonthly(w http.ResponseWriter, r *http.Request) {
	db = dal.GetDB()
	rows, err := db.Query("select water_intake, date  from public.water_details where user_id=$1 AND date >= NOW() - INTERVAL '30 days' order by date desc", UserID.UserID)
	if err != nil {
		errors.MessageShow(500, "Internal Server Error", w)
		return
	}
	var water []models.Water
	i := 0
	for rows.Next() {
		emptyWater := models.Water{}
		water = append(water, emptyWater)
		err := rows.Scan(&water[i].WaterIntake, &water[i].Date)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		i += 1
	}
	defer rows.Close()

	dailyWaterDetails, _ := json.MarshalIndent(water, "", "  ")
	w.Write(dailyWaterDetails)
}
