package handlers

import (
	"encoding/json"
	"fmt"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"net/http"
)

func UserProfileShow(w http.ResponseWriter, r *http.Request) {

	session, err := validSession(r)
	if err != nil {
		errorMessage := errorShow(401, "need to login First")
		w.WriteHeader(401)
		w.Write(errorMessage)
		return
	}
	fmt.Println(session.Values["authenticated"])
	db = dal.GetDB()
	var user models.Users
	errIfNoRows := db.QueryRow("select age, gender, height, weight, health_goal, profile_photo from public.user_profile_details where user_id=$1", session.Values["userID"]).Scan(&user.Age, &user.Gender, &user.Height, &user.Weight, &user.HealthGoal, &user.ProfilePhoto)
	errIfZeroRows := db.QueryRow("select email, fullname from public.user_registration_details where user_id=$1", session.Values["userID"]).Scan(&user.Email, &user.FullName)
	if errIfNoRows == nil || errIfZeroRows == nil {
		user.UserID, _ = session.Values["userID"].(int)
		user_data, _ := json.MarshalIndent(user, "", "  ")
		w.Write(user_data)
		return
	}
	fmt.Fprint(w, "First create profile")
}
