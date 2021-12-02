package validate

import "net/mail"

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

func Email(email string) string {
	if isEmpty(email) {
		return emptyStringError
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return emailFormatError
	}
	return ""
}

func Password(password string) string {
	if isEmpty(password) {
		return emptyStringError
	}
	if len(password) < minPasswordLength {
		return tooShortPasswordError
	}
	if len(password) > maxPasswordLength {
		return tooLongPasswordError
	}
	return ""
}

func Name(name string) string {
	if isEmpty(name) {
		return emptyStringError
	}
	if len(name) > maxStringLength {
		return tooLongStringError
	}
	if len(name) < minStringLength {
		return tooShortStringError
	}
	return ""
}

func isEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}