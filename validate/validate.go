package validate

import "net/mail"

const (
	minPasswordLength = 8
	maxPasswordLength = 64
	maxStringLength = 255

	emptyStringError = "Empty string error!"
	tooShortPasswordError = "Too short password! Password must be at least 8 characters."
	tooLongPasswordError = "Too long password! Password must be no more than 64 characters."
	tooLongStringError = "Too long string! Max length 255 symbols."
	emailFormatError = "Please enter your email address in format: yourname@example.com"
)

type regError struct {
	err string
}

func (r regError) Error() string {
	return r.err
}

func Email(email string) error {
	err := isEmpty(email)
	if err != nil {
		return err
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		return regError{emailFormatError}
	}
	return nil
}

func Password(password string) error {
	err := isEmpty(password)
	if err != nil {
		return err
	}

	if len(password) < minPasswordLength {
		return &regError{tooShortPasswordError}
	}

	if len(password) > maxPasswordLength {
		return &regError{tooLongPasswordError}
	}

	return nil
}

func Name(name string) error {
	err := isEmpty(name)
	if err != nil {
		return err
	}

	if len(name) > maxStringLength {
		return &regError{tooLongStringError}
	}

	return nil
}

func isEmpty(str string) error {
	if str == "" {
		return regError{emptyStringError}
	}
	return nil
}