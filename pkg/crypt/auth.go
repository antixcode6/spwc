package crypt

import (
	"fmt"
	"log"
	"syscall"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"golang.org/x/term"
)

// Auth takes a user password and method from caller and
//
// uses this to either encrypt an insert or decrypt on Ls using the user
//
// provided passphrase.
func (f *File) Auth() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	switch f.Method {
	case "Insert":
		password, err := f.EncryptPwd(string(bytePassword))
		if err != nil {
			return "", err
		}
		return password, nil

	case "Ls":
		password, err := f.DecryptPwd(string(bytePassword), f.CypherText)
		if err != nil {
			return "", err
		}
		return password, nil
	}
	return "", nil
}

// EncryptPwd takes a user provided passphrase and encrypts the password (to be stored in .cache).
func (f *File) EncryptPwd(password string) (string, error) {
	armor, err := helper.EncryptMessageWithPassword([]byte(password), f.Password)
	if err != nil {
		return "", err
	}
	return armor, nil
}

// DecryptPwd takes a user provided passphrase and decrypts the password(s) from .cache.
func (f *File) DecryptPwd(password, cypherText string) (string, error) {
	armor, err := helper.DecryptMessageWithPassword([]byte(password), cypherText)
	if err != nil {
		log.Printf("Unable to decrypt password: %v", err)
		return "", err
	}
	return armor, nil
}
