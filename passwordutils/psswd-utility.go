package passwordutils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (encryptedPassword string, err error) {

	var encryptedPasswordInBytes []byte

	if encryptedPasswordInBytes, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		err := fmt.Errorf("failed to hash password: %s", err)
		return "", err
	}

	encryptedPassword = string(encryptedPasswordInBytes)
	return
}

func PasswordsAreEqual(rawFormPassword, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawFormPassword)); err != nil {
		return false
	}
	return true
}

