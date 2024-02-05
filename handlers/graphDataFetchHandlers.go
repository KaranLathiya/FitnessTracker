package handlers

import (
	"encoding/json"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
	"net/http"
	"strconv"
)

func FetchYearlyWeightDetails(w http.ResponseWriter, r *http.Request) {
	db := dal.GetDB()
	var yearlyWeight []models.YearlyWeight
	year := r.FormValue("year")
	yearIntValue, err := strconv.Atoi(year)
	if err != nil || 0 > yearIntValue || 9999 < yearIntValue {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	rows, err := db.Query(`
	WITH RECURSIVE month_series AS (
		SELECT
		  DATE '`+year+`-01-01' + i * INTERVAL '1 month' AS month
		FROM
		  generate_series(0, 11, 1) AS i
	    )
	    SELECT
		  EXTRACT(MONTH FROM month) AS month,
		  COALESCE(avg(daily_weight),0)
	    FROM
		  month_series
	    LEFT JOIN
		  public.weight_details  ON DATE_TRUNC('month', month) = DATE_TRUNC('month', public.weight_details.date) and user_id = $1
	    GROUP BY
		  month
	    ORDER BY
		  month;`, UserID.UserID)
	if err != nil {
		response.MessageShow(500, "Internal Server Error", w)
		return
	}
	i := 0
	for rows.Next() {
		emptyYearlyWeight := models.YearlyWeight{}
		yearlyWeight = append(yearlyWeight, emptyYearlyWeight)
		err := rows.Scan(&yearlyWeight[i].Month, &yearlyWeight[i].AverageMonthlyWeight)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		i += 1
		defer rows.Close()
	}
	yearlyWeightDeatils, _ := json.MarshalIndent(yearlyWeight, "", "  ")
	w.Write(yearlyWeightDeatils)
}

func FetchYearlyCaloriesBurnedDetails(w http.ResponseWriter, r *http.Request) {
	db := dal.GetDB()
	var yearlyCaloriesBurned []models.YearlyCaloriesBurned
	year := r.FormValue("year")
	yearIntValue, err := strconv.Atoi(year)
	if err != nil || 0 > yearIntValue || 9999 < yearIntValue {
		response.MessageShow(400, "Invalid data", w)
		return
	}
	rows, err := db.Query(`
	WITH RECURSIVE month_series AS (
		SELECT
		  DATE '`+year+`-01-01' + i * INTERVAL '1 month' AS month
		FROM
		  generate_series(0, 11, 1) AS i
	    )
	    SELECT
	    	  month,
		  COALESCE(avg(calories_burned),0) as average_calories_burned_monthly
	    FROM
		(SELECT
			EXTRACT(MONTH FROM month) AS month,
			COALESCE(sum(calories_burned),0) AS calories_burned
			,date
		FROM
			month_series
		LEFT JOIN
			public.exercise_details  ON DATE_TRUNC('month', month) = DATE_TRUNC('month', public.exercise_details.date) and user_id = $1
		GROUP BY
			month,date  
		ORDER BY
			month)
	    GROUP BY
		month
	    ORDER BY
		month;`, UserID.UserID)
	if err != nil {
		response.MessageShow(500, "Internal Server Error", w)
		return
	}
	i := 0
	for rows.Next() {
		emptyYearlyCaloriesBurned := models.YearlyCaloriesBurned{}
		yearlyCaloriesBurned = append(yearlyCaloriesBurned, emptyYearlyCaloriesBurned)
		err := rows.Scan(&yearlyCaloriesBurned[i].Month, &yearlyCaloriesBurned[i].AverageMonthlyCaloriesBurned)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		i += 1
	}
	defer rows.Close()

	yearlyCaloriesBurnedDeatils, _ := json.MarshalIndent(yearlyCaloriesBurned, "", "  ")
	w.Write(yearlyCaloriesBurnedDeatils)
}

// func FetchWaterIntakeMonthly(w http.ResponseWriter, r *http.Request) {
// 	db := dal.GetDB()
// 	rows, err := db.Query("select water_intake, date  from public.water_details where user_id=$1 AND date >= NOW() - INTERVAL '30 days' order by date desc", UserID.UserID)
// 	if err != nil {
// 		response.MessageShow(500, "Internal Server Error", w)
// 		return
// 	}
// 	var water []models.Water
// 	i := 0
// 	for rows.Next() {
// 		emptyWater := models.Water{}
// 		water = append(water, emptyWater)
// 		err := rows.Scan(&water[i].WaterIntake, &water[i].Date)
// 		if err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			return
// 		}
// 		i += 1
// 	}
// 	defer rows.Close()

// 	dailyWaterDetails, _ := json.MarshalIndent(water, "", "  ")
// 	w.Write(dailyWaterDetails)
// }
