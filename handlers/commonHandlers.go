package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"karanlathiya/FitnessTracker/dal"
	"karanlathiya/FitnessTracker/models"
	"karanlathiya/FitnessTracker/response"
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
			response.MessageShow(498, "Invalid token", w)
			return
		}
		db := dal.GetDB()
		errIfNoRows := db.QueryRow("SELECT user_id FROM public.user_details WHERE user_id=$1", UserID.UserID).Scan(&UserID.UserID)
		if errIfNoRows != nil {
			if errIfNoRows.Error() == "sql: no rows in result set" {
				response.MessageShow(498, "Invalid token", w)
				return
			}
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
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
	// fmt.Println(UserID.UserID)
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

func dataReadFromBody(r *http.Request, bodyData interface{}) (interface{}, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return bodyData, err
	}
	// fmt.Println(string(body))
	json.Unmarshal(body, &bodyData)
	err = validate.Struct(bodyData)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid data")
	}
	// fmt.Println(bodyData)
	return bodyData, err
}

// UserSignup example
//
// @tags User
//	@Summary		Add a new user
//	@Description	add new user details with Email, FullName, Password
//	@ID				user-signup
//	@Accept			json
//	@Produce		json
// @Param request body models.UserSignup true "The input for add new user"
//	@Success		200		{object}	models.UserID	
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/signup [post]
func UserSignup(w http.ResponseWriter, r *http.Request) {
	var userSignup models.UserSignup
	var userID models.UserID
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &userSignup)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(userSignup.Password), 14)
	userSignup.Password = string(bytes)
	errIfNoRows := db.QueryRow("INSERT INTO public.user_details( email, fullname, password) VALUES ( $1, $2, $3) RETURNING user_id ;", userSignup.Email, userSignup.FullName, userSignup.Password).Scan(&userID.UserID)
	if errIfNoRows != nil {
		if errIfNoRows.Error() == "sql: no rows in result set" {
			response.MessageShow(401, "Email id doesn't exist", w)
			return
		}
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	userID_data, _ := json.MarshalIndent(userID, "", "  ")
	w.Write(userID_data)
}

// UserLogin example
//
// @tags User
//	@Summary		login for a user
//	@Description	login user with Email, Password
//	@ID				user-login
//	@Accept			json
//	@Produce		json
// @Param request body models.UserLogin true "The input for login for user"
//	@Success		200		{object}	models.UserID	
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		401		{object}	models.Message	"Email id doesn't exist / Wrong password"
//	@Failure		409		{object}	models.Message	"This record contains duplicated data that conflicts with what is already in the database"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/login [post]
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var userID models.UserID
	var userLogin models.UserLogin
	var user models.UserLogin
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &user)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	errIfNoRows := db.QueryRow("SELECT user_id, email , password FROM public.user_details WHERE email=$1", user.Email).Scan(&userID.UserID, &userLogin.Email, &userLogin.Password)
	if errIfNoRows != nil {
		if errIfNoRows.Error() == "sql: no rows in result set" {
			response.MessageShow(401, "Email id doesn't exist", w)
			return
		}
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(errIfNoRows)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))
	if err != nil {
		response.MessageShow(401, "Wrong password", w)
		return
	}
	userID_data, _ := json.MarshalIndent(userID, "", "  ")
	w.Write(userID_data)
}

// OTPRequest example
//
// @tags User
//	@Summary		otp for forgot user password
//	@Description	send otp in registered email for set new user password in case of forgot password with Email, EventType
//	@ID				user-otprequest
//	@Accept			json
// @Param request body models.RequestOTP true "The input for otp for forgot password"
//	@Produce		json
//	@Success		200				string		"OTP sent to email Successfully"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data"
//	@Failure		401		{object}	models.Message	"Email id doesn't exist""
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/otp/request [post]
func OTPRequest(w http.ResponseWriter, r *http.Request) {
	var requestOTP models.RequestOTP
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &requestOTP)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var email string
	errIfNoRows := db.QueryRow("SELECT email FROM public.user_details WHERE email=$1", requestOTP.Email).Scan(&email)
	if errIfNoRows != nil {
		if errIfNoRows.Error() == "sql: no rows in result set" {
			response.MessageShow(401, "Email id doesn't exist", w)
			return
		}
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(errIfNoRows)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	otp, err := generateOTP(6)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	err = sendMail(email, otp)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(otp), 14)
	hashedOTP := string(bytes)
	err = storeOTP(email, hashedOTP, requestOTP.EventType)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	response.MessageShow(200, "OTP sent to email Successfully", w)
}

