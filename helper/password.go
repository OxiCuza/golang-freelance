package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)

	return string(bytes), nil
}
