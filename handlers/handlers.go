package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"

	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	// "github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	aes256Key, _ = generateAESKey(32)
	// aes256Key = []byte("super-secret-key")
	store    = sessions.NewCookieStore(aes256Key)
	validate = validator.New()
)

func generateAESKey(keyLength int) ([]byte, error) {
	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func validSession(r *http.Request) (*sessions.Session, bool) {
	session, err := store.Get(r, "val")
	fmt.Println(session)
	if err != nil {
		fmt.Println("Error no session Found !", err)
		return session, false
	}
	if session.Values["authenticated"] != true {
		fmt.Println("First login")
		return session, false
	}
	return session, true
}

func dataReadFromBody(r *http.Request, bodyData interface{}) (interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return bodyData, err
	}
	// fmt.Println(string(body))
	json.Unmarshal(body, &bodyData)

	err = validate.Struct(bodyData)
	if err != nil {
		fmt.Println("Error in passing data through json")
		return bodyData, err
	}
	fmt.Println(bodyData)
	return bodyData, err
}
func AddUserProfileDetails(w http.ResponseWriter, r *http.Request) {

	var user models.Users
	db = dal.GetDB()
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	_, err := dataReadFromBody(r, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.UserID, _ = session.Values["userID"].(int)
	if _, err := db.Exec(
		"INSERT INTO public.user_profile_details( user_id, age, gender, height, weight, health_goal, profile_photo) VALUES ( $1, $2, $3, $4, $5, $6, $7);", user.UserID, user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "User details Successfully added")
}
func UserSignup(w http.ResponseWriter, r *http.Request) {

	var userSignup models.UserSignup
	db = dal.GetDB()

	_, err := dataReadFromBody(r, &userSignup)
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userSignup.Password), 14)
	userSignup.Password = string(bytes)
	fmt.Println(userSignup)
	// fmt.Print(db)
	if _, err := db.Exec(
		"INSERT INTO public.user_registration_details( email, fullname, password) VALUES ( $1, $2, $3);", userSignup.Email, userSignup.FullName, userSignup.Password); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "User Successfully registered.")
}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserSignup
	var user models.UserSignup
	db = dal.GetDB()
	fmt.Println("start")
	_, err := dataReadFromBody(r, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	errIfNoRows := db.QueryRow("select user_id, email , password from public.user_registration_details where email=$1", user.Email).Scan(&userLogin.UserID, &userLogin.Email, &userLogin.Password)
	if errIfNoRows == nil {
		err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))
		if err != nil {
			fmt.Fprintf(w, "Wrong password")
			return
		}
		session, err := store.Get(r, "val")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(session)
		session.Values["userID"] = userLogin.UserID
		session.Values["authenticated"] = true
		err = session.Save(r, w)
		if err != nil {
			fmt.Println("cant store session")
			return
		}
		fmt.Println(session)
		fmt.Fprintf(w, "Sucessfull login")

		//LOGIC FOR COOKIE
		// expirationTime := time.Now().Add(time.Minute * 5)
		// http.SetCookie(w,
		// 	&http.Cookie{
		// 		Name:  "UserID",
		// 		Value: strconv.Itoa(userLogin.UserID),
		// 		Expires:expirationTime,
		// 	})

		fmt.Println(strconv.Itoa(userLogin.UserID))
		return
	}
	fmt.Fprintf(w, "Invalid Email")
}
func UserProfileShow(w http.ResponseWriter, r *http.Request) {

	session, bool_validSession := validSession(r)
	if !bool_validSession {
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
		fmt.Fprint(w, string(user_data))
		return
	}
	fmt.Fprint(w, "First create profile")
}
func AddExerciseDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var exercise models.Exercise
	db = dal.GetDB()
	_, err := dataReadFromBody(r, &exercise)
	if err != nil {
		fmt.Println(err)
		return
	}

	RowsAffected, err := dal.MustExec("INSERT INTO public.exercise_details( user_id, exercise_type, duration, calories_burned, date) VALUES ( $1, $2, $3, $4, $5);", session.Values["userID"], exercise.ExerciseType, exercise.Duration, exercise.CaloriesBurned, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}

func AddMealDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var meal models.Meal
	db = dal.GetDB()
	_, err := dataReadFromBody(r, &meal)
	if err != nil {
		fmt.Println(err)
		return
	}
	RowsAffected, err := dal.MustExec(
		"INSERT INTO public.meal_details( user_id, meal_type, ingredients, calories_consumed, date) VALUES ( $1, $2, $3, $4, $5);", session.Values["userID"], meal.MealType, meal.Ingeredients, meal.CaloriesConsumed, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")
}
func AddWeightDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var weight models.Weight
	db = dal.GetDB()
	_, err := dataReadFromBody(r, &weight)
	if err != nil {
		fmt.Println(err)
		return
	}
	RowsAffected, err := dal.MustExec(
		"INSERT INTO public.weight_details( user_id, daily_weight, date) VALUES ( $1, $2, $3);", session.Values["userID"], weight.DailyWeight, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}
func AddWaterDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var water models.Water
	db = dal.GetDB()
	_, err := dataReadFromBody(r, &water)
	if err != nil {
		fmt.Println(err)
		return
	}
	RowsAffected, err := dal.MustExec("INSERT INTO public.water_details( user_id, water_intake, date) VALUES ( $1, $2, $3);", session.Values["userID"], water.WaterIntake, time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully added")

}
func EditUserProfileDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var user models.Users
	_, err := dataReadFromBody(r, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.user_profile_details set  age=$2, gender=$3, height=$4, weight=$5, health_goal=$6, profile_photo=$7  where user_id=$1 ;", session.Values["userID"], user.Age, user.Gender, user.Height, user.Weight, user.HealthGoal, user.ProfilePhoto)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.user_profile_details where user_id=$1 ;", session.Values["userID"])
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}
func EditMealDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var meal models.Meal
	_, err := dataReadFromBody(r, &meal)
	if err != nil {
		fmt.Println(err)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.meal_details set ingredients=$1, calories_consumed=$2  where user_id=$3 AND date=$4 AND meal_type=$5;", meal.Ingeredients, meal.CaloriesConsumed, session.Values["userID"], meal.Date, meal.MealType)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.meal_details where user_id=$1 AND date=$2 AND meal_type=$3;", session.Values["userID"], meal.Date, meal.MealType)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditExerciseDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var exercise models.Exercise
	_, err := dataReadFromBody(r, &exercise)
	if err != nil {
		fmt.Println(err)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.exercise_details set duration=$1, calories_burned=$2  where user_id=$3 AND date=$4 AND exercise_type=$5;", exercise.Duration, exercise.CaloriesBurned, session.Values["userID"], exercise.Date, exercise.ExerciseType)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.exercise_details where user_id=$1 AND date=$2 AND exercise_type=$3;", session.Values["userID"], exercise.Date, exercise.ExerciseType)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditWeightDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var weight models.Weight
	_, err := dataReadFromBody(r, &weight)
	if err != nil {
		fmt.Println(err)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.weight_details set daily_weight=$2  where user_id=$2 AND date=$3 ;", weight.DailyWeight, session.Values["userID"], weight.Date)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.weight_details where user_id=$1 AND date=$2 ;", session.Values["userID"], weight.Date)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func EditWaterDetails(w http.ResponseWriter, r *http.Request) {
	session, bool_validSession := validSession(r)
	if !bool_validSession {
		return
	}
	var water models.Water
	_, err := dataReadFromBody(r, &water)
	if err != nil {
		fmt.Println(err)
		return
	}
	var RowsAffected int64
	if r.Method == "PUT" {
		RowsAffected, err = dal.MustExec("UPDATE public.water_details set water_intake=$2  where user_id=$2 AND date=$3 ;", water.WaterIntake, session.Values["userID"], water.Date)
	} else if r.Method == "DELETE" {
		RowsAffected, err = dal.MustExec("DELETE FROM public.water_details where user_id=$1 AND date=$2 ;", session.Values["userID"], water.Date)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(RowsAffected)
	fmt.Fprintf(w, "User details Successfully updated")
}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "val")
	if err != nil {
		fmt.Println("Error no session Found !")
		return
	}
	fmt.Println(session.Values)
	// Revoke users authentication
	fmt.Println("Session found")
	session.Values["authenticated"] = false

	err = session.Save(r, w)
	if err != nil {
		fmt.Println("error in storing sessions")
	}
	fmt.Println(session.Values)
}
