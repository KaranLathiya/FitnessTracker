package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/errors"
	"karanlathiya/FitnessTracker/models"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/markbates/going/randx"
	"golang.org/x/crypto/bcrypt"
)

const otpChars = "1234567890"
var err error
var UserID models.UserID
var validate = validator.New()
var ok bool

func Authentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserID.UserID, ok = validation(r)
		if !ok {
			errors.MessageShow(498, "Invalid token", w)
			return
		}
		db := dal.GetDB()
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
	// fmt.Println(bodyData)
	return bodyData, err
}

func UserSignup(w http.ResponseWriter, r *http.Request) {

	var userSignup models.UserSignup
	var userID models.UserID
	db := dal.GetDB()

	_, err = dataReadFromBody(r, &userSignup)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userSignup.Password), 14)
	userSignup.Password = string(bytes)
	// fmt.Println(userSignup)
	// fmt.Print(db)

	err = db.QueryRow("INSERT INTO public.user_details( email, fullname, password) VALUES ( $1, $2, $3) returning user_id ;", userSignup.Email, userSignup.FullName, userSignup.Password).Scan(&userID.UserID)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	userID_data, _ := json.MarshalIndent(userID, "", "  ")
	w.Write(userID_data)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	var userID models.UserID
	var userLogin models.UserLogin
	var user models.UserLogin
	db := dal.GetDB()
	// fmt.Println("start")
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}

	// fmt.Println(user)

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

func ForgotPassword(w http.ResponseWriter, r *http.Request) {

	var forgotPasswordInput models.ForgotPasswordInput
	db := dal.GetDB()
	// fmt.Println("start")
	_, err = dataReadFromBody(r, &forgotPasswordInput)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var email string
	errIfNoRows := db.QueryRow("select email from public.user_details where email=$1", forgotPasswordInput.Email).Scan(&email)
	if errIfNoRows == nil {
		otp, err := generateOTP(6)
		if err != nil {
			errors.MessageShow(400, err.Error(), w)
			return
		}
		fmt.Println(otp)
		err = sendMail(email, otp)
		if err != nil {
			errors.MessageShow(400, err.Error(), w)
			return
		}
		bytes, _ := bcrypt.GenerateFromPassword([]byte(otp), 14)
		hashedOTP := string(bytes)
		err = storeOTP(email, hashedOTP, forgotPasswordInput.EventType)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
			errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		errors.MessageShow(200, "OTP sent to email Successfully", w)
		return
	}
	errors.MessageShow(404, "Email id doesn't exist", w)
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var validateOTP models.ValidateOTP
	db := dal.GetDB()
	currentFormattedTime := currentTimeConvertToCurrentFormattedTime()
	_, err = dataReadFromBody(r, &validateOTP)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	rows, err := db.Query("select otp from public.otp_details where email=$1 and event_type=$2 and expires_at >= $3 ", validateOTP.Email, validateOTP.EventType, currentFormattedTime)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
		errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	defer rows.Close()
	for rows.Next() {
		storedOTP := models.ValidateOTP{}
		err := rows.Scan(&storedOTP.OTP)
		if err != nil {
			errors.MessageShow(400, err.Error(), w)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(storedOTP.OTP), []byte(validateOTP.OTP))
		if err == nil {
			var token models.Token
			token.Token, err = postOTPVerificationProcess(validateOTP.Email, validateOTP.EventType)
			if err != nil {
				databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
				errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
				return
			}
			tokenData, _ := json.MarshalIndent(token, "", "  ")
			w.Write(tokenData)
			return
		}
	}
	errors.MessageShow(401, "Invalid OTP", w)
}

func SetNewPassword(w http.ResponseWriter, r *http.Request) {
	var setNewPaswordInput models.SetNewPaswordInput
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &setNewPaswordInput)
	if err != nil {
		errors.MessageShow(400, err.Error(), w)
		return
	}
	var hashedToken string
	errIfNoRows := db.QueryRow("select token from public.token_details where email=$1 and event_type=$2", setNewPaswordInput.Email, setNewPaswordInput.EventType).Scan(&hashedToken)
	if errIfNoRows != nil {
		errors.MessageShow(400, "Invalid email or eventType", w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(setNewPaswordInput.Token))
	if err == nil {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(setNewPaswordInput.NewPassword), 14)
		hashedNewPassword := string(bytes)
		RowsAffected, err := dal.MustExec("UPDATE public.user_details set password=$1 where email=$2;", hashedNewPassword, setNewPaswordInput.Email)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := errors.DatabaseErrorShow(err)
			errors.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		if RowsAffected == 0 {
			errors.MessageShow(400, "Invalid data", w)
			return
		}
		errors.MessageShow(200, "Password successfully changed", w)
		return
	}
	errors.MessageShow(401, "Invalid token", w)
}

func postOTPVerificationProcess(email string, eventType string) (string, error) {
	db := dal.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	_, err = tx.Exec("DELETE FROM public.otp_details WHERE email=$1 AND event_type=$2;", email, eventType)
	if err != nil {
		return "", err
	}
	token := randx.String(8)
	tokenBytes, _ := bcrypt.GenerateFromPassword([]byte(token), 14)
	hashedToken := string(tokenBytes)
	_, err = tx.Exec("UPSERT INTO public.token_details( email, token, event_type) VALUES ( $1, $2, $3)", email, hashedToken, eventType)
	if err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return token, nil
}

func sendMail(email string, otp string) error {
	from := "fitnesstrackerdaily@gmail.com"
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: OTP for new password\n\n" +
		otp
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}
	emailPass := os.Getenv("EMAILPASSWORD")
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, emailPass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func currentTimeConvertToCurrentFormattedTime() string {
	currentTime := time.Now().UTC().Add(time.Minute)
	outputFormat := "2006-01-02 15:04:05-07:00"
	currentFormattedTime := currentTime.Format(outputFormat)
	return currentFormattedTime
}

func storeOTP(email string, otp string, eventType string) error {
	db := dal.GetDB()
	otpExpiryTime := time.Now().UTC().Add(time.Minute * time.Duration(5))
	outputFormat := "2006-01-02 15:04:05-07:00"
	otpExpiryFormattedTime := otpExpiryTime.Format(outputFormat)
	_, err := db.Exec("INSERT INTO public.otp_details( email, otp, event_type, expires_at) VALUES ( $1, $2, $3, $4) ;", email, otp, eventType, otpExpiryFormattedTime)
	if err != nil {
		return err
	}
	return nil
}

func generateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	return string(buffer), nil
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
