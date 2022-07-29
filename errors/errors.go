package errors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidCredentials = status.New(1, "invalid credentials").Err()

	// general api errors
	ErrInvalidArgument     = status.New(2, "invalid argument").Err()
	ErrInternalServerError = status.New(3, "internal server error").Err()
	ErrInvalidLimit        = status.New(4, "invalid limit").Err()
	ErrInvalidPage         = status.New(5, "invalid page").Err()
	ErrSystemError         = status.New(6, "system error").Err()
	ErrPermissionDenied    = status.New(7, "permission denied").Err()
	ErrMarshal             = status.New(8, "marshal failed").Err()
	ErrUnmarshal           = status.New(9, "unmarshal failed").Err()
	ErrGenerateTokenError  = status.New(10, "generate token failed").Err()
)

// ErrorWithMessage wraps detail error
func ErrorWithMessage(err error, message string) error {
	s, ok := status.FromError(err)
	if !ok {
		return errors.Wrap(err, message)
	}

	status := status.New(s.Code(), s.Message()+": "+message)

	return status.Err()
}
