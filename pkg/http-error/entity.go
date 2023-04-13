package errorsHandling

var (
	ErrEmailAlreadyExist    = "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)"
	ErrUsernameAlreadyExist = "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)"
)

type Form struct {
	Field   string
	Message string
}
