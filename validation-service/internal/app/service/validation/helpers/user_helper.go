package helpers

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"time"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/validation-service/internal/app/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	emailRegex     = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	passwordLength = 8
	sexVals        = []string{"male", "female"}
)

func ValidatePassword(pass, passConf string) error {
	if passConf == "" {
		log.Println("password confirm is missing, but now its OK")
	}

	if len(pass) < passwordLength {
		log.Println(errVals.ErrInvalidPasswordText)
		return fmt.Errorf(errVals.ErrInvalidPasswordText)
	}

	return nil
}

func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		log.Println(errVals.ErrInvalidEmailText)
		return fmt.Errorf(errVals.ErrInvalidEmailText)
	}

	return nil
}

func ValidateBirthdate(birthdate int) error {
	ts := timestamppb.New(time.Now())
	if int(ts.Seconds) < birthdate {
		log.Println(errVals.ErrInvalidBirthdateCode)
		return fmt.Errorf(errVals.ErrInvalidBirthdateText)
	}

	return nil
}

func ValidateSex(sex string) error {
	if !slices.Contains(sexVals, sex) {
		log.Println(errVals.ErrInvalidSexCode)
		return fmt.Errorf(errVals.ErrInvalidSexText)
	}

	return nil
}
