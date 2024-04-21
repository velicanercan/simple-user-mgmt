package errors

type UserError struct {
	Code    int
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}

var (
	ErrUserNotFound      = &UserError{Code: 404, Message: "User not found"}
	ErrUserAlreadyExists = &UserError{Code: 409, Message: "User already exists"}
	ErrUserMailExists    = &UserError{Code: 409, Message: "User mail already exists"}
	ErrUserAge           = &UserError{Code: 400, Message: "User age must be greater than 18"}
	ErrInvalidBirthDate  = &UserError{Code: 400, Message: "Invalid birth date format it must be yyyy-mm-dd"}
	ErrInvalidEmail      = &UserError{Code: 400, Message: "Invalid email format"}
)
