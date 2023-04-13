package errorsHandling

import "errors"

var (
	ErrEmailAlreadyExist    = errors.New("ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
	ErrUsernameAlreadyExist = errors.New("ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)")
	ErrDataNotFound         = errors.New("record not found")
	ErrEmailNotFound        = errors.New("email not found")
	ErrUsernameNotFound     = errors.New("username not found")
	ErrUserNotFound         = errors.New("No user found on with that ID")
)

type Form struct {
	Field   string
	Message string
}
