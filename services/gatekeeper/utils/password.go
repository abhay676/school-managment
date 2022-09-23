package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func CompareHashPassword(encyptPwd string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(encyptPwd), []byte(pwd))
}
