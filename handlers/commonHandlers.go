package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"strings"
	"net/http"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error
var UserID models.UserID
var validate = validator.New()
var ok bool

// var (
// 	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
// 	aes256Key, _ = generateAESKey(32)
// 	// aes256Key = []byte("super-secret-key")
// 	store    = sessions.NewCookieStore(aes256Key)
//
// )

//	func generateAESKey(keyLength int) ([]byte, error) {
//		key := make([]byte, keyLength)
//		_, err := rand.Read(key)
//		if err != nil {
//			return nil, err
//		}
//		return key, nil
//	}

func Authentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserID.UserID, ok = validation(r)
		if !ok {
			errors.MessageShow(498, "Invalid token", w)
			return
		}
		db = dal.GetDB()
		rows, err := db.Query("select user_id from public.user_details where user_id=$1", UserID.UserID)
		if err != nil {
			errors.MessageShow(500, "Internal Server Error", w)
			return
		}
		i := 0
		for rows.Next() {
			// err := rows.Scan(&UserID.UserID)
			// if err != nil {
			// 	fmt.Println("Error scanning row:", err)
			// 	return
			// }
			i += 1
		}
		defer rows.Close()
		if i == 0 {
			errors.MessageShow(498, "Invalid token", w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func validation(r *http.Request) (string, bool) {
	UserID_slice, ok := r.Header["Authorization"]
	if !ok {
		return "", false
	}
	UserID.UserID = strings.Join(UserID_slice, " ")
	fmt.Println(UserID.UserID)
	return UserID.UserID, true
}

func HandleCORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "json/application")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// func validSession(r *http.Request) (*sessions.Session, error) {
// 	session, err := store.Get(r, "val")
// 	fmt.Println(session)
// 	if err != nil {
// 		fmt.Println("Error no session Found !", err)
// 		session.Options.MaxAge = -1
// 		return session, err
// 	}
// 	if session.Values["authenticated"] != true {
// 		err = errors.New("need to login First")
// 		fmt.Println("First login")
// 		return session, err
// 	}
// 	return session, err
// }

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

func UserSignup(w http.ResponseWriter, r *http.Request) {

	var userSignup models.UserSignup
	var userID models.UserID
	db = dal.GetDB()

	_, err = dataReadFromBody(r, &userSignup)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userSignup.Password), 14)
	userSignup.Password = string(bytes)
	fmt.Println(userSignup)
	// fmt.Print(db)

	err = db.QueryRow("INSERT INTO public.user_details( email, fullname, password) VALUES ( $1, $2, $3) returning user_id ;", userSignup.Email, userSignup.FullName, userSignup.Password).Scan(&userID.UserID)
	if err != nil {
		errors.MessageShow(409, err.Error(), w)
		return
	}
	userID_data, _ := json.MarshalIndent(userID, "", "  ")
	w.Write(userID_data)
}
func UserLogin(w http.ResponseWriter, r *http.Request) {

	var userID models.UserID
	var userLogin models.UserLogin
	var user models.UserLogin
	db = dal.GetDB()
	fmt.Println("start")
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}

	fmt.Println(user)

	errIfNoRows := db.QueryRow("select user_id, email , password from public.user_details where email=$1", user.Email).Scan(&userID.UserID, &userLogin.Email, &userLogin.Password)
	if errIfNoRows == nil {
		err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))
		if err != nil {
			errors.MessageShow(401, "Wrong password", w)
			return
		}
		userID_data, _ := json.MarshalIndent(userID, "", "  ")
		w.Write(userID_data)

		//LOGIC FOR COOKIE
		// expirationTime := time.Now().Add(time.Minute * 5)
		// http.SetCookie(w,
		// 	&http.Cookie{
		// 		Name:  "UserID",
		// 		Value: strconv.Itoa(userLogin.UserID),
		// 		Expires:expirationTime,
		// 	})
		return
	}
	errors.MessageShow(404, "Email id doesn't exist", w)

}

// func Logout(w http.ResponseWriter, r *http.Request) {

// session, err := store.Get(r, "val")
// session.Options.MaxAge = -1
// if err != nil {
// 	fmt.Println("Error no session Found !")
// 	errors.MessageShow(200, "Successfull logout", w)
// 	return
// }
// fmt.Println(session.Values)
// session.Values["authenticated"] = false

// err = session.Save(r, w)
// if err != nil {
// 	errors.MessageShow(500, "Internal server error", w)
// 	return
// }
// 	errors.MessageShow(200, "Successfull logout", w)
// }