// VerifyOTP example
//
// @tags User
//	@Summary		verify otp for forgot user password
//	@Description	otp verification for otp sent in registered email for set new user password in case of forgot password with Email, EventType, OTP
//	@ID				user-verifyotp
//	@Accept			json
// @Param request body models.VerifyOTP true "The input for verify otp for forgot password"
//	@Produce		json
//	@Success		200		{object}	models.Token	
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		410		{object}	models.Message	"OTP Expired"
//	@Failure		401		{object}	models.Message	"Invalid OTP""
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/otp/verify [post]
func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var verifyOTP models.VerifyOTP
	db := dal.GetDB()
	currentFormattedTime := currentTimeConvertToCurrentFormattedTime()
	_, err = dataReadFromBody(r, &verifyOTP)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	rows, err := db.Query("SELECT otp, case WHEN $3>expires_at THEN true ELSE false END AS otp_expired FROM public.otp_details WHERE email=$1 AND event_type=$2", verifyOTP.Email, verifyOTP.EventType, currentFormattedTime)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var storedOTP string
		var expiredOTP bool
		err := rows.Scan(&storedOTP, &expiredOTP)
		if err != nil {
			response.MessageShow(400, err.Error(), w)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(storedOTP), []byte(verifyOTP.OTP))
		if err == nil {
			if expiredOTP {
				response.MessageShow(410, "OTP Expired", w)
				return
			}
			var token models.Token
			token.Token, err = postOTPVerificationProcess(verifyOTP.Email, verifyOTP.EventType)
			if err != nil {
				databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
				response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
				return
			}
			tokenData, _ := json.MarshalIndent(token, "", "  ")
			w.Write(tokenData)
			return
		}
	}
	response.MessageShow(401, "Invalid OTP", w)
}

// ForgotPassword example
//
// @tags User
//	@Summary		set new password for user in case of forgot password
//	@Description	after otp verification set new password with Email, EventType, Token, NewePassword
//	@ID				user-forgotpassword
//	@Accept			json
// @Param request body models.ForgotPasswordInput true "The input for set new password"
//	@Produce		json
//	@Success		200				string		"Password successfully changed"
//	@Failure		498		{object}	models.Message	"Invalid token"
//	@Failure		400		{object}	models.Message	"Invalid data "
//	@Failure		401		{object}	models.Message	"Invalid email or eventType or token"
//	@Failure		500		{object}	models.Message	"Internal server error"
//	@Router			/forget-password [post]
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forgotPaswordInput models.ForgotPasswordInput
	db := dal.GetDB()
	_, err = dataReadFromBody(r, &forgotPaswordInput)
	if err != nil {
		response.MessageShow(400, err.Error(), w)
		return
	}
	var hashedToken string
	errIfNoRows := db.QueryRow("SELECT token FROM public.token_details WHERE email=$1 AND event_type=$2", forgotPaswordInput.Email, forgotPaswordInput.EventType).Scan(&hashedToken)
	if errIfNoRows != nil {
		if errIfNoRows.Error() == "sql: no rows in result set" {
			response.MessageShow(401, "Invalid email or eventType or token", w)
			return
		}
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(errIfNoRows)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(forgotPaswordInput.Token))
	if err == nil {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(forgotPaswordInput.NewPassword), 14)
		hashedNewPassword := string(bytes)
		tx, err := db.Begin()
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		defer tx.Rollback()
		_, err = tx.Exec("UPDATE public.user_details SET password=$1 WHERE email=$2;", hashedNewPassword, forgotPaswordInput.Email)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		_, err = tx.Exec("DELETE FROM public.token_details WHERE email=$1 AND event_type=$2;", forgotPaswordInput.Email, forgotPaswordInput.EventType)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		err = tx.Commit()
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			return
		}
		response.MessageShow(200, "Password successfully changed", w)
		return
	}
	response.MessageShow(401, "Invalid token", w)
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
		"Subject: OTP for set new password\n\n" +
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
// 	response.MessageShow(200, "Successfull logout", w)
// 	return
// }
// fmt.Println(session.Values)
// session.Values["authenticated"] = false

// err = session.Save(r, w)
// if err != nil {
// 	response.MessageShow(500, "Internal server error", w)
// 	return
// }
// 	response.MessageShow(200, "Successfull logout", w)
// }

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
