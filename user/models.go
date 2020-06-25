package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	emailRE *regexp.Regexp
)

func init() {
	// credit: https://www.golangprograms.com/regular-expression-to-validate-email-address.html
	emailRE = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

type User struct {
	ID           uuid.UUID
	Email        string
	Password     string
	PasswordHash string
}

func (u *User) Validate() error {
	errStr := []string{}
	if u.Email == "" {
		errStr = append(errStr, "email must not be blank")
	}

	if !emailRE.MatchString(u.Email) {
		errStr = append(errStr, "email is not valid")
	}

	if u.Password == "" {
		errStr = append(errStr, "password must not be blank")
	}

	if len(u.Password) < 4 {
		errStr = append(errStr, "password must be at least four character long")
	}

	if len(errStr) > 0 {
		return fmt.Errorf("%s", strings.Join(errStr[:], ", "))
	}
	return nil
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}
