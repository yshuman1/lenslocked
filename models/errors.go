package models

import (
	"strings"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound modelError = "models: resource not found"

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrPasswordIncorrect modelError = "models: incorrect password provided"

	// ErrEmailRequired is returned when email address isnt provided when creating a user
	ErrEmailRequired modelError = "models: Email address is required"

	// ErrEmailInvalid is returned when email address doesnt match requirements
	ErrEmailInvalid modelError = "models: Email address is not valid"

	// ErrEmailTaken is returned when an update or create is attempted with an email address already in use
	ErrEmailTaken modelError = "models: email is already taken"

	// ErrPasswordTooShort returned if pw set is shorter than 8 characters
	ErrPasswordTooShort modelError = "models: password must be at least 8 characters long"

	// ErrPasswordRequired is returned when a user create is attempted without a pw being provided
	ErrPasswordRequired modelError = "models: password is required"

	// ErrRememberRequired is returned when a create or update is attempted without a user remember token
	ErrRememberRequired modelError = "models: remember token is required"

	// ErrIDInvalid is returned when an invalid ID is provided
	// to a method like Delete.
	ErrIDInvalid privateError = "models: ID provided was invalid"

	// ErrRememberTooShort is returned when a remember token is not at least 32 bytes
	ErrRememberTooShort privateError = "models: remember token must be at least 32 bytes"

	// ErrUserIDRequired is returned when a user ID is incorrect or not provied
	ErrUserIDRequired privateError = "models: user ID is required"

	// ErrTokenInvalid is returned when token given during pw reset isnt any good
	ErrTokenInvalid modelError = "models: token provided is not valid"

	// ErrTitleRequired is returned if a gallery is created without a title
	ErrTitleRequired modelError = "models: title is required"
)
