package errors

import "errors"

var (
	ServerError         = GenerateError("Something went wrong! Please try again later")
	UserNotExist        = GenerateError("User not exists")
	UnauthorisedError   = GenerateError("You are not authorised to perform this action")
	TimeStampError      = GenerateError("time should be a unix timestamp")
	InternalServerError = GenerateError("internal server error")
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func GenerateError(err string) error {
	return errors.New(err)
}
