package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"

	// "log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/lib/pq"

	// "github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var errMessage models.MyError

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

func validSession(r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "val")
	fmt.Println(session)
	if err != nil {
		fmt.Println("Error no session Found !", err)
		session.Options.MaxAge = -1
		return session, err
	}
	if session.Values["authenticated"] != true {
		err = errors.New("need to login First")
		fmt.Println("First login")
		return session, err
	}
	return session, err
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
func errorShow(code int, message string) []byte {
	errMessage.Code = code
	errMessage.Message = message
	user_data, _ := json.MarshalIndent(errMessage, "", "  ")
	return user_data
}
func databaseErrorShow(err error) string {
	if dbErr, ok := err.(*pq.Error); ok { // For PostgreSQL database driver (pq)
		// Access PostgreSQL-specific error fields
		// errCode,_ :=  strconv.Atoi(dbErr.Code)
		errMessage := dbErr.Detail
		// Handle the PostgreSQL-specific error
		return errMessage
	}
	return "Databse Error"
}
func UserSignup(w http.ResponseWriter, r *http.Request) {

	var userSignup models.UserSignup
	db = dal.GetDB()

	_, err := dataReadFromBody(r, &userSignup)
	if err != nil {
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
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
		errorMessage := errorShow(400, err.Error())
		w.WriteHeader(400)
		w.Write(errorMessage)
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
		c := &http.Cookie{
			Name:     "val",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		aes256Key, _ = generateAESKey(32)
		store = sessions.NewCookieStore(aes256Key)
		session, _ := store.Get(r, "val")
		session.Values["userID"] = userLogin.UserID
		session.Values["authenticated"] = true
		err = session.Save(r, w)
		if err != nil {
			errorShow(500, "Session storing error"+err.Error())
			return
		}
		fmt.Fprintf(w, "Successfull login")

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

func Logout(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "val")
	session.Options.MaxAge = -1
	if err != nil {
		fmt.Println("Error no session Found !")
		fmt.Fprintf(w, "Successfull logout")
		return
	}
	fmt.Println(session.Values)
	session.Values["authenticated"] = false

	err = session.Save(r, w)
	if err != nil {
		errorShow(500, "Internal server error")
		return
	}

	fmt.Fprintf(w, "Successfull logout")
}
