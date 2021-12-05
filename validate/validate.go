package validate

import (
	"fmt"
	"net/mail"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 64
	maxStringLength = 255
	minStringLength = 3

	emptyStringError = "Empty string error!"
	tooShortPasswordError = "Too short password! Password must be at least 8 characters."
	tooLongPasswordError = "Too long password! Password must`nt be more than 64 characters."
	tooLongStringError = "Too long string! Max length 255 symbols."
	tooShortStringError = "Too short name!"
	emailFormatError = "Please enter your email address in format: yourname@example.com"
)

func Email(email string) error {
	if isEmpty(email) {
		return fmt.Errorf("%v", emptyStringError)
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("%v", emailFormatError)
	}
	return nil
}

func Password(password string) error {
	if isEmpty(password) {
		return fmt.Errorf("%v", emptyStringError)
	}
	if len(password) < minPasswordLength {
		return fmt.Errorf("%v", tooShortPasswordError)
	}
	if len(password) > maxPasswordLength {
		return fmt.Errorf("%v", tooLongPasswordError)
	}
	return nil
}

func Name(name string) error {
	if isEmpty(name) {
		return fmt.Errorf("%v", emptyStringError)
	}
	if len(name) > maxStringLength {
		return fmt.Errorf("%v", tooLongStringError)
	}
	if len(name) < minStringLength {
		return fmt.Errorf("%v", tooShortStringError)
	}
	return nil
}

func isEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}