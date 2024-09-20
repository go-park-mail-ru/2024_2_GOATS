package helpers

import (
	"log"
	"regexp"
	"slices"
	"time"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const PasswordLength = 8
const Male = 0
const Female = 1
const Other = 2

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	sexVals    = []int32{Male, Female, Other}
)

func ValidatePassword(pass, passConf string) error {
	if passConf == "" {
		log.Println("password confirm is missing, but now its OK")
	}

	if len(pass) < PasswordLength {
		log.Println(errVals.ErrInvalidPasswordText.Error())
		return errVals.ErrInvalidPasswordText
	}

	return nil
}

func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		log.Println(errVals.ErrInvalidEmailText.Error())
		return errVals.ErrInvalidEmailText
	}

	return nil
}

func ValidateBirthdate(birthdate int) error {
	ts := timestamppb.New(time.Now())
	if int(ts.Seconds) < birthdate {
		log.Println(errVals.ErrInvalidBirthdateText.Error())
		return errVals.ErrInvalidBirthdateText
	}

	return nil
}

func ValidateSex(sex int32) error {
	if !slices.Contains(sexVals, sex) {
		log.Println(errVals.ErrInvalidSexText.Error())
		return errVals.ErrInvalidSexText
	}

	return nil
}
