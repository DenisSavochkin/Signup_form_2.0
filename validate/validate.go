package validate

const (
	minPasswordLength = 8
	maxPasswordLength = 64

	emptyStringError = "Empty string error"
	tooShortPasswordError = "Too short password"
	tooLongPasswordError = "Too long password"
)

type regError struct {
	err string
}

func (r regError) Error() string {
	return r.err
}

func Password(password string) error {
	if password == "" {
		return &regError{emptyStringError}
	}

	if len(password) < minPasswordLength {
		return &regError{tooShortPasswordError}
	}

	if len(password) > maxPasswordLength {
		return &regError{tooLongPasswordError}
	}

	return nil
}
